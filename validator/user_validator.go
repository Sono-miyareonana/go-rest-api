/*31.validator*/
package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

/*interface*/
type IUserValidator interface {
	UserValidate(user model.User) error
}

/*struct*/
type userValidator struct{}

/*NewUserValidator method*/
func NewUserValidator() IUserValidator {
	return &userValidator{}
}

/*UserValidate method*/
func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
		),
	)
}

/*
解説：このコードは、ユーザー情報のバリデーションを行うための機能を定義しています。
1. `IUserValidator` インターフェース:
    ユーザー情報のバリデーションを行うメソッド`UserValidate` を定義しています。
2. `userValidator` 構造体: 
    ユーザーバリデーション機能を実装するための構造体です。
3. `NewUserValidator` 関数: 
   `IUserValidator` インターフェースを実装する新しい `userValidator` インスタンスを返します。
4. `UserValidate` メソッド: 
   `userValidator` のメソッドとして定義され、引数として `model.User` を受け取り、
   バリデーションを実行します。このメソッドは、次のバリデーションルールを適用します。
   - メールアドレス (`user.Email`) に対するバリデーション:
     - 必須チェック (空でないこと)
     - 文字数制限 (1～30文字)
     - メールアドレスの形式チェック
   - パスワード (`user.Password`) に対するバリデーション:
     - 必須チェック (空でないこと)
     - 文字数制限 (6～30文字)
　　`UserValidate` メソッドは、バリデーションが成功した場合は `nil` を返し、
　　失敗した場合はエラーメッセージを含むエラーオブジェクトを返します。
　　これにより、呼び出し元はバリデーションの結果に応じて適切な処理やエラーハンドリングを行うことができます。

インスタンスとは
クラスや構造体は、プログラム内でオブジェクトを作成するための設計図のようなもので、
それに基づいて具体的なオブジェクトが作成されます。この具体的なオブジェクトがインスタンスです。
インスタンスは、クラスや構造体が持つデータ（プロパティ）と操作（メソッド）を具体的なオブジェクトとして実体化します。
これにより、プログラムの他の部分から、インスタンスを使ってプロパティにアクセスしたり、メソッドを実行することができます。
*/