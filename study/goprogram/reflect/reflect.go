package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

/*
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
			if value, ok := struct_tag.Lookup(string(field.Tag), `sql`); !ok || value != "-" {
				fn(field)
			}
		}
	}
	return true*/
func Sprint(x interface{}) string {
	typ := reflect.ValueOf(x)
	columns := make([]string, 0)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Struct {
		for i := 0; i < typ.NumField(); i++ {
			name := typ.Type().Field(i).Name
			field := typ.Field(i)

			if field.Kind() == reflect.Struct {
				columns = append(columns, fmt.Sprintf(`"%s":%v`, name, Sprint(field)))
			} else if field.Kind() == reflect.Array || field.Kind() == reflect.Slice {
				for k := 0; k < field.Len(); k++ {
					fmt.Println(field.Type().Field(i).Name)
					//columns = append(columns, fmt.Sprintf(`"%s":%v`, name, Sprint(field)))
				}
			} else {
				columns = append(columns, fmt.Sprintf(`"%s":%v`, name, field))
			}

		}
	}
	return fmt.Sprintf("{%s}", strings.Join(columns, ","))
}

type MyObj struct {
	Name  string `json:"name" comment:"名称"`
	Score int
	Books []string
	Subs  []MyObj
}

func main() {
	obj := MyObj{
		Name:  "小明",
		Score: 90,
		Books: []string{"语文", "数学"},
		Subs: []MyObj{MyObj{
			Name:  "小东",
			Score: 80,
			Books: []string{"语文", "数学"},
			Subs:  nil,
		}},
	}
	fmt.Println(Sprint(obj))

	bt, _ := json.Marshal(&obj)
	fmt.Println(string(bt))
}
