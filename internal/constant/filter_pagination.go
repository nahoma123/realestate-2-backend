package constant

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Filter struct {
	Field    string
	Value    string
	Operator string
}

type Pagination struct {
	Page    int
	PerPage int
	Sort    map[string]string
}

type Lookup struct {
	Field        string
	As           string
	Collection   string
	ForeignField string
	Filters      []Filter
}

type FilterPagination struct {
	Filters    []Filter
	Pagination Pagination
	Lookups    []Lookup
	TotalCount int64
	TotalPages int
}

func AddFilter(fp FilterPagination, field, value, operator string) *FilterPagination {
	fp.Filters = append(fp.Filters, Filter{
		Field:    field,
		Value:    value,
		Operator: operator,
	})
	return &fp
}

func AddLookup(fp FilterPagination, field, foreignField, as, collection string) *FilterPagination {
	fp.Lookups = append(fp.Lookups, Lookup{
		Field:        field,
		As:           as,
		ForeignField: foreignField,
		Collection:   collection,
	})
	return &fp
}

func ExtractValueFromFilter(fp *FilterPagination, field string) interface{} {
	if fp == nil || fp.Filters == nil {
		return ""
	}

	for _, filter := range fp.Filters {
		if filter.Field == field {
			return filter.Value
		}
	}

	return ""
}

func ParseFilterPagination(c *gin.Context) *FilterPagination {
	var filterPagination FilterPagination
	filterPagination.Filters = make([]Filter, 0)
	filterPagination.Pagination.Page = 1
	filterPagination.Pagination.PerPage = 10
	filterPagination.Pagination.Sort = make(map[string]string)

	// Extract filter and pagination information from query parameters
	if value := c.Query("page"); value != "" {
		page, err := strconv.Atoi(value)
		if err == nil {
			filterPagination.Pagination.Page = page
		}
	}
	if value := c.Query("per_page"); value != "" {
		perPage, err := strconv.Atoi(value)
		if err == nil {
			filterPagination.Pagination.PerPage = perPage
		}
	}
	filterString := c.Query("filter")
	if filterString != "" {
		// Parse filter string into slice of filters
		var filters []Filter
		err := json.Unmarshal([]byte(filterString), &filters)
		if err == nil {
			filterPagination.Filters = filters
		}
	}

	// Extract sort information from query parameters
	sortString := c.Query("sort")
	if sortString != "" {
		// Parse sort string into map of sort fields and their directions
		err := json.Unmarshal([]byte(sortString), &filterPagination.Pagination.Sort)
		if err != nil {
			// Handle error if sort string cannot be parsed
		}
	}

	return &filterPagination
}

func LocationFilter(coordinates []float64, maxDistance int) bson.M {

	return bson.M{
		"$geoNear": bson.M{
			"near":               coordinates,
			"distanceField":      "distance",
			"maxDistance":        maxDistance,
			"spherical":          true,
			"distanceMultiplier": 1000,
		},
	}
}

//	"$geoNear": bson.M{
//		"near": bson.M{
//			"type":        "Point",
//			"coordinates": coordinates,
//		},
//		"distanceField": "distance",
//		"maxDistance":   10000000,
//		"includeLocs":   "location",
//		"spherical":     true,
//	},
func DeleteFilter(fp *FilterPagination, field string) {
	if fp == nil || fp.Filters == nil {
		return
	}

	for i, filter := range fp.Filters {
		if filter.Field == field {
			fp.Filters = append(fp.Filters[:i], fp.Filters[i+1:]...)
		}
	}

}
