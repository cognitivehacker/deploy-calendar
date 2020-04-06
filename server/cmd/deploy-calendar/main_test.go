package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

func errorCode(expected int, received int, route string) string {
	return fmt.Sprintf("Expecting Http Status Code: %v, but received: %v at %s", expected, received, route)
}

func TestMain(t *testing.T) {

	var url = "http://127.0.0.1:8080"

	doneChan := make(chan struct{})
	go func() {
		main()
	}()

	time.Sleep(1 * time.Second) // Wait for fasthttp start

	go func() {

		var resp []byte

		code, _, _ := fasthttp.Get(resp, url)
		if code != 200 {
			t.Errorf(errorCode(200, code, url))
		}

		doneChan <- struct{}{}
	}()

	<-doneChan
}

func ExampleMain() {

	time.Sleep(1 * time.Second) // Wait for fasthttp start

	main()
	// Output:
	// Error in ListenAndServe:  listen tcp4 :8080: bind: address already in use
}
