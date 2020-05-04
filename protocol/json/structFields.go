package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Stu struct {
	OP string
	AA int
	Ar []int8
}

type Element struct {
	AOP int `json:"aOP" comment:"哈"`
	BP  string
	CP  bool
	Stu
	Stus *Stu
}

func main() {
	strs := getSql(reflect.TypeOf(Element{}))
	fmt.Printf("{%s}\n", strings.Join(strs, " "))
}

func getSql(tpy reflect.Type) []string {
	if tpy.Kind() == reflect.Ptr {
		tpy = tpy.Elem()
	}
	if tpy.Kind() != reflect.Struct {
		return nil
	}

	strs := make([]string, 0)
	for i := 0; i < tpy.NumField(); i++ {
		field := tpy.Field(i)
		if field.Anonymous ||
			field.Type.Kind() == reflect.Ptr ||
			field.Type.Kind() == reflect.Struct {
			strs = append(strs, fmt.Sprintf("{%s}", strings.Join(getSql(field.Type), " ")))
		} else {
			fmt.Printf("%s: %+v\n", field.Name, field.Tag)
			strs = append(strs, field.Name)
		}
	}
	return strs
}
