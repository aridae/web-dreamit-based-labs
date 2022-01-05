package controllers

import (
	"reflect"
	"testing"

	userrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/user_repo"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

func TestNewUserController(t *testing.T) {
	type args struct {
		UserRepo userrepo.Repository
	}
	tests := []struct {
		name string
		args args
		want *UserController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserController(tt.args.UserRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToUid(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToUid(tt.args.str); got != tt.want {
				t.Errorf("StringToUid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_SignUp(t *testing.T) {
	type fields struct {
		UserRepo userrepo.Repository
	}
	type args struct {
		signupUser *domain.SignupUserData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				UserRepo: tt.fields.UserRepo,
			}
			got, err := u.SignUp(tt.args.signupUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserController.SignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_LogIn(t *testing.T) {
	type fields struct {
		UserRepo userrepo.Repository
	}
	type args struct {
		loginUser *domain.LoginUserData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				UserRepo: tt.fields.UserRepo,
			}
			got, err := u.LogIn(tt.args.loginUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.LogIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserController.LogIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_LogInKeycloak(t *testing.T) {
	type fields struct {
		UserRepo userrepo.Repository
	}
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				UserRepo: tt.fields.UserRepo,
			}
			got, err := u.LogInKeycloak(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.LogInKeycloak() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserController.LogInKeycloak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_GetSelfProfile(t *testing.T) {
	type fields struct {
		UserRepo userrepo.Repository
	}
	type args struct {
		userId uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.UserProfile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				UserRepo: tt.fields.UserRepo,
			}
			got, err := u.GetSelfProfile(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.GetSelfProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserController.GetSelfProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_DeleteSelfProfile(t *testing.T) {
	type fields struct {
		UserRepo userrepo.Repository
	}
	type args struct {
		userId uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				UserRepo: tt.fields.UserRepo,
			}
			if err := u.DeleteSelfProfile(tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("UserController.DeleteSelfProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserController_GetUsers(t *testing.T) {
	type fields struct {
		UserRepo userrepo.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.UserProfile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserController{
				UserRepo: tt.fields.UserRepo,
			}
			got, err := u.GetUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserController.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
