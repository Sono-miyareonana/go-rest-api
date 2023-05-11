/*20.router.go*/
package router

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4" //27.echo-jwtのインポート(ターミナルでgo mod tidyを実行)
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*21.NewRouterの実装*/ /*28.TaskControllerの追加*/
func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	/*34.CORSmiddlewareの追加*/
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowCredentials: true,
	}))
	/*35.CSRFmiddlewareの追加*/
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath: "/",
		CookieDomain: os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge: 60,
	}))

	//userのルーティング
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	/*27.taskのルーティング*/
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	/*28.TaskControllerの追加*/
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}

/*
27解説
この`NewRouter`関数は、APIのルーティング構成を行い、
エンドポイントとそれらに対応するハンドラー関数を関連付けます。
また、JWT認証を適用することで、セキュリティを強化し、
認証されたユーザーのみが特定のエンドポイントにアクセスできるようにします。
1. Echoフレームワークの新しいインスタンスを作成します。
2. ユーザー関連のルーティングを設定します。`/signup`、`/login`、`/logout`のエンドポイントに対して、それぞれ`uc.SignUp`、`uc.LogIn`、`uc.LogOut`ハンドラー関数を登録します。
3. タスク関連のルーティングを設定するために、`/tasks`をベースパスとする新しいルーティンググループを作成します。
4. JWT認証を有効にするため、`t.Use(echojwt.WithConfig(echojwt.Config{...}))`を使用して、作成したルーティンググループにミドルウェアを適用します。ミドルウェアは、リクエストが処理される前に実行される関数で、ここではJWT認証を行います。認証に使用される秘密鍵、トークンの検索方法（ここではクッキー）など、JWT認証の設定を提供します。
5. 設定されたルーティングを持つ`*echo.Echo`インスタンスを返します。

28解説
1.`NewRouter`関数の引数に`tc controller.ItaskController`を追加しました。これにより、タスクコントローラーが関数に渡されるようになります。
2.タスク関連のエンドポイントとハンドラー関数を関連付けました：
  - `t.GET("", tc.GetAllTasks)`: 全てのタスクを取得するエンドポイント（`/tasks`）に`tc.GetAllTasks`ハンドラー関数を登録します。
  - `t.GET("/:taskId", tc.GetTaskById)`: 特定のタスクを取得するエンドポイント（`/tasks/:taskId`）に`tc.GetTaskById`ハンドラー関数を登録します。
  - `t.POST("", tc.CreateTask)`: 新しいタスクを作成するエンドポイント（`/tasks`）に`tc.CreateTask`ハンドラー関数を登録します。
  - `t.PUT("/:taskId", tc.UpdateTask)`: 特定のタスクを更新するエンドポイント（`/tasks/:taskId`）に`tc.UpdateTask`ハンドラー関数を登録します。
  - `t.DELETE("/:taskId", tc.DeleteTask)`: 特定のタスクを削除するエンドポイント（`/tasks/:taskId`）に`tc.DeleteTask`ハンドラー関数を登録します。
これらの変更により、`NewRouter`関数は、ユーザー関連のルーティングだけでなく、タスク関連のルーティングも設定し、それらに対応するハンドラー関数を関連付けることができます。また、JWT認証ミドルウェアがタスク関連のルーティンググループに適用されているため、認証されたユーザーのみがこれらのエンドポイントにアクセスできるようになります。
*/