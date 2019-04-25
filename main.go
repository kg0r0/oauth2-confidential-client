package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := getData(r)
	tpl.ExecuteTemplate(w, "home.gohtml", data)
}

func getData(r *http.Request) string {
	return "Hello"
}

func main() {
	http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
