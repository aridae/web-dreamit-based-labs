package sessionrepo

import (
	"fmt"
	"strconv"

	db "github.com/aridae/web-dreamit-api-based-labs/internal/database"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

type SessionRepository struct {
	AccessDB  *db.RedisClient
	RefreshDB *db.RedisClient
}

func NewSessionRedisRepository(AccessDB, RefreshDB *db.RedisClient) Repository {
	return &SessionRepository{
		AccessDB:  AccessDB,
		RefreshDB: RefreshDB,
	}
}

func (r *SessionRepository) InsertToken(userId uint64, token domain.TokenDetails) error {
	data := fmt.Sprintf("%d", userId)
	err := r.AccessDB.Client.Set(token.AccessDetails.Uuid, data, AccessTokenExpires).Err()
	if err != nil {
		return err
	}

	err = r.RefreshDB.Client.Set(token.RefreshDetails.Uuid, data, RefreshTokenExpires).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) SelectUserIdByAccessToken(Uuid string) (uint64, error) {
	data, err := r.AccessDB.Client.Get(Uuid).Bytes()
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
	data, err := r.AccessDB.Client.Get(Uuid).Bytes()
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
	err := r.AccessDB.Client.Del(Uuid).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) DeleteRefreshToken(Uuid string) error {
	err := r.RefreshDB.Client.Del(Uuid).Err()
	if err != nil {
		return err
	}

	return nil
}
