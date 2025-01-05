package dbcontext

import "time"

// Date/time fields filters

type DateTimeFieldOperands struct {
	FieldValueOperandsParams
}

// The values is equals to field record.
func (s *DateTimeFieldOperands) Equals(value time.Time) *GLobalFilter {
	return &GLobalFilter{
		operand:     "AND",
		fieldFilter: []FieldFilter{*createFilter(s.FieldName, "=", value)},
	}
}

// Filter records where value less then inputted value.
func (s *DateTimeFieldOperands) Before(value time.Time) *GLobalFilter {
	return &GLobalFilter{
		operand:     "AND",
		fieldFilter: []FieldFilter{*createFilter(s.FieldName, "<", value)},
	}
}

// Filter records where value less then inputted value.
func (s *DateTimeFieldOperands) After(value time.Time) *GLobalFilter {
	return &GLobalFilter{
		operand:     "AND",
		fieldFilter: []FieldFilter{*createFilter(s.FieldName, ">", value)},
	}
}
