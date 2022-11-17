package repo

import "time"

type GetPostQuery struct {
	Page  int
	Limit int
	Search string
}

type GetAllPostResult struct {
	Post  []*Post
	Count int
}

type Post struct {
	Id          int
	Title       string
	Description string
	ImageUrl    string
	UserId      string
	CategoryId  string
	UpdatedAt   string
	ViewsCount  string
	CreatedAt   time.Time
}

type PostStorageI interface {
	Create(p *Post) (*Post, error)
	Get(id int) (*Post, error)
	GetAll(param GetPostQuery) (*GetAllPostResult, error)
	Update(usr *Post) (*Post, error)
	Delete(id int) error
}


