package main

import (
	"fmt"
	"net/url"
)

func main() {
	brandMap := make(map[string][]int)
	brands, ok := brandMap["cookie"]
	if !ok {
		brandMap["cookie"] = make([]int, 0)
		brands, _ = brandMap["cookie"]
	}
	bs := make([]int, 10)
	ap(bs)
	brands = append(brands, 1)
	brands = append(brands, 2)
	brandMap["cookie"] = append(brandMap["cookie"], 12)

	fmt.Printf("%+v\n%+v\n%+v\n", brands, bs, brandMap)
}

func ap(bs []int) {
	bs[1] = 10
}
