package session

import "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/models"

type UseCase interface {
	GetUserIdByAccessToken(Uuid string) (uint64, error)
	CreateNewSession(userId uint64) (*models.Token, error)
	DestroySession(Uuid string) error
	RefreshSession(Uuid string) (*models.Token, error)
}
