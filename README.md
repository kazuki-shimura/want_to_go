# want_to_go

パッケージ内で絶対参照できるようにする
go mod init want_to_go
いらないものを整理する
go mod tidy

この状態で
go run main.go を実行すると
[エラーメッセージ]
2021/03/13 13:55:43 open config.ini: no such file or directory
exit status 1
が表示される
これをした上で config の下から want_to_go の下に変更すると読み込める
理由はまだわからないのであとで調べる必要がある

GO111MODULE=on して、常に Module-aware モードとなる状態で検証
export GO111MODULE=on
