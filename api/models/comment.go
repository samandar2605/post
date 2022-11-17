package models

import "time"

type Comment struct {
	Id          int       `json:"id" db:"id"`
	PostId      int       `json:"post_id" db:"post_id"`
	UserId      int       `json:"user_id" db:"user_id"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateComment struct {
	PostId      int    `json:"post_id" db:"post_id"`
	UserId      int    `json:"user_id" db:"user_id"`
	Description string `json:"description" db:"description"`
}
