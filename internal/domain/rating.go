package domain

import "time"

type Rating struct {
	ID         string
	ProductID  string
	Type       string
	Value      int
	InsertedAt time.Time
	UpdatedAt  time.Time
}
