package main

import (
	"fmt"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type MyStruct struct {
	Id int
}

func (MyStruct) Read(p []byte) (n int, err error) {
	fmt.Printf("read: %s\n", string(p))
	return 0, nil
}

func (MyStruct) Write(p []byte) (n int, err error) {
	fmt.Printf("write: %s\n", string(p))
	return 0, nil
}

func DoSome(obj ReadWriter) {
	obj.Read([]byte("==="))
	obj.Write([]byte("==="))
}

type Obj struct {
	Id int
	*MyStruct
}

func main() {
	mystruct := MyStruct{}
	DoSome(mystruct)

	obj := Obj{
		1,
		&MyStruct{
			2,
		},
	}
	fmt.Printf("%d, %d\n", obj.MyStruct.Id, obj.Id)

}
