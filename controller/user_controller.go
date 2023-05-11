/*14.controllerの実装*/
package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4" //ターミナルでgo mod tidyを実行
)

/*15.IUserControllerの実装*/ /*36.Csrftokenの追加*/
type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

/*16.userControllerの実装*/
type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

/*17.SignUpの実装*/
func (uc *userController) SignUp(c echo.Context) error {
	//リクエストのバインド
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

/*18.LogInの実装*/
func (uc *userController) LogIn(c echo.Context) error {
	//リクエストのバインド
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//cookieの設定
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

/*19.LogOutの実装*/
func (uc *userController) LogOut(c echo.Context) error {
	//cookieの設定
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
	}
	
	/*37.CsrfTokenの実装*/
	func (uc *userController) CsrfToken(c echo.Context) error {
		token := c.Get("csrf").(string)
		return c.JSON(http.StatusOK, echo.Map{
			"csrf_token": token,
		})
	}