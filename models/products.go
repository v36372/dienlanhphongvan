package models

import "time"

type Product struct {
	Id          int
	Name        string
	Description string
	CategoryId  int
	Price       float32
	Slug        string
	Thumbnail   string
	Image01     string
	Image02     string
	Image03     string
	Image04     string
	Image05     string
	Image06     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
