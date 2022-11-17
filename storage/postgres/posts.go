package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/post/storage/repo"
)

type postRepo struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) repo.PostStorageI {
	return &postRepo{db: db}
}

func (pr *postRepo) Create(p *repo.Post) (*repo.Post, error) {
	query := `
		INSERT INTO posts(
			title,
			description,
			image_url,
			user_id,
			category_id,
			views_count
		)values($1,$2,$3,$4,$5,$6)
		RETURNING id,created_at
	`
	row := pr.db.QueryRow(
		query,
		p.Title,
		p.Description,
		p.ImageUrl,
		p.UserId,
		p.CategoryId,
		p.ViewsCount,
	)

	if err := row.Scan(
		&p.Id,
		&p.CreatedAt,
	); err != nil {
		return nil, err
	}

	return p, nil
}

func (pr *postRepo) Get(id int) (*repo.Post, error) {
	var Post repo.Post

	query := `
		SELECT 
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			views_count,
			created_at
		from posts
		where id=$1
	`
	row := pr.db.QueryRow(query, id)
	if err := row.Scan(
		&Post.Id,
		&Post.Title,
		&Post.Description,
		&Post.ImageUrl,
		&Post.UserId,
		&Post.CategoryId,
		&Post.ViewsCount,
		&Post.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &Post, nil
}

func (pr *postRepo) GetAll(param repo.GetPostQuery) (*repo.GetAllPostResult, error) {
	result := repo.GetAllPostResult{
		Post: make([]*repo.Post, 0),
	}

	offset := (param.Page - 1) * param.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", param.Limit, offset)
	filter := ""
	if param.Search != "" {
		str := "%" + param.Search + "%"
		filter += fmt.Sprintf(`
			where title ILIKE '%s'`, str)
	}

	query := `
		SELECT 
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			views_count,
			created_at
		FROM posts
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var Post repo.Post
		if err := rows.Scan(
			&Post.Id,
			&Post.Title,
			&Post.Description,
			&Post.ImageUrl,
			&Post.UserId,
			&Post.CategoryId,
			&Post.ViewsCount,
			&Post.CreatedAt,
		); err != nil {
			return nil, err
		}
		result.Post = append(result.Post, &Post)
	}
	queryCount := `SELECT count(1) FROM posts ` + filter
	err = pr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (pr *postRepo) Update(post *repo.Post) (*repo.Post, error) {
	query := `
		update posts set 
			title=$1,
			description=$2,
			image_url=$3,
			user_id=$4,
			category_id=$5,
			views_count=$6,
			updated_at=$7
		where id=$8
		RETURNING updated_at
	`
	row := pr.db.QueryRow(
		query,
		post.Title,
		post.Description,
		post.ImageUrl,
		post.UserId,
		post.CategoryId,
		post.ViewsCount,
		time.Now(),
		post.Id,
	)

	if err := row.Scan(&post.UpdatedAt); err != nil {
		return nil, err
	}

	return post, nil
}

func (ur *postRepo) Delete(id int) error {
	res, err := ur.db.Exec("delete from posts where id=$1", id)
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
