package main

import (
    "fmt"
    "os"
    "reflect"
)

type TInfo struct {
    Name  string
    Score int32
}

type TBase struct {
    Name string
    Gard string
}

type TGrid struct {
    X       int32
    Y       int32
    NMap    map[int32]*TBase
    Array   []int32
    Unkonwn interface{}
}

func main() {
    grid := &TGrid{
        X:       1,
        Y:       1,
        NMap:    map[int32]*TBase{1: &TBase{Name: "xx", Gard: "yy"}, 2: &TBase{Name: "xc", Gard: "yc"}},
        Array:   []int32{1, 2},
        Unkonwn: &TInfo{Name: "honghong", Score: 100},
    }
    Display("grid", grid)
    Display("os.Stderr", os.Stderr)
}

func Display(name string, x interface{}) {
    fmt.Printf("Display %s (%T):\n", name, x)
    display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
    switch v.Kind() {
    case reflect.Invalid:
        fmt.Printf("%s = invalid\n", path)
    case reflect.Slice, reflect.Array:
        for i := 0; i < v.Len(); i++ {
            display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
        }
    case reflect.Struct:
        for i := 0; i < v.NumField(); i++ {
            display(fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name), v.Field(i))
        }
    case reflect.Map:
        for _, key := range v.MapKeys() {
            display(fmt.Sprintf("%s[%v]", path, key), v.MapIndex(key))
        }
    case reflect.Ptr:
        if v.IsNil() {
            fmt.Printf("%s = nil\n", path)
        } else {
            display(fmt.Sprintf("(*%s)", path), v.Elem())
        }
    case reflect.Interface:
        if v.IsNil() {
            fmt.Printf("%s = nil\n", path)
        } else {
            fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
            display(fmt.Sprintf("%s.value", path), v.Elem())
        }
    default:
        fmt.Printf("%s = %v\n", path, v)

    }
}
