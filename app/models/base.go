package models

import (
	"database/sql"
	"fmt"
	"log"
	"want_to_go/config"

	_ "github.com/mattn/go-sqlite3"
)

// tableを作成するコードを記載する

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		コマンドを作成していく(%sはtableNameUserを入れる)
		まずはじめにUsertableを作成
	*/
	cmdCreateUserTable := fmt.Sprintf(`CREATE table if not exists %s(
		id integer primary key autoincrement,
		uuid string not null unique,
		name string,
		email string,
		password string,
		created_at datetime)`, tableNameUser)

	/* 作成したコマンドの実行 */
	Db.Exec(cmdCreateUserTable)
}
