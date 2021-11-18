package usecase

import (
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/session"
	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/jwt_token"
)

type SessionUseCase struct {
	sessionRepo session.Repository
}

func NewUseCase(sessionRepo session.Repository) session.UseCase {
	return &SessionUseCase{
		sessionRepo: sessionRepo,
	}
}

func (u *SessionUseCase) GetUserIdByAccessToken(Uuid string) (uint64, error) {
	userId, err := u.sessionRepo.SelectUserIdByAccessToken(Uuid)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (u *SessionUseCase) CreateNewSession(userId uint64) (*api_models.Token, error) {
	token, err := jwt_token.CreateJwtToken()
	if err != nil {
		return nil, err
	}

	if err = u.sessionRepo.InsertToken(userId, token); err != nil {
		return nil, err
	}

	return &token.Token, nil
}

func (u *SessionUseCase) DestroySession(Uuid string) error {
	if err := u.sessionRepo.DeleteAccessToken(Uuid); err != nil {
		return err
	}

	if err := u.sessionRepo.DeleteRefreshToken(Uuid); err != nil {
		return err
	}

	return nil
}

func (u *SessionUseCase) RefreshSession(Uuid string) (*api_models.Token, error) {
	userId, err := u.sessionRepo.SelectUserIdByRefreshToken(Uuid)
	if err != nil {
		return nil, err
	}

	if err = u.sessionRepo.DeleteRefreshToken(Uuid); err != nil {
		return nil, err
	}
	_ = u.sessionRepo.DeleteAccessToken(Uuid) // Check access token

	token, err := jwt_token.CreateJwtToken()
	if err != nil {
		return nil, err
	}

	if err = u.sessionRepo.InsertToken(userId, token); err != nil {
		return nil, err
	}

	return &token.Token, nil
}
