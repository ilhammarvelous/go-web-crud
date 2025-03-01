package entities

import "time"

type Product struct {
	Id          uint
	Name        string
	Category    Category
	Stock       int64
	Description string
	CreateAt    time.Time
	UpdatedAt   time.Time
	No          int
}
