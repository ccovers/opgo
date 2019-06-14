package main

import (
	"fmt"
	"reflect"
)

type Detail struct {
	Name   string
	Amount int64
}

type Obj struct {
	Id   int64
	Name string
	Detail
	DetailP Detail
	Details []Detail
	FID     []int64
}

func main() {
	obj := Obj{
		Id:      int64(100),
		Name:    "xiaoming",
		Detail:  Detail{},
		DetailP: Detail{},
		Details: []Detail{},
	}

	fmt.Printf("%s\n", jsonConvert(obj))
}

// 提供`interface{}`对象，将其内容转换为json格式数据`{"a":1,"b":"xx","c":[1,2,3]}`
func jsonConvert(jsonObj interface{}) string {
	otv := reflect.ValueOf(jsonObj)
	switch otv.Kind() {
	case reflect.Struct:
		var tsr string
		for i := 0; i < otv.NumField(); i++ {
			if len(tsr) > 0 {
				tsr += ","
			}

			tsr += fmt.Sprintf("\"%s\":%s", reflect.TypeOf(jsonObj).Field(i).Name,
				jsonConvert(otv.Field(i).Interface()))
		}
		return fmt.Sprintf("{%s}", tsr)
	case reflect.Slice, reflect.Array:
		var tsr string
		for i := 0; i < otv.Len(); i++ {
			if len(tsr) > 0 {
				tsr += ","
			}
			tsr += jsonConvert(otv.Index(i).Interface())
		}
		return fmt.Sprintf("[%s]", tsr)
	case reflect.String:
		return fmt.Sprintf("\"%v\"", otv)
	default:
		return fmt.Sprintf("%v", otv)
	}
}
