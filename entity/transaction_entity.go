package entity

import "time"

type Transaction struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	BookId       int       `json:"book_id"`
	StrartRental time.Time `json:"start_rental"`
	FinishRental time.Time `json:"finish_rental"`
	HavePenalty  int       `json:"have_penalty"`
	Returned     bool      `json:"returned"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
