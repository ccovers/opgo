package main

import (
	"fmt"
	"net/http"
	"os"
)

type Counter int

func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	*c++
	fmt.Fprintf(w, "counter=%d\n", *c)
}

func ArgServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, os.Args)
}

func main() {
	c := new(Counter)
	http.Handle("/counter", c)
	//http.ListenAndServe(":3000", http.HandlerFunc(\))
	http.ListenAndServe(":3000", c)
}
