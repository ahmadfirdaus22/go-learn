package repository

import (
	"go-api/model"
	"go-api/utils"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.DB.Select(&users, utils.GetAllUsersQuery)
	return users, err
}

func (r *UserRepository) CreateUser(user model.User) error {
	_, err := r.DB.Exec(utils.CreateUserQuery, user.Name, user.Position, user.Password, user.Salary)
	return err
}

func (r *UserRepository) UpdateUser(id int, user model.User) error {
	_, err := r.DB.Exec(utils.UpdateUserQuery, user.Name, user.Position, user.Salary, id)
	return err
}

func (r *UserRepository) GetOneUser(id int) (*model.User, error) {
	var user model.User
	err := r.DB.Get(&user, utils.GetOneUserQuery, id)
	return &user, err
}

func (r *UserRepository) GetOneUserByNameAndPassword(name string, pass string) (*model.User, error) {
	var user model.User
	err := r.DB.Get(&user, utils.GetOneUserByNameAndPasswordQuery, name, pass)
	return &user, err
}

func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.DB.Exec(utils.DeleteUserQuery, id)
	return err
}
