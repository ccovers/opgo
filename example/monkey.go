package main

import (
	"bou.ke/monkey"
	"fmt"
	"work/opgo/example/xpp"
)

func main() {

	p := monkey.Patch(xpp.GetName, func(id int) string {
		return "替换值"
	})
	defer p.Unpatch()

	test()
}

func test() {
	fmt.Println(xpp.GetName(1))
}
