package service

import (
	"encoding/json"
	"fmt"
	"go-api/model"
	"go-api/repository"
	"net/http"
	"os"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.Repo.GetAllTasks()
}

func (s *TaskService) GetOneTask(id int) (*model.Task, error) {
	return s.Repo.GetOneTask(id)
}

func (s *TaskService) CreateTask(task model.Task) error {
	return s.Repo.CreateTask(task)
}

func (s *TaskService) UpdateTask(id int, task model.Task) error {
	return s.Repo.UpdateTask(id, task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.Repo.DeleteTask(id)
}

func (s *TaskService) GetOpenLibraryData() (interface{}, error) {
	url := os.Getenv("OPEN_LIBRARY_URL")
	if url == "" {
		return nil, fmt.Errorf("OPEN_LIBRARY_URL is not set")
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return data, nil
}
