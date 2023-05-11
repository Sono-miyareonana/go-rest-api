/*26*/
package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

/*interfaceを定義*/
type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

/*taskController構造体を定義*/
type taskController struct {
	tu usecase.ITaskUsecase
}

/*NewTaskController関数を定義*/
func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

/*GetAllTasksメソッドを定義*/
func (tc *taskController) GetAllTasks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	tasksRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasksRes)
}
/*
このコードは、Go言語で書かれたWebアプリケーションの一部で、Echoフレームワークを使用しています。具体的には、`taskController`構造体のメソッド`GetAllTasks`を定義しており、これはHTTPリクエストを受け取って、指定されたユーザーのすべてのタスクを取得し、JSON形式で返す役割を果たします。以下に、コードの概要を説明します。
1. `func (tc *taskController) GetAllTasks(c echo.Context) error`：`GetAllTasks`という名前のメソッドを定義しており、taskController構造体に属しています。このメソッドは、Echoフレームワークの`Context`インターフェースを引数に取り、エラーを返すことができます。
2. `user := c.Get("user").(*jwt.Token)`：コンテキストから"user"というキーでJWT（JSON Web Token）を取得し、`*jwt.Token`型にキャストしています。これは、認証されたユーザー情報を取得するために使用されます。
3. `claims := user.Claims.(jwt.MapClaims)`：JWTトークンからクレームを抽出し、`jwt.MapClaims`型にキャストしています。クレームは、トークンに関する追加情報を提供し、ユーザー属性や権限に関する情報が含まれることがあります。
4. `userId := claims["user_id"]`：`user_id`というキーでクレームからユーザーIDを取得しています。
5. `tasksRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))`：`taskController`の`tu`フィールド（タスク関連の処理を行うオブジェクト）の`GetAllTasks`メソッドを呼び出し、ユーザーIDを引数にしています。ユーザーIDを`float64`から`uint`型にキャストしています。このメソッドは、すべてのタスクを取得するために使われ、その結果とエラーが返されます。
6. エラーチェック：`if err != nil`ブロックで、`GetAllTasks`メソッドから返されたエラーをチェックしています。エラーがある場合、HTTPステータスコード500（Internal Server Error）とエラーメッセージをJSON形式で返します。
7. `return c.JSON(http.StatusOK, tasksRes)`：エラーがない場合、HTTPステータスコード200（OK）と取得したタスクのリストをJSON形式で返します。*/

/*GetTaskByIdメソッドを定義*/
func (tc *taskController) GetTaskById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}
/*
このコードもGo言語で書かれており、Echoフレームワークを使用しています。`taskController`構造体の`GetTaskById`メソッドを定義しており、HTTPリクエストを受け取って、指定されたユーザーIDとタスクIDに基づいて特定のタスクを取得し、JSON形式で返す役割を果たします。以下に、コードの概要を説明します。
1. `func (tc *taskController) GetTaskById(c echo.Context) error`：`GetTaskById`という名前のメソッドを定義しており、`taskController`構造体に属しています。このメソッドは、Echoフレームワークの`Context`インターフェースを引数に取り、エラーを返すことができます。
2. `user := c.Get("user").(*jwt.Token)`：コンテキストから"user"というキーでJWT（JSON Web Token）を取得し、`*jwt.Token`型にキャストしています。これは、認証されたユーザー情報を取得するために使用されます。
3. `claims := user.Claims.(jwt.MapClaims)`：JWTトークンからクレームを抽出し、`jwt.MapClaims`型にキャストしています。クレームは、トークンに関する追加情報を提供し、ユーザー属性や権限に関する情報が含まれることがあります。
4. `userId := claims["user_id"]`：`user_id`というキーでクレームからユーザーIDを取得しています。
5. `id := c.Param("taskId")`：ParamはEchoフレームワークのContextインターフェースのメソッドです。c.Param("taskId")は、リクエストURLのパスパラメータからtaskIdを取得しています。これは、URLの一部として送信されるパラメータであり、リソースを識別するために使用されます。
6. `taskId, _ := strconv.Atoi(id)`：strconv.Atoiは、Go言語のstrconvパッケージにある関数で、文字列を整数（具体的にはint型）に変換する機能を提供します。`taskId`を文字列から整数に変換しています。エラーは無視されています。
7. `taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))`：`taskController`の`tu`フィールド（タスク関連の処理を行うオブジェクト）の`GetTaskById`メソッドを呼び出し、ユーザーIDとタスクIDを引数にしています。ユーザーIDを`float64`から`uint`型にキャストしています。このメソッドは、指定されたタスクを取得するために使われ、その結果とエラーが返されます。
8. `エラーチェック：`if err != nil`ブロックで、`GetTaskById`メソッドから返されたエラーをチェックしています。エラーがある場合、HTTPステータスコード500（Internal Server Error）とエラーメッ
*/

/*CreateTaskメソッドを定義*/
func (tc *taskController) CreateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserId = uint(userId.(float64))
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}
/*このコードは、Go言語で書かれた`taskController`構造体の`CreateTask`メソッドです。このメソッドは、Echoフレームワークを利用してHTTPリクエストを処理し、新しいタスクを作成する機能を提供します。
1. メソッドは、Echoフレームワークの`Context`を引数に受け取り、エラーを返します。
2. JWTトークンからユーザー情報を取得し、`claims`に格納します。`userId`は、`claims`から取得されます。
3. `model.Task`型の新しいタスクインスタンスを作成します。
4. `c.Bind(&task)`を使用して、リクエストボディからタスク情報を取得し、`task`インスタンスにバインドします。バインド中にエラーが発生した場合、ステータスコード400（Bad Request）とエラーメッセージを返します。
	Bindは、Echoフレームワークにおいて、HTTPリクエストのボディからデータを読み込み、指定された構造体にデータをマッピングするためのメソッドです。
5. `task.UserId`に、取得した`userId`を`uint`型に変換して設定します。
6. `tc.tu.CreateTask(task)`を呼び出して、新しいタスクを作成します。この操作中にエラーが発生した場合、ステータスコード500（Internal Server Error）とエラーメッセージを返します。
7. タスクの作成が成功した場合、ステータスコード201（Created）と作成されたタスクの情報を返します。
この`CreateTask`メソッドは、タスクを作成するエンドポイントのハンドラーとして機能し、リクエストボディからタスク情報を取得し、データベースまたは他のストレージシステムに保存する処理を行います。*/

/*UpdateTaskメソッドを定義
このUpdateTaskメソッドは、タスクを更新するエンドポイントのハンドラーとして機能し、
リクエストボディからタスク情報を取得し、データベースまたは他のストレージシステムで
既存のタスクを更新する処理を行います。
最終的にエラーが発生しなければ更新後のタスク情報を返します。
*/
func (tc *taskController) UpdateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

/*DeleteTaskメソッドを定義
このDeleteTaskメソッドは、タスクを削除するエンドポイントのハンドラーとして機能し、
リクエストボディからタスク情報を取得し、データベースまたは他のストレージシステムで
既存のタスクを削除する処理を行います。
*/
func (tc *taskController) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
