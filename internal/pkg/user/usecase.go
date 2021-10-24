package user

import (
	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/internal/pkg/models"
)

type UseCase interface {
	SignUp(signupUser *models.SignupUserRequest) (uint64, error)
	LogIn(loginUser *models.LoginUserRequest) (uint64, error)
	LogInKeycloak(code string) (uint64, error)
	GetSelfProfile(userId uint64) (*models.UserData, error)
	DeleteSelfProfile(userId uint64) error
}
