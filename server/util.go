package main

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"github.com/mushahiroyuki/gowebprog/ch02/chitchat/data"
)

// （引数)(返り値)
func session(writer http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	// requestからcookieを取り出す
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		// dbからUuidのセッションがあるかチェック
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// dataは空のインターフェースでどのような型でも受け付ける,fnはテンプレートファイルのリスト
// fn...stringは可変長引数でテンプレートファイルを複数受け付ける
func generateHTML(writer http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html"), file)
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
