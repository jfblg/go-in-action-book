package main

import "net/http"

// server files with custom handler

func main() {
	http.HandleFunc("/", readme)
	http.ListenAndServe(":8088", nil)
}

func readme(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./files/file_1.txt")
}
