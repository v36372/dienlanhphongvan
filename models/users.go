package models

import "time"

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
