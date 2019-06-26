package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
		fmt.Println(err)
	}
}

func httpSend(req interface{}, res interface{}, token string) error {
	buffer, err := json.Marshal(req)
	if err != nil {
		return err
	}
	httpRequest, err := http.NewRequest("POST",
		"http://127.0.0.1:8080/detail",
		bytes.NewBuffer(buffer))
	if err != nil {
		return err
	}
	httpRequest.Header.Set("Cookie", token)

	client := &http.Client{}
	resp, err := client.Do(httpRequest)
	if err != nil {
		return err
	}
	defer resp.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("body: %s\n", string(body))
	err = json.Unmarshal(body, res)
	if err != nil {
		return err
	}
	return nil
}
