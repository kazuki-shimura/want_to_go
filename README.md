# want_to_go

パッケージ内で絶対参照できるようにする

    go mod init want_to_go

いらないものを整理する

    go mod tidy

この状態で下記を実行

    go run main.go

エラーメッセージが表示される
[エラーメッセージ]
2021/03/13 13:55:43 open config.ini: no such file or directory
exit status 1

これをした上で config の下から want_to_go の下に変更すると読み込める
理由はまだわからないのであとで調べる必要がある

GO111MODULE=on して、常に Module-aware モードとなる状態で検証

    export GO111MODULE=on

##SQL の追加

SQLite を使用するために必要ライブラリをインストールする

    go get github.com/mattn/go-sqlite3

sql を読み込むためにドライバを使用できるよに import 文に追記する

    _ "github.com/mattn/go-sqlite3"

##データベースを見る場合
SQLite に入る
sqlite3 webapp.sql

    テーブルを確認
    .table

    SQLiteを抜ける
    .exit
