package user

import (
	"time"

	"gorm.io/plugin/optimisticlock"
)

type User struct {
	ID        int
	Email     string
	Age       int
	Version   optimisticlock.Version
	CreatedAt time.Time
	UpdatedAt time.Time
}
