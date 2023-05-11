/*22.main.goの実装*/
package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

/*main関数を定義*/
/*29.ディペンデンシーインジェクションを追加*/
/*34.validatorを追加*/
func main() {
	db := db.NewDB()
	usetValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, usetValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}

/* 
23.ターミナルで GO_ENV=dev go run main.go を実行
・PostmanでPOST http://localhost:8080/signup　に　send
・pgAdminでpwがハッシュ化されていることなどを確認
・PostmanでPOST http://localhost:8080/login　に　send
・PostmanでPOST http://localhost:8080/logout　に　send
・PostmanでGET http://localhost:8080/users　に　send

29.ディぺンデンシーインジェクションの追加
1. `db.NewDB()`を呼び出し、データベース接続を作成します。
2. `repository.NewUserRepository(db)`と`repository.NewTaskRepository(db)`を呼び出し、ユーザーとタスクのリポジトリインスタンスを作成します。リポジトリは、データベースとのやりとりを担当します。
3. `usecase.NewUserUsecase(userRepository)`と`usecase.NewTaskUsecase(taskRepository)`を呼び出し、ユーザーとタスクのユースケースインスタンスを作成します。ユースケースは、ビジネスロジックを実装します。
4. `controller.NewUserController(userUsecase)`と`controller.NewTaskController(taskUsecase)`を呼び出し、ユーザーとタスクのコントローラーインスタンスを作成します。コントローラーは、HTTPリクエストの処理とレスポンスの生成を担当します。
5. `router.NewRouter(userController, taskController)`を呼び出し、ルーティングを設定し、`*echo.Echo`インスタンスを作成します。
6. `e.Logger.Fatal(e.Start(":8080"))`を呼び出し、ポート8080でHTTPサーバーを起動します。
この`main`関数により、アプリケーションの各コンポーネントが適切に初期化され、組み立てられます。その後、HTTPサーバーが起動し、エンドポイントとそれらに対応するハンドラー関数を利用してリクエストを処理します。

34.validatorを追加
変更前のコードと変更後のコードの違いは、バリデータを作成してユースケースに渡している部分です。
1. `usetValidator` と `taskValidator` が作成されています。
	これらは、それぞれ `validator.NewUserValidator()` と `validator.NewTaskValidator()` によって生成されたバリデータのインスタンスです。
2. `userUsecase` と `taskUsecase` の作成時に、それぞれのバリデータが追加されています。これにより、ユースケース内でバリデーションを行うことができます。
変更後のコードでは、ユーザーとタスクのバリデーションが組み込まれており、より堅牢な実装になっています。この変更により、不正なデータがシステムに入ることを防ぐことができます。
*/