package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // here be dragons
	// _ "runtime/pprof"
)

// http://localhost:8080/debug/pprof/
// go tool pprof -text -nodecount=20 . profile
// go tool pprof http://localhost:8080/debug/pprof/profile?seconds=60
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
