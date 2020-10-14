package data

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func init() {
	// errという変数をerrorの型で定義
	/*
		type error interface {
			Error() string
		}
	*/
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
