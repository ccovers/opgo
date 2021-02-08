package main

import (
	"fmt"
)

func main() {
	mySlice := make([]byte, 0, 5)
	data := []byte{1, 2, 3, 4, 5}

	mySlice = Append(mySlice, data)
	fmt.Printf("slice: %v\n", mySlice)
}

func Append(slice []byte, data []byte) []byte {
	sliLen := len(slice)
	needLen := len(slice) + len(data)

	if needLen > cap(slice) {
		//创建长度足够的新切片
		newSlice := make([]byte, needLen, needLen)
		copy(newSlice, slice)
		slice = newSlice
	}

	//重置切片长度---切片长度可能小于容量
	//slice = slice[0:needLen]

	for i, c := range data {
		slice[sliLen+i] = c
	}
	return slice
}
