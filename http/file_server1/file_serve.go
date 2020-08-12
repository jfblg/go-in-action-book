package main

import "net/http"

func main() {
	dir := http.Dir("./files")
	http.ListenAndServe(":8088", http.FileServer(dir))
}
