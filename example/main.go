package main

import (
	"fmt"
	"sync"
)

var number int = 10
var mutex sync.Mutex

func main() {
	chs := make(chan []string, 1)
	//go func() {
	chs <- []string{}
	//}()
	fmt.Println("len:", len(chs))
	fmt.Println("cap:", cap(chs))

	var n int
	n++
	for ; n > 0; n-- {
		for list := range chs {
			for _, link := range list {
				fmt.Println(link)
				n++
				go func() {
					m := getNumber()
					if m > 0 {
						chs <- []string{fmt.Sprintf("%d", m)}
					} else {
						fmt.Println("xxx")
						chs <- []string{}
					}
				}()
			}
		}
	}
}

func getNumber() int {
	mutex.Lock()
	defer mutex.Unlock()
	if number > 0 {
		number--
	}
	return number
}
