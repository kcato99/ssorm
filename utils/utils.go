package utils

import (
	"reflect"
)

func GetTableName(model interface{}) string {
	results := reflect.Indirect(reflect.ValueOf(model))

	if reflect.TypeOf(model).Kind() == reflect.String {
		return model.(string)
	}

	if kind := results.Kind(); kind == reflect.Slice {
		resultType := results.Type().Elem()
		if resultType.Kind() == reflect.Ptr {
			resultType = resultType.Elem()
			elem := reflect.New(resultType).Interface()
			return reflect.TypeOf(elem).Elem().Name()
		} else {
			return resultType.Name()
		}
	}

	if reflect.TypeOf(model).Kind() == reflect.Ptr {
		modelType := reflect.TypeOf(model)
		modelValue := reflect.New(modelType.Elem()).Interface()
		return reflect.TypeOf(modelValue).Elem().Name()
	}

	return results.Type().Name()
}

func ArrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func ReflectValues(reflectValue reflect.Value, i int) (reflect.StructTag, string, interface{}, reflect.Type) {
	value := reflect.TypeOf(reflectValue.Interface())
	tag := value.Field(i).Tag
	varName := reflectValue.Type().Field(i).Name
	varType := reflectValue.Type().Field(i).Type
	varValue := reflectValue.Field(i).Interface()
	return tag, varName, varValue, varType
}
