package models

import "time"

type Category struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
