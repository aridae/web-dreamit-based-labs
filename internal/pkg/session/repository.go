package session

import "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/models"

type Repository interface {
	InsertToken(userId uint64, token *models.TokenDetails) error
	SelectUserIdByAccessToken(Uuid string) (uint64, error)
	SelectUserIdByRefreshToken(Uuid string) (uint64, error)
	DeleteAccessToken(Uuid string) error
	DeleteRefreshToken(Uuid string) error
}
