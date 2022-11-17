package repo

import "time"

type GetPostQuery struct {
	Page   int `json:"page" db:"page" binding:"required" default:"1"`
	Limit  int `json:"limit" db:"limit" binding:"required" default:"10"`
	Search string `json:"search"`
}

type GetAllPostResult struct {
	Post  []*Post
	Count int
}

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

type PostStorageI interface {
	Create(p *Post) (*Post, error)
	Get(id int) (*Post, error)
	GetAll(param GetPostQuery) (*GetAllPostResult, error)
	Update(usr *Post) (*Post, error)
	Delete(id int) error
}
