package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Class struct {
	ID    int32
	Score int32
}
type MyObj struct {
	Name      string `json:"name" comment:"名称"`
	Score     int32  //`json:"score" comment:"分数"`
	Books     []string
	NationMap map[string][]string
	Class
	ClassT     Class
	ClassP     *Class
	SubClasses []*Class
	Unkown     interface{}
}

func display(field reflect.Value, anonymous bool, nMap map[string]struct{}) string {
	switch field.Kind() {
	case reflect.Struct:
		nmap := map[string]struct{}{}
		for i := 0; i < field.NumField(); i++ {
			name := field.Type().Field(i).Name
			if tname := field.Type().Field(i).Tag.Get("json"); tname != "" {
				name = tname
			}
			nmap[name] = struct{}{}
		}
		columns := make([]string, 0)
		for i := 0; i < field.NumField(); i++ {
			name := field.Type().Field(i).Name
			if name[0] >= 'A' && name[0] <= 'Z' {
				if tname := field.Type().Field(i).Tag.Get("json"); tname != "" {
					name = tname
				}
				if _, ok := nMap[name]; ok {
					continue
				}
				if field.Type().Field(i).Anonymous && field.Field(i).Kind() == reflect.Struct {
					columns = append(columns, display(field.Field(i), true, nmap))
				} else {
					columns = append(columns, fmt.Sprintf(`"%s":%s`, name, display(field.Field(i), false, nil)))
				}
			}
		}
		if anonymous {
			return fmt.Sprintf("%s", strings.Join(columns, ","))
		}
		return fmt.Sprintf("{%s}", strings.Join(columns, ","))
	case reflect.Array, reflect.Slice:
		columns := make([]string, 0)
		for i := 0; i < field.Len(); i++ {
			columns = append(columns, display(field.Index(i), false, nil))
		}
		if len(columns) == 0 {
			return fmt.Sprintf(`null`)
		}
		return fmt.Sprintf(`[%s]`, strings.Join(columns, ","))
	case reflect.Map:
		columns := make([]string, 0)
		for _, key := range field.MapKeys() {
			columns = append(columns, fmt.Sprintf(`"%s":%s`, key, display(field.MapIndex(key), false, nil)))
		}
		if len(columns) == 0 {
			return fmt.Sprintf(`null`)
		}
		return fmt.Sprintf(`{%s}`, strings.Join(columns, ","))
	case reflect.Ptr, reflect.Interface:
		if field.IsNil() {
			return fmt.Sprintf("null")
		}
		return display(field.Elem(), false, nil)
	case reflect.Invalid:
		panic(fmt.Errorf("invalid reflect type: %v", field))
	case reflect.String:
		return fmt.Sprintf(`"%v"`, field)
	default:
		return fmt.Sprintf(`%v`, field)
	}
}

func Display(obj interface{}) string {
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
		fmt.Println("...")
		return ""
	}
	return display(field, false, nil)
}

func main() {

	obj := MyObj{
		Name:       "小明",
		Books:      []string{"语文", "数学"},
		NationMap:  map[string][]string{"体育": []string{"篮球", "足球"}, "学习": []string{"科学"}},
		Class:      Class{ID: 1, Score: 1},
		ClassT:     Class{ID: 2, Score: 2},
		ClassP:     &Class{ID: 3, Score: 3},
		SubClasses: []*Class{&Class{ID: 4, Score: 4}, &Class{ID: 5, Score: 5}},
		Unkown:     &Class{ID: 6, Score: 6},
	}
	fmt.Println("display: ", Display(&obj))

	bt, _ := json.Marshal(&obj)
	fmt.Println("marshal: ", string(bt))
}
