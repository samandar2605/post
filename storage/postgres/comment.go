package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/samandar2605/post/storage/repo"
)

type commentRepo struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) repo.CommentStorageI {
	return &commentRepo{db: db}
}

func (cr *commentRepo) Create(comment *repo.Comment) (*repo.Comment, error) {
	query := `
		INSERT INTO comments(
			post_id,
			user_id,
			description,
			created_at
		) values ($1,$2,$3,$4)
		RETURNING
			id,
			created_at
	`
	result := cr.db.QueryRow(
		query,
		comment.PostId,
		comment.UserId,
		comment.Description,
		time.Now(),
	)
	if err := result.Scan(
		&comment.Id,
		&comment.CreatedAt,
	); err != nil {
		return nil, err
	}
	return comment, nil
}

func (cr *commentRepo) Get(id int) (*repo.Comment, error) {
	var Comment repo.Comment
	query := `
		SELECT
			id,
			post_id,
			user_id,
			description,
			created_at
		FROM comments
		WHERE id=$1
	`

	result := cr.db.QueryRow(
		query,
		id,
	)
	if err := result.Scan(
		&Comment.Id,
		&Comment.PostId,
		&Comment.UserId,
		&Comment.Description,
		&Comment.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &Comment, nil
}

func (cr *commentRepo) GetAll(param repo.GetCommentQuery) (*repo.GetAllCommentsResult, error) {
	result := repo.GetAllCommentsResult{
		Comments: make([]*repo.Comment, 0),
	}

	offset := (param.Page - 1) * param.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", param.Limit, offset)
	filter := ""
	if param.PostId > 0 {
		filter += fmt.Sprintf("where post_id=%d", param.PostId)
	}
	if param.UserId > 0 {
		if filter == "" {
			filter += fmt.Sprintf("where user_id=%d", param.PostId)
		} else {
			filter += fmt.Sprintf("or user_id=%d", param.PostId)
		}
	}
	query := `
		SELECT 
			id,
			post_id,
			user_id,
			description,
			created_at
		FROM comments
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var Comment repo.Comment
		if err := rows.Scan(
			&Comment.Id,
			&Comment.PostId,
			&Comment.UserId,
			&Comment.Description,
			&Comment.CreatedAt,
		); err != nil {
			return nil, err
		}
		result.Comments = append(result.Comments, &Comment)
	}
	queryCount := `SELECT count(1) FROM comments ` + filter
	err = cr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr *commentRepo) Update(comment *repo.Comment) (*repo.Comment, error) {
	query := `
		update comments set 
			post_id=$1,
			user_id=$2,
			description=$3,
			updated_at=$4
		where id=$5
		RETURNING
			id,
			created_at,
			updated_at
	`
	result := cr.db.QueryRow(
		query,
		comment.PostId,
		comment.UserId,
		comment.Description,
		time.Now(),
		comment.Id,
	)

	if err := result.Scan(
		&comment.Id,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return comment, nil
}

func (cr *commentRepo) Delete(id int) error {
	res, err := cr.db.Exec("delete from comments where id=$1", id)
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
