package models

import "time"

type UserModel struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
