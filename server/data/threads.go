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
