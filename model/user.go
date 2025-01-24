package model

type User struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
	Position string `db:"position" json:"position"`
	Salary   int    `db:"salary" json:"salary"`
}
