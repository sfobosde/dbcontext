package dbcontext

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// Querry operands.
type GlobalOperands interface {
	// Logical "AND".
	And(filters ...GLobalFilter) *GLobalFilter

	// Logical "OR"
	Or(filters ...GLobalFilter) *GLobalFilter

	// Logical "NOT"
	Not(filters ...GLobalFilter) *GLobalFilter
}

// Структура для реализации интерфейса логических операций.
type Operands struct {
	GlobalOperands
}

// Logical "И".
func (glb *Operands) And(globalFilters ...GLobalFilter) *GLobalFilter {
	globalFilter := &GLobalFilter{
		operand: "AND",
		filters: make([]GLobalFilter, len(globalFilters)),
	}

	copy(globalFilter.filters, globalFilters)

	return globalFilter
}

// Logical "OR"
func (glb *Operands) Or(filters ...GLobalFilter) *GLobalFilter {
	globalFilter := &GLobalFilter{
		operand: "OR",
	}

	globalFilter.filters = filters
	return globalFilter
}

// Logical "NOT"
func (glb *Operands) Not(filters ...GLobalFilter) *GLobalFilter {
	globalFilter := &GLobalFilter{
		operand: "NOT",
	}

	globalFilter.filters = filters
	return globalFilter
}

// Filter structure.
type GLobalFilter struct {
	operand      string
	fieldFilter  []FieldFilter
	filters      []GLobalFilter
	relationName string
	joinParams   string
}

// Arragning inner filters.
func (gf *GLobalFilter) arrangeFilters(db *gorm.DB) (*gorm.DB, error) {
	var fieldFilterExpression string
	var fieldFilterValues []interface{}

	if len(gf.fieldFilter) > 0 {
		fieldFilterExpression, fieldFilterValues = gf.arrangeFieldFilters()

		db = userFilter(db, gf, fieldFilterExpression, fieldFilterValues)
	}

	for _, filter := range gf.filters {
		fieldFilterExpression, fieldFilterValues = filter.arrangeGlobalFilters()
		db = userFilter(db, gf, fieldFilterExpression, fieldFilterValues)
	}

	return db, nil
}

// Add filter queery chain.
func userFilter(db *gorm.DB, gf *GLobalFilter, fieldFilterExpression string, fieldFilterValues []interface{}) *gorm.DB {
	if gf.operand == "AND" {
		db = db.Where(fieldFilterExpression, fieldFilterValues...)
	} else if gf.operand == "OR" {
		db = db.Or(fieldFilterExpression, fieldFilterValues...)
	} else if gf.operand == "NOT" {
		db = db.Not(fieldFilterExpression, fieldFilterValues...)
	} else if gf.operand == "JOIN" {
		params := strings.Split(gf.joinParams, ",")

		db = db.
			Joins(fmt.Sprintf("JOIN %s ON %s.%s = %s", params[0], params[0], params[1], params[3])).
			Joins(fmt.Sprintf("JOIN roles ON %s = %s.%s", params[4], params[0], params[2])).
			Where(fieldFilterExpression, fieldFilterValues...)
	}

	return db
}

// Arrange field filters.
func (gf *GLobalFilter) arrangeFieldFilters() (string, []interface{}) {
	fieldFilterExpression := ""
	fieldFilterValues := make([]interface{}, 0)

	if len(gf.fieldFilter) > 0 {
		for _, fieldFilter := range gf.fieldFilter {
			if len(fieldFilterExpression) > 0 {
				fieldFilterExpression = fieldFilterExpression + " " + gf.operand + " "
			}

			if fieldFilter.value != nil {
				fieldFilterExpression = fieldFilterExpression + fieldFilter.fieldName + " " + fieldFilter.operator + " ? "
				fieldFilterValues = append(fieldFilterValues, fieldFilter.value)
			} else {
				fieldFilterExpression = fieldFilterExpression + fieldFilter.fieldName + " IS NULL "
			}

		}
	}

	return fieldFilterExpression, fieldFilterValues
}

// Arrange global filters.
func (gf *GLobalFilter) arrangeGlobalFilters() (string, []interface{}) {
	fieldFilterExpression := ""
	fieldFilterValues := make([]interface{}, 0)

	if len(gf.filters) > 0 {
		for _, filter := range gf.filters {
			expression, values := filter.arrangeGlobalFilters()

			fieldFilterExpression = fieldFilterExpression + expression
			fieldFilterValues = append(fieldFilterValues, values...)
		}
	}

	if len(gf.fieldFilter) > 0 {
		expression, values := gf.arrangeFieldFilters()

		fieldFilterExpression = fieldFilterExpression + expression
		fieldFilterValues = append(fieldFilterValues, values...)
	}

	return fieldFilterExpression, fieldFilterValues
}
