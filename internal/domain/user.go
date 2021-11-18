package domain

// модели уровня бизнес логики

type UserData struct {
	Id         uint64
	FirstName  string
	LastName   string
	Email      string
	Login      string
	Avatar     string
	Background string
	Password   []byte
}

type UserProfile struct {
	Id         uint64
	FirstName  string
	LastName   string
	Email      string
	Login      string
	Avatar     string
	Background string
}

type AuthUserData struct {
	FirstName    string
	LastName     string
	AccessToken  string
	RefreshToken string
	ServiceType  string
	Email        string
	Login        string
	Password     []byte
	AuthId       uint64
}

type SignupUserData struct {
	Email    string
	Login    string
	Password string
}

type LoginUserData struct {
	EmailOrLogin string
	Password     string
}
