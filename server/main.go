package main

import (
	"net/http"
)

func main() {
	// リクエストをハンドラに渡すマルチプレクサ
	mux := http.NewServeMux()
	// 静的なファイルを配信するディレクトリ
	files := http.FileServer(http.Dir("/public"))
	// /static/がつくリクエストURLはpuclic配下にあるか探してそのまま返送するようにする
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// ルートURLとハンドラ関数を引数にとる　ハンドラ関数はResponseWirterとRequestを引数にするため、引数を渡す必要はない？
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	files := []string{
// 		"templates/layout.html",
// 		"templates/navbar.html",
// 		"templates/index.html",
// 	}
// 	templates := template.Must(template.ParseFiles(files...))
// 	threads, err := data.Threads()
// 	if err == nil {
// 		templates.ExecuteTemplate(w, "layout", threads)
// 	}
// }
