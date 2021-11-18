package session

import "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"

type Repository interface {
	InsertToken(userId uint64, token *api_models.TokenDetails) error
	SelectUserIdByAccessToken(Uuid string) (uint64, error)
	SelectUserIdByRefreshToken(Uuid string) (uint64, error)
	DeleteAccessToken(Uuid string) error
	DeleteRefreshToken(Uuid string) error
}
