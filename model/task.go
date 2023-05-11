/*taskのmodelを作成*/
package model

import "time"

/*３.taskの構造体を作成*/
type Task struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

/*補足:
`json:"user" gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE"`
 データが更新された時に、関連するデータも更新するために設定 */

/*4.クライアント側からリクエストされた時に返すtaskのデータ構造作成*/
type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}