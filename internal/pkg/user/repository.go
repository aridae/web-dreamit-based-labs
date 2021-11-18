package user

import "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"

type Repository interface {
	InsertUser(user *api_models.UserData) (uint64, error)
	InsertAuthUser(user *api_models.AuthUserData) (uint64, error)
	SelectUserByEmailOrLogin(emailOrLogin string) (*api_models.UserData, error)
	SelectUserById(userId uint64) (*api_models.UserData, error)
	SelectUserByAuthId(authId uint64, authService string) (*api_models.UserData, error)
	SelectNewUniqLogin(login string) (string, error)
	DeleteSelfProfile(userId uint64) error
}
