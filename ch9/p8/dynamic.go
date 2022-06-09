// Serving content generated with templates
package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("base.html"))
}

func main() {
	http.Handle("/", http.HandlerFunc(home))

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, "base.html")
}
