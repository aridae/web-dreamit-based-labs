package repository

import (
	"fmt"
	"strconv"

	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/session"

	"github.com/go-redis/redis"
)

type SessionRepository struct {
	connAccessDB  *redis.Client
	connRefreshDB *redis.Client
}

func NewSessionRedisRepository(connAccessDB, connRefreshDB *redis.Client) session.Repository {
	return &SessionRepository{
		connAccessDB:  connAccessDB,
		connRefreshDB: connRefreshDB,
	}
}

func (r *SessionRepository) InsertToken(userId uint64, token *api_models.TokenDetails) error {
	data := fmt.Sprintf("%d", userId)
	err := r.connAccessDB.Set(token.AccessDetails.Uuid, data, api_models.AccessTokenExpires).Err()
	if err != nil {
		return err
	}

	err = r.connRefreshDB.Set(token.RefreshDetails.Uuid, data, api_models.RefreshTokenExpires).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) SelectUserIdByAccessToken(Uuid string) (uint64, error) {
	data, err := r.connAccessDB.Get(Uuid).Bytes()
	if err != nil {
		return 0, err
	}

	userId, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *SessionRepository) SelectUserIdByRefreshToken(Uuid string) (uint64, error) {
	data, err := r.connAccessDB.Get(Uuid).Bytes()
	if err != nil {
		return 0, err
	}

	userId, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *SessionRepository) DeleteAccessToken(Uuid string) error {
	err := r.connAccessDB.Del(Uuid).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) DeleteRefreshToken(Uuid string) error {
	err := r.connRefreshDB.Del(Uuid).Err()
	if err != nil {
		return err
	}

	return nil
}
