package dbcontext

type FieldValueOperandsParams struct {
	// Field db code-name.
	FieldName string
}

// Field filter structure.
type FieldFilter struct {
	fieldName string
	operator  string
	value     interface{}
}

// Filter stucture constructor.
func createFilter[T any](fieldName, operator string, value T) *FieldFilter {
	filter := new(FieldFilter)
	filter.fieldName = fieldName
	filter.operator = operator
	filter.value = value

	return filter
}

// Field is empty.
func isNull(FieldName string) *FieldFilter {
	return &FieldFilter{
		fieldName: FieldName,
		operator:  "is",
		value:     nil,
	}
}
