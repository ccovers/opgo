package main

import (
	"fmt"
	"sort"
)

type MyDefSlice struct {
	Id   int
	Name string
}
type MySlice []MyDefSlice

func (p MySlice) Less(i, j int) bool {
	return p[i].Id > p[j].Id
}
func (p MySlice) Len() int {
	return len(p)
}
func (p MySlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	mySlice := make(MySlice, 0)
	for i := 0; i < 10; i++ {
		mySlice = append(mySlice, MyDefSlice{
			Id:   i,
			Name: fmt.Sprintf("slice_%d", i),
		})
	}
	sort.Slice(mySlice, func(i, j int) bool { return mySlice[i].Id > mySlice[j].Id })
	//sort.Sort(mySlice)

	fmt.Printf("mySlice: %v\n", mySlice)
}
