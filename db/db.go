/*5.データベースを操作するために作成*/
package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // 外部パッケージのためgo mod tidyをターミナルで実行する
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*DBの構造体を作成 gorm.DBのポイント型を返り値として指定する
(gormは外部パッケージのためgo mod tidyをターミナルで実行する)
このコードは、gormライブラリを使用してデータベースに接続するための関数です。
os.Getenv("GO_ENV") == "dev"の部分は、環境変数GO_ENVの値がdevである場合に限り、
ローカルで動作する環境に対して設定ファイル.envから環境変数を取得するための処理を行います。
環境変数がdevでない場合は何もしません。
godotenv.Load()は、.envファイルから環境変数をロードする関数です。
この関数がエラーを返した場合、log.Fatalln(err)でエラーを出力してプログラムを終了させます。*/

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev"{
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	/*6.データベースに接続するための関数を作成*/
	// url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?", os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
	// 	os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		"5432", os.Getenv("POSTGRES_DB"))

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connceted")
	return db
}

/*7.データベースをCloseするためのCloseDB*/
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}