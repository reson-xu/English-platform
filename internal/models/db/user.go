package db

import "time"

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Nickname     string
	Role         string
	Status       string
	LastLoginAt  *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
