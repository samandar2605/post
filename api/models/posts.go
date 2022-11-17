package models

import "time"

type Post struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	ImageUrl    string    `json:"image_url" db:"image_url"`
	UserId      string    `json:"user_id" db:"user_id"`
	CategoryId  string    `json:"category_id" db:"category_id"`
	UpdatedAt   string    `json:"updated_at" db:"updated_at"`
	ViewsCount  string    `json:"views_count" db:"views_count"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type CreatePost struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	ImageUrl    string `json:"image_url" db:"image_url"`
	UserId      string `json:"user_id" db:"user_id"`
	CategoryId  string `json:"category_id" db:"category_id"`
	ViewsCount  string `json:"views_count" db:"views_count"`
}
