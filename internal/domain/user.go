package domain

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"displayname"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"name"`
	Email    string `json:"email"`
}

type UserLoginInfo struct {
	IdToken string `json:"idToken"`
}

type UserTokenClaims struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`	
	DisplayName string `json:"displayName"`
}
