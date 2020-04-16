package main

import (
	"fmt"
	"github.com/httprouter"
	"net/http"
	"sync"
	// "time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		http_router()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		http_server()
	}()

	/*ch := make(<-chan int)
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("===")
			case <-time.After(1 * time.Second):
				fmt.Println("second")
			}
		}
	}()*/

	wg.Wait()
	fmt.Printf("listen over!\n")
}

type Handler struct {
}

func http_server() {
	http.ListenAndServe(":8081", &Handler{})
}
func (handler *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		resp.Write([]byte("welcome! (server)\n"))
	} else if req.Method == "GET" {
		if len(req.URL.Path) > 4 && req.URL.Path[0:4] == "/src" {
			req.URL.Path = req.URL.Path[4:]
			fileServer := http.FileServer(http.Dir("./"))
			fileServer.ServeHTTP(resp, req)
		} else {
			resp.Write([]byte("bad http method request\n"))
		}
	} else {
		resp.Write([]byte("bad http method request\n"))
	}
}

func http_router() {
	router := httprouter.New()
	router.GET("/", indexRouter)
	router.GET("/hello/:name", helloRouter)
	router.ServeFiles("/src/*filepath", http.Dir("/"))

	router.PanicHandler = func(resp http.ResponseWriter, req *http.Request, data interface{}) {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(resp, "error:", data)
	}
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("listen err: %s\n", err.Error())
	}
	fmt.Println("http_router over")
}

func indexRouter(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	panic("故意抛出的异常")
	// resp.Write([]byte("welcome! (router)\n"))
}

func helloRouter(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	resp.Write([]byte(fmt.Sprintf("hello %s (router)\n", params.ByName("name"))))
}
