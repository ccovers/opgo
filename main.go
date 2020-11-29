package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		fmt.Println("hello")
	})
	http.Handle("/", r)

	ch := make(chan int, 0)
	go func() {
		cnt := 0
		timer := time.NewTimer(2 * time.Second)
		for {
			select {
			case <-timer.C:
				fmt.Println("over")
			default:
				fmt.Println("xxx")
				time.Sleep(1 * time.Second)
				cnt++
				if cnt > 3 {
					ch <- 0
					return
				}
				timer.Reset(1 * time.Second)
			}
		}
	}()
	x := <-ch
	fmt.Println("...", x)
}
