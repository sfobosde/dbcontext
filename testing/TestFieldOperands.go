package dbcontext_test

// Тest FieldOperands filters.
func testFieldOperands(models *model) {
	testStringFieldOperands(models)
	testDateTimeFieldOperands(models)
	testBoolFieldOperands(models)
}
