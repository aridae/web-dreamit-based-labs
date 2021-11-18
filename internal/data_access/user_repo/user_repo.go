package userrepo

import "github.com/aridae/web-dreamit-api-based-labs/internal/domain"

type Repository interface {
	GetUsers() ([]domain.UserProfile, error)
	SelectUserById(userId uint64) (*domain.UserProfile, error)

	InsertUser(user *domain.UserData) (uint64, error)
	InsertAuthUser(user *domain.AuthUserData) (uint64, error)
	SelectUserByEmailOrLogin(emailOrLogin string) (*domain.UserData, error)
	SelectUserByAuthId(authId uint64, authService string) (*domain.UserData, error)
	SelectNewUniqLogin(login string) (string, error)
	DeleteSelfProfile(userId uint64) error
}
