package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Example #1",
		Content: "Which shows how to use html templates in Go",
	}
	t := template.Must(template.ParseFiles("templates/simple.html"))
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8088", nil)
}
