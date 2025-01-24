package service

import (
	"encoding/json"
	"fmt"
	"go-api/model"
	"go-api/repository"
	"net/http"
	"os"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetOneUser(id int) (*model.User, error) {
	return s.Repo.GetOneUser(id)
}

func (s *UserService) GetOneUserByNameAndPassword(name string, pass string) (*model.User, error) {
	return s.Repo.GetOneUserByNameAndPassword(name, pass)
}

func (s *UserService) CreateUser(user model.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) UpdateUser(id int, user model.User) error {
	return s.Repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}

func (s *UserService) GetOpenLibraryData() (interface{}, error) {
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
