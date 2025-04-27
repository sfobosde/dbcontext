package dbcontext

import "fmt"

type LinkedObjectFieldOperands[TModel interface{}, TSearch interface{}] struct {
	FieldValueOperandsParams
	JoinParams string
}

func (s *LinkedObjectFieldOperands[TModel, TSearchFields]) Has(values ...TModel) *GLobalFilter {
	valuesIds := make([]string, len(values))

	for i, value := range values {
		id, _ := GetFieldValue(value, "ID")
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
