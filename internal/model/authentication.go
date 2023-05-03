package model

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

type UserLoginResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Role        string `json:"role"`
	AccessToken string `json:"access_token"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
