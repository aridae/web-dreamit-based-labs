package sessionrepo

import "github.com/aridae/web-dreamit-api-based-labs/internal/domain"

type Repository interface {
	InsertToken(userId uint64, token domain.TokenDetails) error
	SelectUserIdByAccessToken(Uuid string) (uint64, error)
	SelectUserIdByRefreshToken(Uuid string) (uint64, error)
	DeleteAccessToken(Uuid string) error
	DeleteRefreshToken(Uuid string) error
}
