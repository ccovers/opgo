package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type INT int
type A struct {
	OP string
}
type MyObj struct {
	Name      string `json:"name" comment:"名称"`
	Score     int    `json:"-"`
	Books     []string
	Subs      []MyObj
	NationMap map[string][]string
	A
	Bb_BB_bB A
	int
	INT
	MyObj *MyObj
}

func fieldsFromStruct(typ reflect.Type, fn func(reflect.StructField)) {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Anonymous && typ.Field(i).Type.Kind() == reflect.Struct {
			fieldsFromStruct(typ.Field(i).Type, fn)
		} else {
			fn(typ.Field(i))
		}
	}
}

func FieldsFromStruct(obj interface{}) []string {
	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	columns := make([]string, 0)
	if typ.Kind() == reflect.Struct {
		fieldsFromStruct(typ, func(field reflect.StructField) {
			tag, _ := field.Tag.Lookup("json")
			if field.Name[0] >= 'A' && field.Name[0] <= 'Z' && tag != "-" {
				columns = append(columns, field.Name)
			}
		})
	}
	return columns
}

/* 单词边界有两种
1. 非大写字符，且下一个是大写字符
2. 大写字符，且下一个是非大写字符，且上一个是大写字符，
*/
func ColumnsFromFields(fields []string) []string {
	column := func(str string) string {
		columns := make([]string, 0)
		start := 0
		for end, ch := range str {
			if end+1 < len(str) && ch != '_' {
				next := str[end+1]
				if ch < 'A' || ch > 'Z' {
					if next >= 'A' && next <= 'Z' {
						columns = append(columns, strings.ToLower(str[start:end+1]))
						start = end + 1
					}
				} else {
					if next < 'A' || next > 'Z' {
						if start < end && next != '_' {
							columns = append(columns, strings.ToLower(str[start:end]))
							start = end
						}
					}
				}
			}
		}
		columns = append(columns, strings.ToLower(str[start:]))
		return strings.Join(columns, "_")
	}

	for i, _ := range fields {
		fields[i] = column(fields[i])
	}
	return fields
}

func main() {
	fmt.Println(strings.Join(ColumnsFromFields(FieldsFromStruct(&MyObj{})), ","))

	bt, _ := json.Marshal(&MyObj{})
	fmt.Println(string(bt))
}
