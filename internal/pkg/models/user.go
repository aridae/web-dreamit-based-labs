package models

type SignupUserRequest struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	EmailOrLogin string `json:"emailOrLogin"`
	Password     string `json:"password"`
}

type UserData struct {
	Id         uint64 `json:"-"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Login      string `json:"login"`
	Avatar     string `json:"avatar"`
	Background string `json:"background"`
	Password   []byte `json:"-"`
}

type VKAuthResponse struct {
	UserData []VKAuthUserData `json:"response"`
}

type VKAuthUserData struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

type AuthUserData struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ServiceType  string `json:"serviceType"`
	Email        string `json:"email"`
	Login        string `json:"login"`
	Password     []byte `json:"-"`
	AuthId       uint64 `json:"authId"`
}
