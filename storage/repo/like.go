package repo

type GetLikesQuery struct {
	Page   int `json:"page" db:"page" binding:"required" default:"1"`
	Limit  int `json:"limit" db:"limit" binding:"required" default:"10"`
	PostId int `json:"post_id"`
	UserId int `json:"user_id"`
}

type GetAllLikesResult struct {
	Like  []*Like
	Count int
}

type Like struct {
	Id     int
	PostId int
	UserId int
	Status string
}

type LikeStorageI interface {
	Create(l *Like) (*Like, error)
	Get(id int) (*Like, error)
	GetAll(param GetLikesQuery) (*GetAllLikesResult, error)
	Update(usr *Like) (*Like, error)
	Delete(id int) error
}
