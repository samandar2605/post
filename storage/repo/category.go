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
	Update(category Category) (*Category, error)
	Delete(id int) error
}

type GetCategoryQuery struct {
	Page  int
	Limit int
	Search string
}

type GetAllCategoriesResult struct {
	Categories []*Category
	Count      int
}
