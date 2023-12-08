package constant

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

/*
Filters
 1. by gte for date, by gte for age [gte]
 2. equals [=]
 3. contains [contains]
 4. not equals [!=]
*/

func GetResults(ctx context.Context, db *gorm.DB, tableName string, filterPagination *FilterPagination, results *[]map[string]interface{}) error {
	// Create a GORM DB instance for the specified table
	table := db.Table(tableName)

	// Apply filters
	for _, f := range filterPagination.Filters {
		if f.Operator == "gte" {
			table = table.Where(fmt.Sprintf("%s >= ?", f.Field), f.Value)
		} else if f.Operator == "lte" {
			table = table.Where(fmt.Sprintf("%s <= ?", f.Field), f.Value)
		} else if f.Operator == "=" {
			if strings.Contains(f.Value, "||") {
				values := strings.Split(f.Value, "||")

				// Convert values to []interface{}
				valueInterfaces := make([]interface{}, len(values))
				for i, val := range values {
					valueInterfaces[i] = val
				}

				table = table.Where(fmt.Sprintf("%s IN (?)", f.Field), valueInterfaces...)
			} else {
				table = table.Where(fmt.Sprintf("%s = ?", f.Field), f.Value)
			}
		} else if f.Operator == "contains" {
			if strings.Contains(f.Value, "||") {
				values := strings.Split(f.Value, "||")
				orConditions := make([]string, len(values))
				valueInterfaces := make([]interface{}, len(values))
				for i, val := range values {
					orConditions[i] = fmt.Sprintf("%s ILIKE ?", f.Field)
					valueInterfaces[i] = fmt.Sprintf("%%%s%%", val)
				}
				table = table.Where(strings.Join(orConditions, " OR "), valueInterfaces...)
			} else {
				table = table.Where(fmt.Sprintf("%s ILIKE ?", f.Field), fmt.Sprintf("%%%s%%", f.Value))
			}
		} else if f.Operator == "!=" {
			table = table.Where(fmt.Sprintf("%s != ?", f.Field), f.Value)
		} else {
			// Handle other operators
		}
	}

	// Get the total count of documents that match the filter
	var totalCount int64
	if err := table.Count(&totalCount).Error; err != nil {
		return err
	}

	// Apply sorting
	for field, order := range filterPagination.Pagination.Sort {
		if order == "asc" {
			table = table.Order(fmt.Sprintf("%s asc", field))
		} else if order == "desc" {
			table = table.Order(fmt.Sprintf("%s desc", field))
		}
	}

	// Apply pagination
	table = table.Offset((filterPagination.Pagination.Page - 1) * filterPagination.Pagination.PerPage).Limit(filterPagination.Pagination.PerPage)

	// Execute the query and retrieve the results
	if err := table.Find(results).Error; err != nil {
		return err
	}

	filterPagination.TotalCount = totalCount
	filterPagination.TotalPages = int(totalCount/int64(filterPagination.Pagination.PerPage) + 1)

	return nil
}
