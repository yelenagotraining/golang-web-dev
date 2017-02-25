package main

import (
	"net/http"
	"html/template"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	subscribe := r.FormValue("subscribe")
	dessert := r.FormValue("dessert")
	fmt.Println("RESULTS:", subscribe, " - ", dessert)
}
