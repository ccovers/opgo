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

type TKey struct {
    Id   int32
    Name string
}

type TGrid struct {
    X       int32
    Y       int32
    NMap    map[TKey]*TBase
    Array   []int32
    Unkonwn interface{}
}

func main() {
    grid := &TGrid{
        X:       1,
        Y:       1,
        NMap:    map[TKey]*TBase{TKey{Id: 1, Name: "111"}: &TBase{Name: "xx", Gard: "yy"}, TKey{Id: 2, Name: "222"}: &TBase{Name: "xc", Gard: "yc"}},
        Array:   []int32{1, 2},
        Unkonwn: &TInfo{Name: "honghong", Score: 100},
    }
    Display("grid", grid)
    Display("os.Stderr", os.Stderr)
}

func Display(name string, x interface{}) {
    fmt.Printf("Display %s (%T):\n", name, x)
    cycle := make(map[string]struct{})
    display(name, reflect.ValueOf(x), cycle)
}

func display(path string, v reflect.Value, cycle map[string]struct{}) {
    switch v.Kind() {
    case reflect.Invalid:
        fmt.Printf("%s = invalid\n", path)
    case reflect.Slice, reflect.Array:
        for i := 0; i < v.Len(); i++ {
            display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), cycle)
        }
    case reflect.Struct:
        for i := 0; i < v.NumField(); i++ {
            if _, ok := cycle[v.Type().Field(i).Name]; ok {
                continue
            }
            cycle[v.Type().Field(i).Name] = struct{}{}
            display(fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name), v.Field(i), cycle)
        }
    case reflect.Map:
        for _, key := range v.MapKeys() {
            display(fmt.Sprintf("%s[%v]", path, key), v.MapIndex(key), cycle)
        }
    case reflect.Ptr:
        if v.IsNil() {
            fmt.Printf("%s = nil\n", path)
        } else {
            display(fmt.Sprintf("(*%s)", path), v.Elem(), cycle)
        }
    case reflect.Interface:
        if v.IsNil() {
            fmt.Printf("%s = nil\n", path)
        } else {
            fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
            display(fmt.Sprintf("%s.value", path), v.Elem(), cycle)
        }
    default:
        fmt.Printf("%s = %v\n", path, v)
    }
}
