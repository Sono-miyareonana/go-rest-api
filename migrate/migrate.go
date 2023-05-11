package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

/*8.マイグレーションを実行するための関数を作成
マイグレーションとは、
データベースにテーブルを作成したり、テーブルの構造を変更したりするためのプロセスです。
アプリケーションのバージョンアップや機能追加に伴って、データベースの構造を変更する必要がある場合に使用されます。
・main() 関数内では、データベースへの接続を確立するために、 db.NewDB() 関数を呼び出します
・次に、データベースに対してマイグレーションを実行します。
　これは、データベーススキーマを変更するための操作を、アプリケーションの起動時に自動的に行うために使用されます。
・最後に、defer 文を使用して、アプリケーションの処理が終了する際に、データベース接続をクローズすることを保証します。
つまり、このコードはデータベースに接続して、 User と Task のマイグレーションを自動化するものです。*/

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}

// ターミナルで実行　GO_ENV=dev go run migrate/migrate.go