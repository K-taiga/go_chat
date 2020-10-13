package main

import (
	"net/http"

	// "github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"github.com/mushahiroyuki/gowebprog/ch02/chitchat/data"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	// GoでPostの値を取得するにはParseFormを使う
	r.ParseForm()

	// emailを与えるとuser structを返す
	user, _ := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}

	// pwをencryptしてチェック
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		// Session構造体を作成
		session, err := user.CreateSession()
		/*
			type Session struct {
				Id        int
				Uuid      string　[ユニークなID]
				Email     string
				UserId	  string
				CreatedAt time.Time
			}
		*/

		// ブラウザに保存するCookie構造体を作成
		// TTLは設定しない　ブラウザを消したらcookieも削除
		cookie := http.Cookie{
			Name: "_cookie",
			// session structで作ったUuid
			Value: session.Uuid,
			// HTTPとHTTPSでのサーバーのやりとりのみでしかcookieにアクセスできない
			// jsとかからのAPI接続ではアクセスできなくする
			HttpOnly: true,
		}
		// レスポンスヘッダにクッキーを追加
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "login", 302)
	}

}
