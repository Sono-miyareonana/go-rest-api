/*30.validator*/
package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

/*interface*/
type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

/*struct*/
type taskValidator struct{}

/*method*/
func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

/*TaskvalidatorのTaskValidateメソッド実装*/
func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
         &task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		),
	)
}
/*
解説
このコード全体で、タスクのバリデーションを行う `ITaskValidator` インターフェースと、
その実装である `taskValidator` 構造体を定義しています。
これにより、他の部分でタスクのバリデーションを行いたい場合、`ITaskValidator` を利用することができます。
1. **interface**：`ITaskValidator` インターフェースを定義しています。
   このインターフェースには、`TaskValidate` というメソッドが定義されています。
2. **struct**：`taskValidator` という構造体（struct）を定義しています。
   この構造体は、後で `ITaskValidator` インターフェースを実装します。
3. **method**：`NewTaskValidator` という関数を定義しています。
   この関数は、`ITaskValidator` インターフェースを実装した
   新しい `taskValidator` 構造体のポインタを返します。
4. **TaskvalidatorのTaskValidateメソッド実装**：
   `taskValidator` 構造体に対して、`ITaskValidator` インターフェースで定義されている 
   `TaskValidate` メソッドを実装しています。
   このメソッドは、引数として `model.Task` を受け取り、バリデーションを行い、エラーがあれば返します。
   バリデーションでは、`validation.ValidateStruct` 関数を使って、
   構造体のフィールドに対してバリデーションルールを適用しています。
   このコードでは、`title` フィールドに対して以下のバリデーションルールが適用されています。
   - `validation.Required`：
   `title` が必須であることをチェックし、エラーメッセージ "title is required" を設定しています。
   - `validation.RuneLength(1, 10)`：
   `title` の長さが1文字以上10文字以下であることをチェックし、エラーメッセージ "limited max 10 char" を設定しています。

バリデーションとは、データが特定のルールや条件に従っているかどうかを検証するプロセスです。
バリデーションは、プログラムが予期しないデータや誤った形式のデータを受け取らないようにするために重要な役割を果たします。
データバリデーションは、ユーザー入力や外部からのデータ、または他のシステムとのインタラクションの際に特に重要です。
バリデーションは、以下のような目的で実行されます。
1. データの整合性を保つ：データがアプリケーションの要件に従っていることを確認し、
   データベースや他のシステムに正しいデータが保存されるようにします。
2. エラーを回避する：不正なデータや形式のデータを処理しようとすると、プログラムがエラーを起こしたり、
   予期しない動作をしたりする可能性があります。バリデーションによって、エラーを事前に回避できます。
3. セキュリティを向上させる：悪意のあるデータや攻撃を防ぐために、バリデーションを使ってデータをチェックすることができます。
4. ユーザーエクスペリエンスを向上させる：バリデーションを使って、ユーザーに入力エラーや不足している情報をフィードバックすることで、
   ユーザーエクスペリエンスを向上させることができます。
バリデーションは、アプリケーションの種類や要件に応じて、さまざまなルールや方法で実行されます。
例えば、文字列の長さ、数値の範囲、日付の形式、必須フィールドの存在、正規表現によるパターンマッチングなどがあります。
*/