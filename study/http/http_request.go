package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

var httpTransport *http.Transport = &http.Transport{
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

func main() {
	ch := make(chan int, 0)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		time.Sleep(100 * time.Millisecond)

		go func(n int, c chan int, w *sync.WaitGroup) {
			defer w.Done()

			httpRequest, err := http.NewRequest("POST", "https://xxx", bytes.NewBuffer([]byte("{0}")))
			if err != nil {
				fmt.Println(n, err)
				return
			}
			httpRequest.Header.Set("Cookie", "token-qa2=xxx")

			client := &http.Client{
				Transport: httpTransport,
			}
			resp, err := client.Do(httpRequest)
			if err != nil {
				fmt.Println(n, err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(n, err)
				return
			}
			if n%100 == 0 {
				fmt.Println("\n===", n)
			}
			if resp.StatusCode == http.StatusOK {
				fmt.Printf("\rbody[%d]: %d", n, len(string(body)))
				c <- 1
			}
		}(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	cnt := 0
	for {
		k, ok := <-ch
		if !ok {
			break
		}
		cnt += k
	}
	fmt.Println("\nsucces: ", cnt)
}
