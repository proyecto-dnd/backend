package domain

type User struct {
	Id       string `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
}