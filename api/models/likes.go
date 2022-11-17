package models

type Like struct {
	Id     int    `json:"id"`
	PostId int    `json:"post_id"`
	UserId int    `json:"user_id"`
	Status string `json:"status"`
}

type CreateLike struct {
	PostId int    `json:"post_id"`
	UserId int    `json:"user_id"`
	Status string `json:"status"`
}
