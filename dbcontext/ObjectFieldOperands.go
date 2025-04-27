package dbcontext

type ObjectFieldOperands[TModel any] struct {
	FieldValueOperandsParams
}

func (s *ObjectFieldOperands[TModel]) Link(values ...*TModel) *GLobalFilter {
	var ids []any
	for _, value := range values {
		id, _ := GetFieldValue(*value, "ID")
		ids = append(ids, id)
	}

	return &GLobalFilter{
		operand:     "OR",
		fieldFilter: []FieldFilter{*createFilter(s.FieldName, "in", ids)},
	}
}
