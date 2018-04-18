package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("me.gohtml"))

}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index page")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog  page")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "me.gohtml", "我的模板.....")
}
