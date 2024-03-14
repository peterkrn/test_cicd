package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	UserType int    `json:"type"`
}

type UserResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}
