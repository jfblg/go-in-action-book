package main

import "net/http"

func main() {
	dir := http.Dir("./files")
	handler := http.StripPrefix("/static/", http.FileServer(dir))
	http.Handle("/static/", handler)

	http.HandleFunc("/", readme)
	http.ListenAndServe(":8088", nil)
}

func readme(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./files/file_1.txt")
}
