package user

import "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/models"

type Repository interface {
	InsertUser(user *models.UserData) (uint64, error)
	InsertAuthUser(user *models.AuthUserData) (uint64, error)
	SelectUserByEmailOrLogin(emailOrLogin string) (*models.UserData, error)
	SelectUserById(userId uint64) (*models.UserData, error)
	SelectUserByAuthId(authId uint64, authService string) (*models.UserData, error)
	SelectNewUniqLogin(login string) (string, error)
	DeleteSelfProfile(userId uint64) error
}
