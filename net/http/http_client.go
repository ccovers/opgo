package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Req struct {
	Id   int
	Name string
}

type Res struct {
	Status int
}

func main() {
	res := Res{}
	err := httpSend(&Req{
		Id:   1024,
		Name: "xiaoming",
	}, &res, "xxx000")
	if err != nil {
		return
	}
	fmt.Printf("res: %+v\n", res)
}

func httpSend(req interface{}, res interface{}, token string) error {
	buffer, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Marshal", err)
		return err
	}
	httpRequest, err := http.NewRequest("POST",
		"http://127.0.0.1:8080/detail?brand=xx",
		bytes.NewBuffer(buffer))
	if err != nil {
		fmt.Println("NewRequest", err)
		return err
	}
	httpRequest.Header.Set("Cookie", token)
	//httpRequest.Header.Add(key, value)

	httpTransport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{
		Transport: httpTransport,
	}
	resp, err := client.Do(httpRequest)
	if err != nil {
		fmt.Println("Do", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll", err)
		return err
	}
	bbd := string(body)
	fmt.Printf("xx bodys: %v\n", bbd)

	err = json.Unmarshal(body, res)
	if err != nil {
		fmt.Println("Unmarshal", err)
		return err
	}
	return nil
}
