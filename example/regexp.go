package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "90jewadksackfi"
	b, err := regexp.Match("sa", []byte(str))
	if err != nil {
		fmt.Println("err:", b)
		return
	}

	fmt.Println(b)
}
