package model

import (
	"time"
)

type TransactionRespon struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	BookId       int       `json:"book_id"`
	StrartRental time.Time `json:"start_rental"`
	FinishRental time.Time `json:"finish_rental"`
	HavePenalty  int       `json:"have_penalty"`
	Returned     bool      `json:"returned"`
	BookName     string    `json:"book_name"`
	StudentName  string    `json:"student_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type TransactionRequest struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	BookId       int       `json:"book_id"`
	StrartRental time.Time `json:"start_rental"`
	FinishRental time.Time `json:"finish_rental"`
	HavePenalty  int       `json:"have_penalty"`
	Returned     bool      `json:"returned"`
	IsAvailable  bool      `json:"is_available"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
