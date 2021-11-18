package session

import "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"

type UseCase interface {
	GetUserIdByAccessToken(Uuid string) (uint64, error)
	CreateNewSession(userId uint64) (*api_models.Token, error)
	DestroySession(Uuid string) error
	RefreshSession(Uuid string) (*api_models.Token, error)
}
