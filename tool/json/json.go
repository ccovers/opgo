package main

import (
	"encoding/json"
	"fmt"
)

var jsonChar string = `{"id":1, "name":"小明", "age":12}`

// 打印 json 字符串的字段
func main() {
	columnMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonChar), &columnMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, v := range columnMap {
		fmt.Printf("%s: %v\n", key, v)
	}
}

/*
id: 1
name: 小明
age: 12
*/
