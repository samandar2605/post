package repo

type GetLikesQuery struct {
	Page   int
	Limit  int
	PostId int
	UserId int
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
