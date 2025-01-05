package dbcontext

// String field type filters.

type StringFieldOperands struct {
	FieldValueOperandsParams
}

// Equals filter.
func (s *StringFieldOperands) Equals(value string) *GLobalFilter {
	globalFilter := &GLobalFilter{
		operand: "AND",
	}

	filter := createFilter(s.FieldName, "=", value)

	globalFilter.fieldFilter = append(globalFilter.fieldFilter, *filter)

	return globalFilter
}

// Searching value is on the list (value).
func (s *StringFieldOperands) In(value []string) *GLobalFilter {
	globalFilter := &GLobalFilter{
		operand: "AND",
	}

	globalFilter.fieldFilter = append(globalFilter.fieldFilter, *createFilter(s.FieldName, "in", value))

	return globalFilter
}

// Searching value contains value as part.
func (s *StringFieldOperands) Like(value string) *GLobalFilter {
	globalFilter := &GLobalFilter{
		operand: "AND",
	}

	filter := createFilter(s.FieldName, "like", "%"+value+"%")

	globalFilter.fieldFilter = append(globalFilter.fieldFilter, *filter)

	return globalFilter
}
