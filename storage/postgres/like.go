package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/post/storage/repo"
)

type likeRepo struct {
	db *sqlx.DB
}

func NewLike(db *sqlx.DB) repo.LikeStorageI {
	return &likeRepo{db: db}
}

func (cr *likeRepo) Create(like *repo.Like) (*repo.Like, error) {
	query := `
		INSERT INTO likes(
			post_id,
			user_id,
			status
		) values ($1,$2,$3)
		RETURNING id
	`
	result := cr.db.QueryRow(
		query,
		like.PostId,
		like.UserId,
		like.Status,
	)
	if err := result.Scan(
		&like.Id,
	); err != nil {
		return nil, err
	}
	return like, nil
}

func (cr *likeRepo) Get(id int) (*repo.Like, error) {
	var like repo.Like
	query := `
		SELECT
			id,
			post_id,
			user_id,
			status
		FROM likes
		WHERE id=$1
	`

	result := cr.db.QueryRow(
		query,
		id,
	)
	if err := result.Scan(
		&like.Id,
		&like.PostId,
		&like.UserId,
		&like.Status,
	); err != nil {
		return nil, err
	}

	return &like, nil
}

func (cr *likeRepo) GetAll(param repo.GetLikesQuery) (*repo.GetAllLikesResult, error) {
	result := repo.GetAllLikesResult{
		Like: make([]*repo.Like, 0),
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
			status
		FROM likes
		` + filter + `
		ORDER BY post_id desc
		` + limit

	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var like repo.Like
		if err := rows.Scan(
			&like.Id,
			&like.PostId,
			&like.UserId,
			&like.Status,
		); err != nil {
			return nil, err
		}
		result.Like = append(result.Like, &like)
	}
	queryCount := `SELECT count(1) FROM likes ` + filter
	err = cr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr *likeRepo) Update(like *repo.Like) (*repo.Like, error) {
	query := `
		update likes set 
			post_id =$1,
			user_id =$2,
			status=$3
		where id=$4
		RETURNING id
	`
	result := cr.db.QueryRow(
		query,
		like.PostId,
		like.UserId,
		like.Status,
	)

	if err := result.Scan(
		&like.Id,
	); err != nil {
		return nil, err
	}

	return like, nil
}

func (cr *likeRepo) Delete(id int) error {
	res, err := cr.db.Exec("delete from likes where id=$1", id)
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
