package dbcontext

import "fmt"

// Use Tag by template: `dbcontext:"Roles:user_roles,user_id,role_id,users.id,roles.id"`
type LinkedObjectFieldOperands struct {
	FieldValueOperandsParams
	JoinParams string
}

// Select on joined values. values must contains "id" with string type.
func (s *LinkedObjectFieldOperands) Has(values ...interface{}) *GLobalFilter {
	valuesIds := make([]string, len(values))

	for i, value := range values {
		id, _ := GetFieldValue(value, "Id")
		valuesIds[i] = fmt.Sprint(id)
	}

	filter := createFilter(s.FieldName+".id", "IN", valuesIds)

	return &GLobalFilter{
		operand:      "JOIN",
		fieldFilter:  []FieldFilter{*filter},
		relationName: s.FieldName,
		joinParams:   s.JoinParams,
	}
}
