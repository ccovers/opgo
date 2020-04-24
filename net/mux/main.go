package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	s = s.PathPrefix("/user").Subrouter()
	s.HandleFunc("/list", UserListHandler).Methods("POST")
	s.HandleFunc("/add", UserAddHandler).Methods("GET")
	s.HandleFunc("/orders/{orderId}:insure", UserAddHandler).Methods("POST")

	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	fmt.Println("dir:", dir)
	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		fmt.Println(methods, pathTemplate)
		return nil
	})
	if err != nil {
		log.Fatal("walk: ", err)
	}

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user list"))
}

func UserAddHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user add"))
}
