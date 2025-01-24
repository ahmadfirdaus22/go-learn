package model

type Task struct {
	ID          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"name"`
	Description string `db:"description" json:"description"`
	Completed   bool   `db:"completed" json:"completed"`
	UserId      int    `db:"user_id" json:"user_id"`
}
