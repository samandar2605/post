package models

import "time"

type Category struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateCategory struct {
	Title string `json:"title" binding:"required"`
}
