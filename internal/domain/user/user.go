package user

import (
	"time"
)

type User struct {
	ID        int
	Email     string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
