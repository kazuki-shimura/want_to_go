package models

import (
	"log"
	"time"
)

// Userの方を定義する
type User struct {
	ID        string
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	/*
		コマンドを実行するExecメソッドを使用する
		引数はクエリ文字列とパラメータ返り値はResult型とerror型を返す
		パスワードやUUIDはbase.goでメソッドを作成しているが何もしていない場合は下記となる予定
		_, err = Db.Exec(cmd, u.UUID, u.Name, u.Email, u.Password, u.CreatedAt)
		time.Nowは現在の時間を取得している
	*/
	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
