package sessionrepo

import "time"

var (
	AccessTokenExpires  = time.Hour * 15
	RefreshTokenExpires = time.Hour * 24 * 7
)

type TokenDetails struct {
	Token          Token
	AccessDetails  AccessTokenDetails
	RefreshDetails RefreshTokenDetails
}

type AccessTokenDetails struct {
	Uuid    string
	Expires int64
}

type RefreshTokenDetails struct {
	Uuid    string
	Expires int64
}

type Token struct {
	AccessToken  string
	RefreshToken string
}
