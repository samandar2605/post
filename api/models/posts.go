package models

import "time"

type Post struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	ImageUrl    string    `json:"image_url" db:"image_url"`
	UserId      int       `json:"user_id" db:"user_id"`
	CategoryId  int       `json:"category_id" db:"category_id"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	ViewsCount  int       `json:"views_count" db:"views_count"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type CreatePost struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	ImageUrl    string `json:"image_url" db:"image_url"`
	UserId      int `json:"user_id" db:"user_id"`
	CategoryId  int `json:"category_id" db:"category_id"`
	ViewsCount  int `json:"views_count" db:"views_count"`
}
