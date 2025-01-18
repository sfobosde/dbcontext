package dbcontext

import (
	"fmt"
	"reflect"
)

// Get db connection string.
func GetConnectionConfig(properties *ConnectionProperties) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		properties.Host,
		properties.User,
		properties.Password,
		properties.DBName,
		properties.Port,
	)
}

// Get field tag value.
func GetFieldTag(field reflect.StructField, tagKey string) (string, bool) {
	tag := field.Tag.Get(tagKey)
	return tag, tag != ""
}

// Set SearchField object fields values by tag value.
func SetFieldCodes[TSearch any](searchFields *TSearch) {
	v := reflect.ValueOf(searchFields)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag, ok := GetFieldTag(field, "dbcontext")
		if ok {
			operand := reflect.New(field.Type.Elem()).Interface()
			v.Field(i).Set(reflect.ValueOf(operand))
			operandValue := reflect.ValueOf(operand).Elem()
			fieldCodeField := operandValue.FieldByName("FieldName")

			if fieldCodeField.IsValid() {
				fieldCodeField.SetString(tag)
			}
		}
	}
}

// Set object property value by field code.
func SetObjectProperty[TObject any](object *TObject, fieldName string, value string) error {
	v := reflect.ValueOf(object).Elem()
	f := v.FieldByName(fieldName)

	if !f.IsValid() {
		return fmt.Errorf("поле '%s' не найдено", fieldName)
	}

	f.SetString(value)

	return nil
}

// Get object property value by field code.
func GetFieldValue(s interface{}, fieldName string) (interface{}, error) {
	val := reflect.ValueOf(s)

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %s", val.Kind())
	}

	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, fmt.Errorf("field '%s' not found", fieldName)
	}

	return field.Interface(), nil
}
