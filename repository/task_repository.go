package repository

import (
	"go-api/model"
	"go-api/utils"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	DB *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.DB.Select(&tasks, utils.GetAllTasksQuery)
	return tasks, err
}

func (r *TaskRepository) CreateTask(task model.Task) error {
	_, err := r.DB.Exec(utils.CreateTaskQuery, task.Title, task.Description, task.UserId)
	return err
}

func (r *TaskRepository) UpdateTask(id int, task model.Task) error {
	_, err := r.DB.Exec(utils.UpdateTaskQuery, task.Title, task.Description, task.UserId, id)
	return err
}

func (r *TaskRepository) GetOneTask(id int) (*model.Task, error) {
	var task model.Task
	err := r.DB.Get(&task, utils.GetOneTaskQuery, id)
	return &task, err
}

func (r *TaskRepository) DeleteTask(id int) error {
	_, err := r.DB.Exec(utils.DeleteTaskQuery, id)
	return err
}
