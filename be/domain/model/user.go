package model

import (
	"time"
)

type UserID uint64

type User struct {
	ID        UserID
	Name      string
	CreatedAt time.Time
}
