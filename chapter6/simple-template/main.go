package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title   string
	Content string
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "HTML GO Template",
		Content: "Have fun good beauty paradise",
	}
	t := template.Must(template.ParseFiles("templates/simple.html"))
	t.Execute(w, p)
}
