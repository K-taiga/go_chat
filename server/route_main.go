package main

import (
	"html/template"
	"net/http"

	"github.com/mushahiroyuki/gowebprog/ch02/chitchat/data"
)

// GET /err?msg=
// error画面
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()

	_, err := session(writer, request)

	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
