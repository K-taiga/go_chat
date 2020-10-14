package data

import "time"

// Thread DDL データ定義言語 テーブルを作成する際に使うstruct
type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// NumReplies postsのcountを返すメソッドを構造体に定義
func (thread *Thread) NumReplies() (count int) {
	// $1はプレイスホルダー　引数に置き換え
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return

}

// Threads sqlを発行して構造体にいれる
// Threadのキーを持ったスライス（参照型の可変長配列）が返り値
func Threads() (threads []Thread, err error) {
	// Dbにクエリを送信し、レコードを取得
	rows, err := Db.Query("SELECT id, uuid, topic, user_id,created_at FROM threads ORDER BY created_at DESC")

	if err != nil {
		return
	}

	// レコード分繰り返し
	for rows.Next() {
		th := Thread{}
		// if文実行前にScanでレコードを構造体に取り込みerrチェック
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return
		}
		// スライスにappend
		threads = append(threads, th)
	}
	rows.Close()
	return
}
