package main

import (
	"fmt"
)

type TradableScope struct {
	CompanyId int64
	Scope
}

type Scope struct {
	VehBrands []string
	Brands    []string
}

func main() {
	var t interface{}
	var b bool
	t = &b

	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}

	vehBrands := []string{"xxx"}
	brands := []string{"ddd"}
	memberId := int64(1)

	tscope := TradableScope{
		memberId,
		Scope{
			VehBrands: vehBrands,
			Brands:    brands,
		},
	}

	fmt.Printf("scope: %v, %s\n", tscope.Scope.Brands, tscope.Brands)
}
