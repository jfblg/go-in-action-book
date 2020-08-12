package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// custom http client
func main() {
	cc := &http.Client{Timeout: time.Second} // this modifies timeout value
	res, err := cc.Get("http://example.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s\n", b)
}
