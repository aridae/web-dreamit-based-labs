package session

import "lab.qoollo.com/practice/2021/dreamit/dreamit-api/internal/pkg/models"

type Repository interface {
	InsertToken(userId uint64, token *models.TokenDetails) error
	SelectUserIdByAccessToken(Uuid string) (uint64, error)
	SelectUserIdByRefreshToken(Uuid string) (uint64, error)
	DeleteAccessToken(Uuid string) error
	DeleteRefreshToken(Uuid string) error
}
