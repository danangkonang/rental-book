package entity

import "time"

type Book struct {
	Id            int       `json:"id"`
	BookName      string    `json:"book_name"`
	MaxRentalDays int       `json:"max_rental_days"`
	PenaltyPerDay int       `json:"penalty_per_day"`
	IsAvailable   bool      `json:"is_available"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
