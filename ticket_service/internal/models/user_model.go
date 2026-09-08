package models

import "uuid"

type User struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}
