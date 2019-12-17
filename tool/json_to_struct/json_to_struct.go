package json_to_struct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type fieldT struct {
	name    string
	jsonKey string
}

func Affected(data []byte, p interface{}) ([]string, error) {
	// json字符串字段
	fieldMap := make(map[string]interface{})
	err := json.Unmarshal(data, &fieldMap)
	if err != nil {
		return nil, err
	}

	// 结构体中的字段
	fields, err := getFields(p)
	if err != nil {
		return nil, err
	}

	// 对比相同的字段
	result := make([]string, 0)
	for _, field := range fields {
		_, ok := fieldMap[field.jsonKey]
		if ok {
			result = append(result, field.name)
		}
	}
	return result, nil
}

func getFields(p interface{}) ([]fieldT, error) {
	typ := reflect.TypeOf(p)
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	fields, err := newFields(typ)
	if err != nil {
		return nil, err
	}
	return fields, nil
}

func newFields(typ reflect.Type) ([]fieldT, error) {
	var fields = make([]fieldT, 0, typ.NumField())
	var m1 = make(map[string][]fieldT)
	var m2 = make(map[string][]fieldT)

	traverseStructFields(typ, func(field reflect.StructField) {
		key := getJSONKey(field.Name, field.Tag.Get("json"))
		if key != "" {
			lower := strings.ToLower(key)
			fields = append(fields, fieldT{name: field.Name, jsonKey: lower})
			m1[lower] = append(m1[lower], fieldT{name: field.Name, jsonKey: key})
			m2[field.Name] = append(m2[field.Name], fieldT{name: field.Name, jsonKey: key})
		}
	})

	for _, conflicts := range m1 {
		if len(conflicts) > 1 {
			return nil, fmt.Errorf("conflicts field jsonKey: %+v", conflicts)
		}
	}
	for _, conflicts := range m2 {
		if len(conflicts) > 1 {
			return nil, fmt.Errorf("conflicts field name: %+v", conflicts)
		}
	}

	return fields, nil
}

func getJSONKey(fieldName, tag string) string {
	if tag == "-" {
		return ""
	}
	if idx := strings.Index(tag, ","); idx != -1 {
		tag = tag[:idx]
	}
	if tag == "" {
		return fieldName
	}
	return tag
}

func traverseStructFields(typ reflect.Type, fn func(field reflect.StructField)) bool {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if (!field.Anonymous || !traverseStructFields(field.Type, fn)) &&
			(field.Name[0] >= 'A' && field.Name[0] <= 'Z') {
			fn(field)
		}
	}
	return true
}
