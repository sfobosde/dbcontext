package dbcontext

import (
	"fmt"
	"reflect"
)

// Получение строки подключения к БД.
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

// Получение тега поля.
func GetFieldTag(field reflect.StructField, tagKey string) (string, bool) {
	tag := field.Tag.Get(tagKey)
	return tag, tag != ""
}

// Простановка кодов полям из SearchFields объекта.
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

// Установка значению полю по указанному коду у переданного объекта.
func SetObjectProperty[TObject any](object *TObject, fieldName string, value string) error {
	v := reflect.ValueOf(object).Elem()
	f := v.FieldByName(fieldName)

	if !f.IsValid() {
		return fmt.Errorf("поле '%s' не найдено", fieldName)
	}

	f.SetString(value)

	return nil
}
