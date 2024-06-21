package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
type UserResponseError struct {
	Error string `json:"error"`
}
