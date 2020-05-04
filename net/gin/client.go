package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	//"net/url"
	"strings"
)

func main() {
	// Get
	resp, err := http.Get("http://0.0.0.0:8080/home/user/xiaoming/palying")
	if err != nil {
		log.Fatalf("/home/user: %s\n", err.Error())
	}
	if printData(resp) != nil {
		return
	}

	// POST
	/*url := url.Values{}
	url.Set("msg", "paly")
	resp, err = http.PostForm("http://0.0.0.0:8080/home/msg", url)
	*/
	resp, err = http.Post("http://0.0.0.0:8080/home/msg", "application/x-www-form-urlencoded", strings.NewReader("msg=play"))
	if err != nil {
		log.Fatalf("/home/msg: %s\n", err.Error())
	}
	if printData(resp) != nil {
		return
	}
	resp, err = http.Post("http://0.0.0.0:8080/home/data", "application/json",
		strings.NewReader("{\"id\":890}"))
	if err != nil {
		log.Fatalf("/home/data: %s\n", err.Error())
	}
	if printData(resp) != nil {
		return
	}
}

func printData(resp *http.Response) error {
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		log.Fatalf("read: %s\n", err.Error())
		return err
	}
	fmt.Printf("%s: [%s - %d] %s\n", resp.Request.URL, resp.Proto, resp.StatusCode, string(data))
	return nil
}
