package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func Marshal(typ reflect.Value) string {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Struct {
		columns := make([]string, 0)
		for i := 0; i < typ.NumField(); i++ {
			name := typ.Type().Field(i).Name

			columns = append(columns, fmt.Sprintf(`"%s":%v`, name, Marshal(typ.Field(i))))
		}
		return fmt.Sprintf("{%s}", strings.Join(columns, ","))
	} else if typ.Kind() == reflect.Array || typ.Kind() == reflect.Slice {
		columns := make([]string, 0)
		for i := 0; i < typ.Len(); i++ {
			columns = append(columns, Marshal(typ.Index(i)))
		}
		return fmt.Sprintf(`[%s]`, strings.Join(columns, ","))
	} else if typ.Kind() == reflect.Map {
		columns := make([]string, 0)
		for _, key := range typ.MapKeys() {
			columns = append(columns, fmt.Sprintf(`"%s":%s`, key, Marshal(typ.MapIndex(key))))
		}
		return fmt.Sprintf(`[%s]`, strings.Join(columns, ","))
	} else {
		if typ.Kind() == reflect.String {
			return fmt.Sprintf(`"%v"`, typ)
		} else {
			return fmt.Sprintf(`%v`, typ)
		}
	}
}

type MyObj struct {
	Name      string `json:"name" comment:"名称"`
	Score     int
	Books     []string
	Subs      []MyObj
	NationMap map[string][]string
}

func main() {
	obj := MyObj{
		Name:  "小明",
		Score: 90,
		Books: []string{"语文", "数学"},
		Subs: []MyObj{MyObj{
			Name:      "小东",
			Score:     80,
			Books:     []string{"语文", "数学"},
			Subs:      nil,
			NationMap: map[string][]string{"体育": []string{"篮球", "足球"}},
		}},
		NationMap: map[string][]string{"体育": []string{"篮球", "足球"}, "学习": []string{"科学", "政治"}},
	}
	fmt.Println(Marshal(reflect.ValueOf(obj)))

	bt, _ := json.Marshal(&obj)
	fmt.Println(string(bt))
}
