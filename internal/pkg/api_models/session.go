package api_models

import "time"

var (
	AccessTokenExpires  = time.Hour * 15
	RefreshTokenExpires = time.Hour * 24 * 7
)

type TokenDetails struct {
	Token          Token               `json:"token"`
	AccessDetails  AccessTokenDetails  `json:"accessDetails"`
	RefreshDetails RefreshTokenDetails `json:"refreshDetails"`
}

type AccessTokenDetails struct {
	Uuid    string `json:"uUid"`
	Expires int64  `json:"expires"`
}

type RefreshTokenDetails struct {
	Uuid    string `json:"uUid"`
	Expires int64  `json:"expires"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
