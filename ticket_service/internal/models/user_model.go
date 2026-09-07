package models

type User struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}
