package entity

import "time"

type Student struct {
	Id          int       `json:"id"`
	StudentName string    `json:"student_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
