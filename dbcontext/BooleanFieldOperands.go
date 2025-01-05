package dbcontext

// Boolean field type filters.

type BooleanFieldOperands struct {
	FieldValueOperandsParams
}

// Equals filter.
func (b *BooleanFieldOperands) Equals(value bool) *GLobalFilter {
	return &GLobalFilter{
		operand: "AND",
		fieldFilter: []FieldFilter{
			*createFilter(b.FieldName, "=", value),
			// *isNotNull(b.FieldName),
		},
	}
}

// The field is equals to value or null (undefined).
func (b *BooleanFieldOperands) EqualsOrNull(value bool) *GLobalFilter {
	return &GLobalFilter{
		operand: "OR",
		fieldFilter: []FieldFilter{
			*createFilter(b.FieldName, "=", value),
			*isNull(b.FieldName),
		},
	}
}
