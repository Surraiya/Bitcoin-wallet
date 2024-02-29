package models

import "time"

type Transaction struct {
	Id        string  `gorm:"primary_key;not null"`
	Amount    float64 `gorm:"not null"`
	Spent     bool
	CreatedAt time.Time
}
