package repo

import "time"

type Category struct {
	Id        int
	Title     string
	CreatedAt time.Time
}

type CategoryStorageI interface {
	Create(u *Category) (*Category, error)
	Get(id int) (*Category, error)
	GetAll(param GetCategoryQuery) (*GetAllCategoriesResult, error)
	Update(category *Category) (*Category, error)
	Delete(id int) error
}

type GetCategoryQuery struct {
	Page   int `json:"page" db:"page" binding:"required" default:"1"`
	Limit  int `json:"limit" db:"limit" binding:"required" default:"10"`
	Search string `json:"search"`
}

type GetAllCategoriesResult struct {
	Categories []*Category
	Count      int
}
