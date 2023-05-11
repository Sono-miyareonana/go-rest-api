/*10*/
package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4" //ターミナルでgo mod tidyを実行
	"golang.org/x/crypto/bcrypt"
)

/*interface*/
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

/*struct*/ //usecaseの中でrepositoryとvalidatorを使う
type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator /*33.validator追加*/
}

/*NewUserUsecase method*/ /*33.validator追加*/
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv} 
}

/*SignUp*/
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	/*33.validator追加*/
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}

	//パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}

	//ユーザーの作成
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	//レスポンスの作成
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

//Login
func (uu *userUsecase) Login(user model.User) (string, error) {
	/*33.validator追加*/
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	
	//ユーザーの取得
	storeUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storeUser, user.Email); err != nil {
		return "", err
	}
	
	//パスワードの検証
	err := bcrypt.CompareHashAndPassword([]byte(storeUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	
	//JWTの発行
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storeUser.ID,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*
10.解説
このコードは、ユーザー認証（サインアップとログイン）を実行するための Go 言語で書かれた実装です。
ここでは、`IUserUsecase` インターフェースが定義されており、`SignUp` と `Login` の2つのメソッドが含まれています。
また、`userUsecase` 構造体がこのインターフェースを実装しています。
1. `NewUserUsecase` 関数は、`IUserUsecase` インターフェースを実装する新しいインスタンスを作成し、
	リポジトリ（`repository.IUserRepository`）を引数として受け取ります。
2. `SignUp` メソッドは、ユーザー情報（`model.User`）を受け取り、パスワードをハッシュ化してデータベースに保存します。保存が成功した場合、新しいユーザー情報（`model.UserResponse`）が返されます。ハッシュ化やユーザーの作成でエラーが発生した場合、エラーが返されます。
3. `Login` メソッドは、ユーザー情報（`model.User`）を受け取り、データベースから該当するユーザーを取得し、
	パスワードの検証を行います。パスワードが一致した場合、JWT トークンが生成され、その文字列が返されます。
	パスワードが一致しない場合や、トークンの生成でエラーが発生した場合、エラーが返されます。
この実装では、ユーザー認証の基本的な処理（サインアップ、ログイン、パスワードのハッシュ化、JWT トークンの生成）が
行われています。これらの機能は、WebアプリケーションやAPIでよく使用されるものです。

33.解説
`userUsecase` 構造体にバリデーション機能を追加したものです。変更点は以下の通りです。
1. `userUsecase` 構造体に、`validator.IUserValidator` を使用するための `uv` フィールドが追加されました。
2. `NewUserUsecase` 関数が変更され、`validator.IUserValidator` を引数として受け取るようになりました。これにより、`userUsecase` 構造体のインスタンスを作成する際に、リポジトリとバリデータを同時に渡すことができます。
3. `SignUp` と `Login` メソッドの冒頭に、バリデーション処理（`uu.uv.UserValidate(user)`）が追加されました。これにより、メソッドが実行される前に、渡されたユーザー情報が適切かどうかを確認できます。
この変更により、ユーザー認証の処理にバリデーション機能が組み込まれ、より堅牢な実装になりました。不正なデータがシステムに入ることを防ぐことができるようになります。
*/