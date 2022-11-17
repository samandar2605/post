package models

type GetAllParams struct {
	Limit  int    `json:"limit" binding:"required" default:"10"`
	Page   int    `json:"page" binding:"required" default:"1"`
	Search string `json:"search"`
}
