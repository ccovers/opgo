package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"./tool/json_to_struct"
)

// var data string = `[{"name":"小明", "age":12, "sub":"英语","score":100},{"name":"小红", "age":12, "sub":"英语","score":100}]`
var data string = `{"name":"小明", "age":12, "sub":"英语","score":100,"home":"china"}`

type Column struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Sub     string `json:"sub"`
	Score   int    `json:"score"`
	Company string `json:"company"`
}

// 打印 json 字符串的字段
func main() {
	column := Column{}
	err := json.Unmarshal([]byte(data), &column)
	if err != nil {
		fmt.Println(err)
		return
	}

	fields, err := json_to_struct.Affected([]byte(data), Column{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", column)
	fmt.Printf("%s\n", strings.Join(fields, ","))
}

/*
id: 1
name: 小明
age: 12
*/
