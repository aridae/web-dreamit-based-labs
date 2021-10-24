package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"gopkg.in/square/go-jose.v2/jwt"
	"io/ioutil"
	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/configer"
	"net/http"
	net_rul "net/url"

	"fmt"

	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/models"
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/user"
	//"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/configer"
	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/hasher"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type UserUseCase struct {
	userRepo user.Repository
}

func NewUseCase(userRepo user.Repository) user.UseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func StringToUid(str string) uint64 {
	var sum uint64
	for _, char := range str {
		sum += uint64(char)
	}
	return sum
}

func (u *UserUseCase) SignUp(signupUser *models.SignupUserRequest) (uint64, error) {
	hashOfPassword, err := hasher.GenerateHashFromPassword(signupUser.Password)
	if err != nil {
		return 0, err
	}

	userId, err := u.userRepo.InsertUser(&models.UserData{
		Email:    signupUser.Email,
		Login:    signupUser.Login,
		Password: hashOfPassword,
	})
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (u *UserUseCase) LogIn(loginUser *models.LoginUserRequest) (uint64, error) {
	fmt.Println(loginUser)
	url := "http://keycloak:8080/auth/realms/demo/protocol/openid-connect/token?"
	params := net_rul.Values{}
	params.Add("client_id",  configer.AppConfig.OAuth2App.Keycloak.ClientID)
	params.Add("grant_type", "password")
	params.Add("client_secret", configer.AppConfig.OAuth2App.Keycloak.ClientSecret)
	params.Add("scope", "openid")
	params.Add("username", loginUser.EmailOrLogin)
	params.Add("password", loginUser.Password)

	jsonStr := []byte("")
	fmt.Println(jsonStr)
	fmt.Println(bytes.NewBuffer(jsonStr))
	fmt.Println( url+params.Encode())
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(params.Encode())))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	var token oauth2.Token
	err = json.Unmarshal(body, &token)

	var claims map[string]interface{}

	parsedToken, _ := jwt.ParseSigned(token.AccessToken)
	_ = parsedToken.UnsafeClaimsWithoutVerification(&claims)

	authUserId := StringToUid(claims["sub"].(string))
	userData, err := u.userRepo.SelectUserByAuthId(authUserId, "keycloak")
	if err == nil {
		return userData.Id, nil
	}

	newLogin, err := u.userRepo.SelectNewUniqLogin(claims["sub"].(string))
	if err != nil {
		return 0, err
	}

	userId, err := u.userRepo.InsertAuthUser(&models.AuthUserData{
		FirstName:    newLogin,
		LastName:     newLogin,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ServiceType:  "keycloak",
		Login:        newLogin,
		AuthId:       authUserId,
	})
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (u *UserUseCase) LogInKeycloak(code string) (uint64, error) {
	fmt.Println(code)
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, configer.AppConfig.OAuth2App.Keycloak.ConfigURL)
	if err != nil {
		panic(err)
	}

	oauth2Config := oauth2.Config{
		ClientID:     configer.AppConfig.OAuth2App.Keycloak.ClientID,
		ClientSecret: configer.AppConfig.OAuth2App.Keycloak.ClientSecret,
		RedirectURL:  configer.AppConfig.OAuth2App.Keycloak.RedirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	state := "somestate"

	fmt.Println(oauth2Config.AuthCodeURL(state))
	token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		return 0, err
	}

	client := oauth2Config.Client(ctx, token)
	fmt.Println(client)

	var claims map[string]interface{}

	parsedToken, _ := jwt.ParseSigned(token.AccessToken)
	_ = parsedToken.UnsafeClaimsWithoutVerification(&claims)

	authUserId := StringToUid(claims["sub"].(string))
	userData, err := u.userRepo.SelectUserByAuthId(authUserId, "keycloak")
	if err == nil {
		return userData.Id, nil
	}

	newLogin, err := u.userRepo.SelectNewUniqLogin(claims["sub"].(string))
	if err != nil {
		return 0, err
	}

	userId, err := u.userRepo.InsertAuthUser(&models.AuthUserData{
		FirstName:    newLogin,
		LastName:     newLogin,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ServiceType:  "keycloak",
		Login:        newLogin,
		AuthId:       authUserId,
	})
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (u *UserUseCase) GetSelfProfile(userId uint64) (*models.UserData, error) {
	userData, err := u.userRepo.SelectUserById(userId)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (u *UserUseCase) DeleteSelfProfile(userId uint64) error {
	err := u.userRepo.DeleteSelfProfile(userId)
	if err != nil {
		return err
	}

	return nil
}
