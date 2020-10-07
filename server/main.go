package main

import (
	"fmt"
	"net/http"
)

// イベントがきっかけとして起動されるコールバック関数 = handler ハンドラ関数
// がレスポンス返すインタフェースのResponseWirterとリクエストを受け取る構造体Request
func handler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

// 実行はこの関数から始まる
func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}