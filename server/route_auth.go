package main

import (
	"net/http"

	"local.packages/data"
)

// GET /login
func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// GET /signup
func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
func signupAccount(writer http.ResponseWriter, request *http.Request) {
	// POSTリクエストをパースして受け取る
	err := request.ParseForm
	if err != nil {
		danger(err, "Cannot parse form")
	}

	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)

}

func authenticate(writer http.ResponseWriter, request *http.Request) {
	// GoでPostの値を取得するにはParseFormを使う
	err := request.ParseForm()
	// emailを与えるとuser structを返す
	user, err := data.UserByEmail(request.PostFormValue("email"))

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
				Uuid      string [ユニークなID]
				Email     string
				UserId	  string
				CreatedAt time.Time
			}
		*/

		if err != nil {
			danger(err, "Cannot create session")
		}
		// ブラウザに保存するCookie構造体を作成
		// TTLは設定しない ブラウザを消したらcookieも削除
		cookie := http.Cookie{
			Name: "_cookie",
			// session structで作ったUuid
			Value: session.Uuid,
			// HTTPとHTTPSでのサーバーのやりとりのみでしかcookieにアクセスできない
			// jsとかからのAPI接続ではアクセスできなくする
			HttpOnly: true,
		}
		// レスポンスヘッダにクッキーを追加
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "login", 302)
	}

}

// GET /logout
// Logs the user out
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
