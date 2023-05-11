/*24.task_repository.go*/
package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*repositoryのinterfaceを定義*/
type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

/*TaskRepository構造体を定義*/
type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

/*GetAllTasksメソッドを定義*/
func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").
		Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

/*GetTaskByIdメソッドを定義*/
func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).
		First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

/*Createメソッドを定義*/
func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

/*UpdateTaskメソッドを定義*/
func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).
	Where("id =? AND user_id=?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

/*DeleteTaskメソッドを定義*/
func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id =? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

/*
このコードはデータベース操作を行うリポジトリ層のコードです。

まず、リポジトリ層で使うためのinterfaceである`ITaskRepository`を定義します。
次に、`taskRepository`構造体を定義し、
`NewTaskRepository`関数で`taskRepository`のインスタンスを生成します。

`taskRepository`は`db *gorm.DB`を持ち、`ITaskRepository`で定義された関数を実装しています。

`GetAllTasks`関数は、全てのタスクを取得する関数です。
`tasks`に取得したタスクのスライスを格納し、`userId`を基にユーザーが所有するタスクのみを取得します。

`GetTaskById`関数は、指定した`taskId`と`userId`に対応するタスクを取得する関数です。

`CreateTask`関数は、新しいタスクをデータベースに登録する関数です。

`UpdateTask`関数は、指定された`taskId`と`userId`に対応するタスクのタイトルを更新する関数です。

`DeleteTask`関数は、指定された`taskId`と`userId`に対応するタスクを削除する関数です。

これらの関数で使用される`tr.db`は、GORMを利用して生成されたデータベースのインスタンスです。
GORMを使用することで、SQL操作を行うことなく、Goのコードを通じてデータベースを操作することができます。*/