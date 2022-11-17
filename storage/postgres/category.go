package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/post/storage/repo"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategory(db *sqlx.DB) repo.CategoryStorageI {
	return &categoryRepo{
		db: db,
	}
}

func (cr *categoryRepo) Create(category *repo.Category) (*repo.Category, error) {
	query := `
		INSERT INTO categories(title) VALUES($1)
		RETURNING id, created_at
	`

	row := cr.db.QueryRow(
		query,
		category.Title,
	)

	err := row.Scan(
		&category.Id,
		&category.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cr *categoryRepo) Get(id int) (*repo.Category, error) {
	var result repo.Category

	query := `
		SELECT
			id,
			title,
			created_at
		FROM categories
		WHERE id=$1
	`

	row := cr.db.QueryRow(query, id)
	err := row.Scan(
		&result.Id,
		&result.Title,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cr *categoryRepo) GetAll(param repo.GetCategoryQuery) (*repo.GetAllCategoriesResult, error) {
	result := repo.GetAllCategoriesResult{
		Categories: make([]*repo.Category, 0),
	}

	offset := (param.Page - 1) * param.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", param.Limit, offset)
	filter := "where true"
	if param.Search != "" {
		str := "%" + param.Search + "%"
		filter += fmt.Sprintf(` 
			and title ILIKE '%s'`, str)
	}

	query := `
		SELECT 
			id,
			title,
			created_at
		FROM categories
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var Categ repo.Category
		if err := rows.Scan(
			&Categ.Id,
			&Categ.Title,
			&Categ.CreatedAt,
		); err != nil {
			return nil, err
		}
		result.Categories = append(result.Categories, &Categ)
	}
	queryCount := `SELECT count(1) FROM categories ` + filter
	err = cr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return &result, nil
}

func (cr *categoryRepo) Update(category *repo.Category) (*repo.Category, error) {
	query := `
		update categories set
			title=$1
		where id=$2
	`
	_, err := cr.db.Exec(query, category.Title, category.Id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (ur *categoryRepo) Delete(id int) error {
	res, err := ur.db.Exec("delete from categories where id=$1", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
