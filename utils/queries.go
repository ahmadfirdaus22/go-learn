package utils

const (
	GetAllUsersQuery                 = `SELECT * FROM employee`
	CreateUserQuery                  = `INSERT INTO employee (name, position, password, salary) VALUES ($1, $2, $3, $4)`
	UpdateUserQuery                  = `UPDATE employee SET name = $1, position = $2, salary = $3 WHERE id = $4`
	DeleteUserQuery                  = `DELETE FROM employee WHERE id=$1`
	GetOneUserQuery                  = `SELECT * FROM employee where id=$1`
	GetOneUserByNameAndPasswordQuery = `SELECT * FROM employee where name=$1 and password = $2`
	GetAllTasksQuery                 = `SELECT * FROM tasks`
	CreateTaskQuery                  = `INSERT INTO tasks (title, description, user_id) VALUES ($1, $2, $3)`
	UpdateTaskQuery                  = `UPDATE tasks SET title = $1, description = $2, user_id = $3 WHERE id = $4`
	DeleteTaskQuery                  = `DELETE FROM tasks WHERE id=$1`
	GetOneTaskQuery                  = `SELECT * FROM tasks where id=$1`
)
