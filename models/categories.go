package models

import "time"

type Category struct {
	Id        int
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
