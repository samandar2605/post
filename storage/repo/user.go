package repo

import "time"

type User struct {
	Id              int       `db:"id"`
	FirstName       string    `db:"first_name"`
	LastName        string    `db:"last_name"`
	PhoneNumber     string    `db:"phone_number"`
	Email           string    `db:"email"`
	Gender          string    `db:"gender"`
	UserName        string    `db:"user_name"`
	Password        string    `db:"password"`
	ProfileImageUrl string    `db:"profile_image_url"`
	Type            string    `db:"type"`
	CreatedAt       time.Time `db:"created_at"`
}

type UserStorageI interface {
	Create(u *User) (*User, error)
	Get(id int) (*User, error)
	GetAll(param GetUserQuery) (*GetAllUsersResult, error)
	Update(usr *User) (*User, error)
	Delete(id int)error
}

type GetUserQuery struct {
	Page  int
	Limit int
	Search string
}

type GetAllUsersResult struct {
	Users []*User
	Count int
}
