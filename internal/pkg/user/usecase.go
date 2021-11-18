package user

import "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"

type UseCase interface {
	SignUp(signupUser *api_models.SignupUserRequest) (uint64, error)
	LogIn(loginUser *api_models.LoginUserRequest) (uint64, error)
	LogInKeycloak(code string) (uint64, error)
	GetSelfProfile(userId uint64) (*api_models.UserData, error)
	DeleteSelfProfile(userId uint64) error
}
