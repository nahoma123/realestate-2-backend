package pgxadapter

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/mmcloughlin/meow"
)

const (
	DefaultTableName = "casbin_rule"
	DefaultTimeout   = time.Second * 10
)

// Adapter represents the github.com/jackc/pgx/v4 adapter for policy storage.
type Adapter struct {
	tableName string
	schema    string
	timeout   time.Duration
	filtered  bool
	db        *pgx.Conn
}

type Filter struct {
	P [][]string
	G [][]string
}

// NewAdapterByDB creates a new adapter with connection conn which must either be a PostgreSQL
// connection or an instance of *pgx.ConnConfig from package github.com/jackc/pgx/v4.
func NewAdapterWithDB(db *pgx.Conn) (*Adapter, error) {
	a := &Adapter{
		tableName: DefaultTableName,
		timeout:   DefaultTimeout,
		db:        db,
	}
	if err := a.createTable(); err != nil {
		return nil, fmt.Errorf("pgxadapter.NewAdapter: %v", err)
	}
	return a, nil
}

func policyLine(ptype string, values ...string) string {
	const sep = ", "
	var sb strings.Builder
	sb.WriteString(ptype)
	for _, v := range values {
		if len(v) == 0 {
			break
		}
		sb.WriteString(sep)
		sb.WriteString(v)
	}
	return sb.String()
}

func (a *Adapter) tableIdentifier() pgx.Identifier {
	if a.schema != "" {
		return pgx.Identifier{a.schema, a.tableName}
	}
	return pgx.Identifier{a.tableName}
}

func (a *Adapter) schemaTable() string {
	if a.schema != "" {
		return fmt.Sprintf("%q.%s", a.schema, a.tableName)
	}
	return a.tableName
}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	var pType, v0, v1, v2, v3, v4, v5, v6, v7 pgtype.Text
	_, err := a.db.QueryFunc(
		ctx,
		fmt.Sprintf(`SELECT "p_type", "v0", "v1", "v2", "v3", "v4", "v5", "v6","v7" FROM %s`, a.schemaTable()),
		nil,
		[]interface{}{&pType, &v0, &v1, &v2, &v3, &v4, &v5, &v6, &v7},
		func(pgx.QueryFuncRow) error {
			persist.LoadPolicyLine(policyLine(pType.String, v0.String, v1.String, v2.String, v3.String, v4.String, v5.String, v6.String, v7.String), model)
			return nil
		},
	)
	if err != nil {
		return err
	}

	a.filtered = false

	return nil
}

func policyID(ptype string, rule []string) string {
	data := strings.Join(append([]string{ptype}, rule...), ",")
	sum := meow.Checksum(0, []byte(data))
	return fmt.Sprintf("%x", sum)
}

func policyArgs(ptype string, rule []string) []interface{} {
	row := make([]interface{}, 10)
	row[0] = pgtype.Text{
		String: policyID(ptype, rule),
		Status: pgtype.Present,
	}
	row[1] = pgtype.Text{
		String: ptype,
		Status: pgtype.Present,
	}
	l := len(rule)
	for i := 0; i < 6; i++ {
		if i < l {
			row[2+i] = pgtype.Text{
				String: rule[i],
				Status: pgtype.Present,
			}
		} else {
			row[2+i] = pgtype.Text{
				Status: pgtype.Null,
			}
		}
	}
	return row
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) error {
	rows := [][]interface{}{}
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			rows = append(rows, policyArgs(ptype, rule))
		}
	}
	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			rows = append(rows, policyArgs(ptype, rule))
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	return a.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		_, err := tx.Exec(context.Background(), fmt.Sprintf("DELETE FROM %s WHERE id IS NOT NULL", a.schemaTable()))
		if err != nil {
			return err
		}
		_, err = tx.CopyFrom(
			context.Background(),
			a.tableIdentifier(),
			[]string{"id", "p_type", "v0", "v1", "v2", "v3", "v4", "v5", "v6", "v7"},
			pgx.CopyFromRows(rows),
		)
		return err
	})
}

func (a *Adapter) insertPolicyStmt() string {
	return fmt.Sprintf(`
		INSERT INTO %s (id, p_type, v0, v1, v2, v3, v4, v5,v6,v7)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9,$10)
	`, a.schemaTable())
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	_, err := a.db.Exec(ctx,
		a.insertPolicyStmt(),
		policyArgs(ptype, rule)...,
	)
	return err
}

// AddPolicies adds policy rules to the storage.
func (a *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	return a.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		b := &pgx.Batch{}
		for _, rule := range rules {
			b.Queue(a.insertPolicyStmt(), policyArgs(ptype, rule)...)
		}
		br := tx.SendBatch(context.Background(), b)
		defer br.Close()
		for range rules {
			_, err := br.Exec()
			if err != nil {
				return err
			}
		}
		return br.Close()
	})
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	id := policyID(ptype, rule)
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	_, err := a.db.Exec(ctx,
		fmt.Sprintf("DELETE FROM %s WHERE id = $1", a.schemaTable()),
		id,
	)
	return err
}

