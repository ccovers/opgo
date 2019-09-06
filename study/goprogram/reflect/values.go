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
	Score     int
	Books     []string
	Subs      []MyObj
	NationMap map[string][]string
	A
	B A
	int
	INT
	MyObj *MyObj
}

func marshal(field reflect.Value, anonymous bool) string {
	if field.Kind() == reflect.Ptr && !field.IsNil() {
		field = field.Elem()
	}

	columns := make([]string, 0)
	switch field.Kind() {
	case reflect.Struct:
		for i := 0; i < field.NumField(); i++ {
			name := field.Type().Field(i).Name
			if name[0] >= 'A' && name[0] <= 'Z' {
				if field.Type().Field(i).Anonymous && field.Field(i).Kind() == reflect.Struct {
					columns = append(columns, marshal(field.Field(i), true))
				} else {
					columns = append(columns,
						fmt.Sprintf(`%v`, marshal(field.Field(i), false)))
				}
			}
		}
		if anonymous {
			return fmt.Sprintf("%s", strings.Join(columns, ","))
		} else {
			return fmt.Sprintf("{%s}", strings.Join(columns, ","))
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < field.Len(); i++ {
			columns = append(columns, marshal(field.Index(i), false))
		}
		if len(columns) > 0 {
			return fmt.Sprintf(`[%s]`, strings.Join(columns, ","))
		} else {
			return fmt.Sprintf(`null`)
		}
	case reflect.Map:
		for _, key := range field.MapKeys() {
			columns = append(columns, fmt.Sprintf(`"%s":%s`, key, marshal(field.MapIndex(key), false)))
		}
		if len(columns) > 0 {
			return fmt.Sprintf(`{%s}`, strings.Join(columns, ","))
		} else {
			return fmt.Sprintf(`null`)
		}
	case reflect.Invalid:
		panic(fmt.Errorf("invalid reflect type: %v", field))
	case reflect.Ptr:
		return fmt.Sprintf(`null`)
	case reflect.String:
		return fmt.Sprintf(`"%v"`, field)
	default:
		return fmt.Sprintf(`%v`, field)
	}
}

func Marshal(obj interface{}) string {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				fmt.Println("panic:", err)
			} else {
				panic(r)
			}
		}
	}()

	field := reflect.ValueOf(obj)

	if field.Kind() == reflect.Ptr {
		field = field.Elem()
	}
	if field.Kind() != reflect.Struct {
		return ""
	} else {
		return marshal(field, false)
	}
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
			//NationMap: map[string][]string{"体育": []string{"篮球", "足球"}},
		}},
		NationMap: map[string][]string{"体育": []string{"篮球", "足球"}, "学习": []string{"科学"}},
	}
	fmt.Println(Marshal(&obj))

	bt, _ := json.Marshal(&obj)
	fmt.Println(string(bt))
}
