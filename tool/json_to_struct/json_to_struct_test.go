package json_to_struct

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Column struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Sub     string `json:"sub"`
	Score   int    `json:"score"`
	Company string `json:"company"`
}

// var data string = `[{"name":"小明", "age":12, "sub":"英语","score":100},{"name":"小红", "age":12, "sub":"英语","score":100}]`
var data string = `{"name":"小明", "age":12, "sub":"英语","score":100,"home":"china"}`

func ExampleJsonToStruct() {
	column := Column{}
	err := json.Unmarshal([]byte(data), &column)
	if err != nil {
		fmt.Println(err)
		return
	}

	fields, err := Affected([]byte(data), Column{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", column)
	fmt.Printf("%s\n", strings.Join(fields, ","))
	// Output:
	// {Name:小明 Age:12 Sub:英语 Score:100 Company:}
	// Name,Age,Sub,Score
}
