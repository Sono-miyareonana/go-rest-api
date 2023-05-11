package model

import "time"

/*1.ユーザー情報を表現する構造体を定義します*/
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/*2.クライアントに返すユーザー情報を表現する構造体を定義します*/
type UserResponse struct {
	ID    uint   `json:"id" gorm:"primarykey"`
	Email string `json:"email" gorm:"unique"`
}

/*①補足１
jsonとは、JavaScript Object Notationの略で、データの構造を表現するための軽量なフォーマットです。
JSONは、人間にとっても機械にとっても読みやすく、様々なプログラミング言語で扱いやすいため、
データのやり取りや保存によく使われます。
・ID uint: ID は、ユーザーの一意の識別子を表します。uint は符号なし整数型です。
・Email string: Email は、ユーザーのメールアドレスを表します。string は文字列型です。
　ユーザーのメールアドレスにunique制約がある場合、
　同じメールアドレスを持つ複数のユーザーが存在することはできません。
　これにより、メールアドレスを使ってユーザーを一意に識別することができます。
・Password string: Password は、ユーザーのパスワードを表します。これも string 型です。
・CreatedAt time.Time: CreatedAt は、ユーザーが作成された日時を表します。
　time.Time は、Go の標準パッケージである time パッケージで定義されている日時型です。
・UpdatedAt time.Time: UpdatedAt は、ユーザー情報が最後に更新された日時を表します。
　これも time.Time 型です。
・フィールド定義の後ろにあるバッククォート（`）で囲まれた部分は、
　タグ（tag）と呼ばれるメタデータです。
　タグは、フィールドに追加情報を付与するために使用されます。
　このコードでは、JSON タグと GORM タグが使用されています。
・JSON タグは、構造体を JSON 形式に変換する際に、フィールド名を指定します。
　例えば、ID uint 'json:"id"' では、ID フィールドを JSON に変換する際に、
　キー名として "id" を使用することを指定しています。
・GORM タグは、GORM という Go の ORM（Object-Relational Mapping）ライブラリで使用されます。
　これにより、データベースと構造体のマッピングが可能になります。
　例えば、Email string 'gorm:"unique"' では、Email フィールドが一意であることを
　GORM に伝えています。
この構造体を使って、ユーザー情報を表現したり、データベースとやり取りしたりすることができます。*/

/*②補足２　gormを使う理由
・コードの簡潔さ：
　gormを使うことで、データベース操作に関連するコードを簡潔に書くことができます。
　これにより、開発者はより少ないコードでデータベース操作を実行できるため、開発速度が向上します。
・データベースの抽象化：
　gormは、様々なデータベースシステム（MySQL、PostgreSQL、SQLiteなど）と
　互換性があります。そのため、データベースシステムを変更しても、gormを使用している限り、
　コードの変更が最小限で済みます。
・エラーの防止：
　gormは、データベース操作に関連する一般的なエラー（SQLインジェクションなど）を防ぐ機能を提供しています。
　これにより、セキュリティが向上します。
以上の理由から、gormを使用することで、開発者はデータベース操作に関するコードを効率的に書くことができ、
プロジェクトの品質とセキュリティが向上します。*/
