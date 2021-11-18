package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	net_rul "net/url"

	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/configer"
	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/hasher"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"gopkg.in/square/go-jose.v2/jwt"

	userrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/user_repo"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

type UserController struct {
	UserRepo userrepo.Repository
}

func NewUserController(UserRepo userrepo.Repository) *UserController {
	return &UserController{
		UserRepo: UserRepo,
	}
}

func StringToUid(str string) uint64 {
	var sum uint64
	for _, char := range str {
		sum += uint64(char)
	}
	return sum
}

func (u *UserController) SignUp(signupUser *domain.SignupUserData) (uint64, error) {
	hashOfPassword, err := hasher.GenerateHashFromPassword(signupUser.Password)
	if err != nil {
		return 0, err
	}

	userId, err := u.UserRepo.InsertUser(&domain.UserData{
		Email:    signupUser.Email,
		Login:    signupUser.Login,
		Password: hashOfPassword,
	})
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (u *UserController) LogIn(loginUser *domain.LoginUserData) (uint64, error) {
	fmt.Println(loginUser)
	url := "http://keycloak:8080/auth/realms/demo/protocol/openid-connect/token?"
	params := net_rul.Values{}
	params.Add("client_id", configer.AppConfig.OAuth2App.Keycloak.ClientID)
	params.Add("grant_type", "password")
	params.Add("client_secret", configer.AppConfig.OAuth2App.Keycloak.ClientSecret)
	params.Add("scope", "openid")
	params.Add("username", loginUser.EmailOrLogin)
	params.Add("password", loginUser.Password)

	jsonStr := []byte("")
	fmt.Println(jsonStr)
	fmt.Println(bytes.NewBuffer(jsonStr))
	fmt.Println(url + params.Encode())
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(params.Encode())))
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
	_ = json.Unmarshal(body, &token)

	var claims map[string]interface{}

	parsedToken, _ := jwt.ParseSigned(token.AccessToken)
	_ = parsedToken.UnsafeClaimsWithoutVerification(&claims)

	authUserId := StringToUid(claims["sub"].(string))
	userData, err := u.UserRepo.SelectUserByAuthId(authUserId, "keycloak")
	if err == nil {
		return userData.Id, nil
	}

	newLogin, err := u.UserRepo.SelectNewUniqLogin(claims["sub"].(string))
	if err != nil {
		return 0, err
	}

	userId, err := u.UserRepo.InsertAuthUser(&domain.AuthUserData{
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

func (u *UserController) LogInKeycloak(code string) (uint64, error) {
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
	userData, err := u.UserRepo.SelectUserByAuthId(authUserId, "keycloak")
	if err == nil {
		return userData.Id, nil
	}

	newLogin, err := u.UserRepo.SelectNewUniqLogin(claims["sub"].(string))
	if err != nil {
		return 0, err
	}

	userId, err := u.UserRepo.InsertAuthUser(&domain.AuthUserData{
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

func (u *UserController) GetSelfProfile(userId uint64) (*domain.UserProfile, error) {
	return u.UserRepo.SelectUserById(userId)
}

func (u *UserController) DeleteSelfProfile(userId uint64) error {
	err := u.UserRepo.DeleteSelfProfile(userId)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserController) GetUsers() ([]domain.UserProfile, error) {
	return u.UserRepo.GetUsers()
}
