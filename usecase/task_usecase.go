/*25*/
package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

/*interfaceを定義*/
type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

/*taskUsecase構造体を定義*/ /*32.validetorを追加*/
type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

/*NewTaskUsecase関数を定義*/ /*32.validatorを追加*/
func NewTaskUsecase(tr repository.ITaskRepository,tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr,tv}
}

/*GetAllTasksメソッドを定義*/
func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t:= model.TaskResponse{
			ID: v.ID,
			Title: v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

/*GetTaskByIdメソッドを定義*/
func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID: task.ID,
		Title: task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

/*CreateTaskメソッドを定義*/
func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	/*32.validatorを追加*/
	if err := tu.tv.TaskValidate(task); err != nil{
		return model.TaskResponse{}, err
	}
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID: task.ID,
		Title: task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

/*UpdateTaskメソッドを定義*/
func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error){
	/*32.validatorを追加*/
	if err := tu.tv.TaskValidate(task); err != nil{
		return model.TaskResponse{}, err
	}	
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID: task.ID,
		Title: task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

/*DeleteTaskメソッドを定義*/
func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}

/*
解説25
このコードは、タスクに関するビジネスロジックを扱うためのUsecaseインターフェースと
その実装であるtaskUsecase構造体、およびその構造体のメソッドで構成されています。

まず、ITaskUsecaseというインターフェースを定義しています。
このインターフェースは、タスクに関するUsecaseが実装すべきメソッドを定義しています。

次に、taskUsecase構造体を定義しています。
この構造体は、ITaskUsecaseインターフェースを実装しています。
taskUsecase構造体は、ITaskRepositoryというインターフェースを実装したリポジトリを使用して、
タスクに関するビジネスロジックを実装します。

NewTaskUsecase関数は、ITaskRepositoryを引数に取り、ITaskUsecaseを返します。
この関数は、taskUsecase構造体のポインタを作成し、
そのポインタをITaskUsecaseインターフェースにキャストして返します。

GetAllTasksメソッドは、ユーザーIDを引数に取り、そのユーザーに関連するすべてのタスクを取得します。
取得されたタスクは、model.TaskResponseという構造体に変換され、スライスとして返されます。

GetTaskByIdメソッドは、ユーザーIDとタスクIDを引数に取り、指定されたIDのタスクを取得します。
取得されたタスクは、model.TaskResponseという構造体に変換され、そのタスクを含む構造体が返されます。

CreateTaskメソッドは、新しいタスクを作成します。
新しいタスクは、model.Taskという構造体として受け取り、データベースに保存されます。
保存されたタスクは、model.TaskResponseという構造体に変換され、そのタスクを含む構造体が返されます。

UpdateTaskメソッドは、指定されたIDを持つタスクを更新します。
更新されたタスクは、model.Taskという構造体として受け取り、データベースに保存されます。
保存されたタスクは、model.TaskResponseという構造体に変換され、そのタスクを含む構造体が返されます。

DeleteTaskメソッドは、指定されたIDを持つタスクを削除します。
ユーザーIDとタスクIDが引数として渡され、該当するタスクがデータベースから削除されます。*/