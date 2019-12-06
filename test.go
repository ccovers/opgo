package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/lovego/pinyin"
)

type Info struct {
	Id int64
}

func main() {
	fmt.Printf("aa09.;p考试\n%s\n", pinyin.Initials("aa09.;p考试"))

	fmt.Println(time.Now().Format("2006-01-02 15:04"))
	fmt.Println(time.ParseInLocation("2006-01-02 15:04", time.Now().Format("2006-01-02 15:04"), time.Local))

	name := "好了\\kk%"
	fmt.Println(strings.Replace(name, "\\", "\\\\", -1))
	fmt.Println(strings.Replace(name, "%", "\\%", -1))

	infoMap := make(map[int64]*Info)
	info := Info{Id: 100}
	infoMap[info.Id] = &info
	v, ok := infoMap[100]
	fmt.Printf("%+v, %+v\n", v, ok)
	v, ok = infoMap[1]
	fmt.Printf("%+v, %+v\n", v, ok)
}
