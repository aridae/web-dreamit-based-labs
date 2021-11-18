package userrepo

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
