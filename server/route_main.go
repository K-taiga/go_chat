package main

import (
	"html/template"
	"net/http"

	// "github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
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

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads(); if err == nill {
		// エラーかどうかだけ知りたいため、_（ブランク識別子)にSession構造体を入れる
		-, err := session(writer,request)

		if err != nil {
			generateHTML(writer,threads,"layout","public.navbar","index")
		} else {
			generateHTML(writer,threads,"layout","private.navbar","index")
		}
	}
}
