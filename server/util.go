package main

import (
	"errors"
	"net/http"

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