// RemovePolicies removes policy rules from the storage.
func (a *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	return a.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		b := &pgx.Batch{}
		for _, rule := range rules {
			id := policyID(ptype, rule)
			b.Queue(fmt.Sprintf("DELETE FROM %s WHERE id = $1", a.schemaTable()), id)
		}
		br := tx.SendBatch(context.Background(), b)
		defer br.Close()
		for range rules {
			_, err := br.Exec()
			if err != nil {
				return err
			}
		}
		return br.Close()
	})
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	var sb strings.Builder
	_, err := sb.WriteString(fmt.Sprintf("DELETE FROM %s WHERE p_type = $1", a.schemaTable()))
	if err != nil {
		return err
	}
	args := []interface{}{ptype}

	idx := fieldIndex + len(fieldValues)
	for i := 0; i < 6; i++ {
		if fieldIndex <= i && idx > i && fieldValues[i-fieldIndex] != "" {
			args = append(args, fieldValues[i-fieldIndex])
			_, err = sb.WriteString(fmt.Sprintf(" AND v%d = $%d", i, len(args)))
			if err != nil {
				return err
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	_, err = a.db.Exec(ctx, sb.String(), args...)
	return err
}

func (a *Adapter) loadFilteredPolicy(model model.Model, filter *Filter, handler func(string, model.Model)) error {
	var (
		ptype, v0, v1, v2, v3, v4, v5, v6, v7 pgtype.Text
		args                                  []interface{}
		sb                                    = &strings.Builder{}
	)

	fmt.Fprintf(sb, `SELECT "p_type", "v0", "v1", "v2", "v3", "v4", "v5","v6","v7" FROM %s WHERE `, a.schemaTable())

	buildQuery := func(policies [][]string, ptype string) {
		if len(policies) > 0 {
			args = append(args, ptype)
			fmt.Fprintf(sb, `(p_type = $%d AND (`, len(args))
			for i, p := range policies {
				fmt.Fprint(sb, `(`)
				for j, v := range p {
					if v == "" {
						continue
					}
					args = append(args, v)
					fmt.Fprintf(sb, `v%d = $%d`, j, len(args))
					if j < len(p)-1 {
						fmt.Fprint(sb, ` AND `)
					}
				}
				fmt.Fprint(sb, `)`)
				if i < len(policies)-1 {
					fmt.Fprint(sb, ` OR `)
				}
			}
			fmt.Fprint(sb, `))`)
		}
	}

	buildQuery(filter.P, "p")
	if len(filter.P) > 0 && len(filter.G) > 0 {
		fmt.Fprint(sb, ` OR `)
	}
	buildQuery(filter.G, "g")

	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	_, err := a.db.QueryFunc(ctx, sb.String(), args, []interface{}{&ptype, &v0, &v1, &v2, &v3, &v4, &v5, &v6, &v7}, func(qfr pgx.QueryFuncRow) error {
		handler(policyLine(ptype.String, v0.String, v1.String, v2.String, v3.String, v4.String, v5.String, v6.String, v7.String), model)
		return nil
	})
	return err
}

// LoadFilteredPolicy can query policies with a filter. Make sure that filter is of type *pgxadapter.Filter
func (a *Adapter) LoadFilteredPolicy(model model.Model, filter interface{}) error {
	if filter == nil {
		return a.LoadPolicy(model)
	}

	filterValue, ok := filter.(*Filter)
	if !ok {
		return fmt.Errorf("filter must be of type *pgxadapter.Filter")
	}
	err := a.loadFilteredPolicy(model, filterValue, persist.LoadPolicyLine)
	if err != nil {
		return err
	}
	a.filtered = true
	return nil
}

func (a *Adapter) IsFiltered() bool {
	return a.filtered
}

// UpdatePolicy updates a policy rule from storage.
// This is part of the Auto-Save feature.
func (a *Adapter) UpdatePolicy(sec string, ptype string, oldRule, newPolicy []string) error {
	return a.UpdatePolicies(sec, ptype, [][]string{oldRule}, [][]string{newPolicy})
}

// UpdatePolicies updates some policy rules to storage, like db, redis.
func (a *Adapter) UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	return a.db.BeginFunc(ctx, func(t pgx.Tx) error {
		b := &pgx.Batch{}
		for _, rule := range oldRules {
			id := policyID(ptype, rule)
			b.Queue(fmt.Sprintf("DELETE FROM %s WHERE id = $1", a.schemaTable()), id)
		}
		for _, rule := range newRules {
			b.Queue(a.insertPolicyStmt(), policyArgs(ptype, rule)...)
		}
		br := t.SendBatch(context.Background(), b)
		defer br.Close()
		for i := 0; i < b.Len(); i++ {
			_, err := br.Exec()
			if err != nil {
				return err
			}
		}
		return br.Close()
	})
}

// UpdateFilteredPolicies deletes old rules and adds new rules.
func (a *Adapter) UpdateFilteredPolicies(sec string, ptype string, newPolicies [][]string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	return nil, fmt.Errorf("not implemented")
}

func (a *Adapter) Close() {
	if a != nil && a.db != nil {
		a.db.Close(context.Background())
	}
}

func (a *Adapter) createTable() error {
	if a.schema != "" {
		ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
		defer cancel()
		if _, err := a.db.Exec(ctx, fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS %q`, a.schema)); err != nil {
			return err
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()
	_, err := a.db.Exec(ctx, fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
			p_type text,
			v0 text,
			v1 text,
			v2 text,
			v3 text,
			v4 text,
			v5 text,
			v6 text,
			v7 text
		)
	`, a.schemaTable()))
	return err
}
