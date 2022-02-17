package user

import "time"

type User struct {
	ID           int
	Name         string
	Email        string
	Password     string
	SecretKey    string
	LinkResource string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
