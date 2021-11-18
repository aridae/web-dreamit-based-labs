package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	sessionrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/session_repo"
	domain "github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/configer"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type TokenType int

const (
	Access TokenType = iota
	Refresh
)

type SessionController struct {
	SessionRepo sessionrepo.Repository
}

func NewSessionController(SessionRepo sessionrepo.Repository) *SessionController {
	return &SessionController{
		SessionRepo: SessionRepo,
	}
}

func (u *SessionController) GetUserIdByAccessToken(Uuid string) (uint64, error) {
	userId, err := u.SessionRepo.SelectUserIdByAccessToken(Uuid)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (u *SessionController) CreateNewSession(userId uint64) (*domain.Token, error) {
	token, err := u.CreateJwtToken()
	if err != nil {
		return nil, err
	}

	if err = u.SessionRepo.InsertToken(userId, *token); err != nil {
		return nil, err
	}

	return &domain.Token{
		AccessToken:  token.Token.AccessToken,
		RefreshToken: token.Token.RefreshToken,
	}, nil
}

func (u *SessionController) DestroySession(Uuid string) error {
	if err := u.SessionRepo.DeleteAccessToken(Uuid); err != nil {
		return err
	}

	if err := u.SessionRepo.DeleteRefreshToken(Uuid); err != nil {
		return err
	}

	return nil
}

func (u *SessionController) RefreshSession(Uuid string) (*domain.Token, error) {
	userId, err := u.SessionRepo.SelectUserIdByRefreshToken(Uuid)
	if err != nil {
		return nil, err
	}

	if err = u.SessionRepo.DeleteRefreshToken(Uuid); err != nil {
		return nil, err
	}
	_ = u.SessionRepo.DeleteAccessToken(Uuid) // Check access token

	token, err := u.CreateJwtToken()
	if err != nil {
		return nil, err
	}

	if err = u.SessionRepo.InsertToken(userId, *token); err != nil {
		return nil, err
	}

	return &domain.Token{
		AccessToken:  token.Token.AccessToken,
		RefreshToken: token.Token.RefreshToken,
	}, nil
}

func (u *SessionController) CreateJwtToken() (*domain.TokenDetails, error) {
	tokenDetails := &domain.TokenDetails{}
	uuid, _ := uuid.NewV4()
	Uuid := uuid.String()

	// Setting domain.Token Details
	tokenDetails.AccessDetails.Expires = time.Now().Add(domain.AccessTokenExpires).Unix()
	tokenDetails.AccessDetails.Uuid = Uuid

	tokenDetails.RefreshDetails.Expires = time.Now().Add(domain.RefreshTokenExpires).Unix()
	tokenDetails.RefreshDetails.Uuid = Uuid

	var err error
	// Creating Access domain.Token
	os.Setenv("ACCESS_SECRET", configer.AppConfig.Secret.AccessSecret)
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["user_uuid"] = tokenDetails.AccessDetails.Uuid
	accessTokenClaims["expires"] = tokenDetails.AccessDetails.Expires
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	tokenDetails.Token.AccessToken, err = accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	// Creating Refresh domain.Token
	os.Setenv("REFRESH_SECRET", configer.AppConfig.Secret.RefreshSecret)
	RefreshTokenClaims := jwt.MapClaims{}
	RefreshTokenClaims["user_uuid"] = tokenDetails.RefreshDetails.Uuid
	RefreshTokenClaims["expires"] = tokenDetails.RefreshDetails.Expires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshTokenClaims)
	tokenDetails.Token.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
}

func (u *SessionController) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return strArr[0]
}

func (u *SessionController) VerifyToken(r *http.Request, tokenType TokenType) (*jwt.Token, error) {
	tokenString := u.ExtractToken(r)

	var tokenSecret string
	switch tokenType {
	case Access:
		tokenSecret = configer.AppConfig.Secret.AccessSecret
	case Refresh:
		tokenSecret = configer.AppConfig.Secret.RefreshSecret
	default:
		return nil, fmt.Errorf("incorrect token type")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (u *SessionController) TokenValid(r *http.Request, tokenType TokenType) error {
	_, err := u.VerifyToken(r, tokenType)
	if err != nil {
		return err
	}

	// if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
	// 	return err
	// }

	return nil
}

func (u *SessionController) ExtractAccessTokenMetadata(r *http.Request) (*domain.AccessTokenDetails, error) {
	token, err := u.VerifyToken(r, Access)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["user_uuid"].(string)
		if !ok {
			return nil, err
		}
		return &domain.AccessTokenDetails{
			Uuid: accessUuid,
		}, nil
	}
	return nil, err
}

func (u *SessionController) ExtractRefreshTokenMetadata(r *http.Request) (*domain.RefreshTokenDetails, error) {
	token, err := u.VerifyToken(r, Refresh)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["user_uuid"].(string)
		if !ok {
			return nil, err
		}
		return &domain.RefreshTokenDetails{
			Uuid: accessUuid,
		}, nil
	}
	return nil, err
}
