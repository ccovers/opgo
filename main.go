package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // here be dragons
	"time"
)

func main() {
	handle(nil)
	return

	tm := time.Now()
	fmt.Println(tm.Year(), int32(tm.Month()), tm.Day())
	tm = tm.AddDate(0, 0, -0)
	fmt.Println(tm.Year(), int32(tm.Month()), tm.Day())

	return
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(stackSealMap map[int32]int32) {
	num := int32(0)
	num += stackSealMap[1]
	fmt.Printf("num: %d\n", num)
}
