package main

import (
	"fmt"
	"net/http"
	"os"
)

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprintf(res, "My homapage....")
}

// env variable PORT with a port number is expected
func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe("localhost:"+os.Getenv("PORT"), nil)
}
