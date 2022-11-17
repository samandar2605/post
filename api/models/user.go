package models

type User struct {
	Id              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PhoneNumber     string `json:"phone_number"`
	Email           string `json:"email"`
	CreatedAt       string `json:"created_at"`
	Gender          string `json:"gender"`
	Password        string `json:"password"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
	Type            string `json:"type"`
}

type CreateUser struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PhoneNumber     string `json:"phone_number"`
	Email           string `json:"email"`
	CreatedAt       string `json:"created_at"`
	Gender          string `json:"gender"`
	Password        string `json:"password"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
	Type            string `json:"type"`
}
