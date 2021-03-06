// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package userrepo

import (
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

var (
	CurrUserId = 2
	MockUsers []domain.UserData = []domain.UserData{
		{
			Id: 0,
			Email: "some@email.ru",
			Login: "MockUser_01",
			Password: []byte("superPass24537"),
		},
		{
			Id: 1,
			Email: "some222@email.ru",
			Login: "MockUser_01",
			Password: []byte("superDDHDH24537"),
		},
	}
)

type RepositoryMock struct {
	// DeleteSelfProfileFunc mocks the DeleteSelfProfile method.
	DeleteSelfProfileFunc func(userId uint64) error

	// GetUsersFunc mocks the GetUsers method.
	GetUsersFunc func() ([]domain.UserProfile, error)

	// InsertAuthUserFunc mocks the InsertAuthUser method.
	InsertAuthUserFunc func(user *domain.AuthUserData) (uint64, error)

	// InsertUserFunc mocks the InsertUser method.
	InsertUserFunc func(user *domain.UserData) (uint64, error)

	// SelectNewUniqLoginFunc mocks the SelectNewUniqLogin method.
	SelectNewUniqLoginFunc func(login string) (string, error)

	// SelectUserByAuthIdFunc mocks the SelectUserByAuthId method.
	SelectUserByAuthIdFunc func(authId uint64, authService string) (*domain.UserData, error)

	// SelectUserByEmailOrLoginFunc mocks the SelectUserByEmailOrLogin method.
	SelectUserByEmailOrLoginFunc func(emailOrLogin string) (*domain.UserData, error)

	// SelectUserByIdFunc mocks the SelectUserById method.
	SelectUserByIdFunc func(userId uint64) (*domain.UserProfile, error)

	// calls tracks calls to the methods.
	calls struct {
		// DeleteSelfProfile holds details about calls to the DeleteSelfProfile method.
		DeleteSelfProfile []struct {
			// UserId is the userId argument value.
			UserId uint64
		}
		// GetUsers holds details about calls to the GetUsers method.
		GetUsers []struct {
		}
		// InsertAuthUser holds details about calls to the InsertAuthUser method.
		InsertAuthUser []struct {
			// User is the user argument value.
			User *domain.AuthUserData
		}
		// InsertUser holds details about calls to the InsertUser method.
		InsertUser []struct {
			// User is the user argument value.
			User *domain.UserData
		}
		// SelectNewUniqLogin holds details about calls to the SelectNewUniqLogin method.
		SelectNewUniqLogin []struct {
			// Login is the login argument value.
			Login string
		}
		// SelectUserByAuthId holds details about calls to the SelectUserByAuthId method.
		SelectUserByAuthId []struct {
			// AuthId is the authId argument value.
			AuthId uint64
			// AuthService is the authService argument value.
			AuthService string
		}
		// SelectUserByEmailOrLogin holds details about calls to the SelectUserByEmailOrLogin method.
		SelectUserByEmailOrLogin []struct {
			// EmailOrLogin is the emailOrLogin argument value.
			EmailOrLogin string
		}
		// SelectUserById holds details about calls to the SelectUserById method.
		SelectUserById []struct {
			// UserId is the userId argument value.
			UserId uint64
		}
	}
	lockDeleteSelfProfile        sync.RWMutex
	lockGetUsers                 sync.RWMutex
	lockInsertAuthUser           sync.RWMutex
	lockInsertUser               sync.RWMutex
	lockSelectNewUniqLogin       sync.RWMutex
	lockSelectUserByAuthId       sync.RWMutex
	lockSelectUserByEmailOrLogin sync.RWMutex
	lockSelectUserById           sync.RWMutex
}

func insertUser(user *domain.UserData) (uint64, error) {
	CurrUserId++
	return uint64(CurrUserId), nil
}

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{
		InsertUserFunc: insertUser,
	}
}

// DeleteSelfProfile calls DeleteSelfProfileFunc.
func (mock *RepositoryMock) DeleteSelfProfile(userId uint64) error {
	if mock.DeleteSelfProfileFunc == nil {
		panic("RepositoryMock.DeleteSelfProfileFunc: method is nil but Repository.DeleteSelfProfile was just called")
	}
	callInfo := struct {
		UserId uint64
	}{
		UserId: userId,
	}
	mock.lockDeleteSelfProfile.Lock()
	mock.calls.DeleteSelfProfile = append(mock.calls.DeleteSelfProfile, callInfo)
	mock.lockDeleteSelfProfile.Unlock()
	return mock.DeleteSelfProfileFunc(userId)
}

// DeleteSelfProfileCalls gets all the calls that were made to DeleteSelfProfile.
// Check the length with:
//     len(mockedRepository.DeleteSelfProfileCalls())
func (mock *RepositoryMock) DeleteSelfProfileCalls() []struct {
	UserId uint64
} {
	var calls []struct {
		UserId uint64
	}
	mock.lockDeleteSelfProfile.RLock()
	calls = mock.calls.DeleteSelfProfile
	mock.lockDeleteSelfProfile.RUnlock()
	return calls
}

// GetUsers calls GetUsersFunc.
func (mock *RepositoryMock) GetUsers() ([]domain.UserProfile, error) {
	if mock.GetUsersFunc == nil {
		panic("RepositoryMock.GetUsersFunc: method is nil but Repository.GetUsers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetUsers.Lock()
	mock.calls.GetUsers = append(mock.calls.GetUsers, callInfo)
	mock.lockGetUsers.Unlock()
	return mock.GetUsersFunc()
}

// GetUsersCalls gets all the calls that were made to GetUsers.
// Check the length with:
//     len(mockedRepository.GetUsersCalls())
func (mock *RepositoryMock) GetUsersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetUsers.RLock()
	calls = mock.calls.GetUsers
	mock.lockGetUsers.RUnlock()
	return calls
}

// InsertAuthUser calls InsertAuthUserFunc.
func (mock *RepositoryMock) InsertAuthUser(user *domain.AuthUserData) (uint64, error) {
	if mock.InsertAuthUserFunc == nil {
		panic("RepositoryMock.InsertAuthUserFunc: method is nil but Repository.InsertAuthUser was just called")
	}
	callInfo := struct {
		User *domain.AuthUserData
	}{
		User: user,
	}
	mock.lockInsertAuthUser.Lock()
	mock.calls.InsertAuthUser = append(mock.calls.InsertAuthUser, callInfo)
	mock.lockInsertAuthUser.Unlock()
	return mock.InsertAuthUserFunc(user)
}

// InsertAuthUserCalls gets all the calls that were made to InsertAuthUser.
// Check the length with:
//     len(mockedRepository.InsertAuthUserCalls())
func (mock *RepositoryMock) InsertAuthUserCalls() []struct {
	User *domain.AuthUserData
} {
	var calls []struct {
		User *domain.AuthUserData
	}
	mock.lockInsertAuthUser.RLock()
	calls = mock.calls.InsertAuthUser
	mock.lockInsertAuthUser.RUnlock()
	return calls
}

// InsertUser calls InsertUserFunc.
func (mock *RepositoryMock) InsertUser(user *domain.UserData) (uint64, error) {
	if mock.InsertUserFunc == nil {
		panic("RepositoryMock.InsertUserFunc: method is nil but Repository.InsertUser was just called")
	}
	callInfo := struct {
		User *domain.UserData
	}{
		User: user,
	}
	mock.lockInsertUser.Lock()
	mock.calls.InsertUser = append(mock.calls.InsertUser, callInfo)
	mock.lockInsertUser.Unlock()
	return mock.InsertUserFunc(user)
}

// InsertUserCalls gets all the calls that were made to InsertUser.
// Check the length with:
//     len(mockedRepository.InsertUserCalls())
func (mock *RepositoryMock) InsertUserCalls() []struct {
	User *domain.UserData
} {
	var calls []struct {
		User *domain.UserData
	}
	mock.lockInsertUser.RLock()
	calls = mock.calls.InsertUser
	mock.lockInsertUser.RUnlock()
	return calls
}

// SelectNewUniqLogin calls SelectNewUniqLoginFunc.
func (mock *RepositoryMock) SelectNewUniqLogin(login string) (string, error) {
	if mock.SelectNewUniqLoginFunc == nil {
		panic("RepositoryMock.SelectNewUniqLoginFunc: method is nil but Repository.SelectNewUniqLogin was just called")
	}
	callInfo := struct {
		Login string
	}{
		Login: login,
	}
	mock.lockSelectNewUniqLogin.Lock()
	mock.calls.SelectNewUniqLogin = append(mock.calls.SelectNewUniqLogin, callInfo)
	mock.lockSelectNewUniqLogin.Unlock()
	return mock.SelectNewUniqLoginFunc(login)
}

// SelectNewUniqLoginCalls gets all the calls that were made to SelectNewUniqLogin.
// Check the length with:
//     len(mockedRepository.SelectNewUniqLoginCalls())
func (mock *RepositoryMock) SelectNewUniqLoginCalls() []struct {
	Login string
} {
	var calls []struct {
		Login string
	}
	mock.lockSelectNewUniqLogin.RLock()
	calls = mock.calls.SelectNewUniqLogin
	mock.lockSelectNewUniqLogin.RUnlock()
	return calls
}

// SelectUserByAuthId calls SelectUserByAuthIdFunc.
func (mock *RepositoryMock) SelectUserByAuthId(authId uint64, authService string) (*domain.UserData, error) {
	if mock.SelectUserByAuthIdFunc == nil {
		panic("RepositoryMock.SelectUserByAuthIdFunc: method is nil but Repository.SelectUserByAuthId was just called")
	}
	callInfo := struct {
		AuthId      uint64
		AuthService string
	}{
		AuthId:      authId,
		AuthService: authService,
	}
	mock.lockSelectUserByAuthId.Lock()
	mock.calls.SelectUserByAuthId = append(mock.calls.SelectUserByAuthId, callInfo)
	mock.lockSelectUserByAuthId.Unlock()
	return mock.SelectUserByAuthIdFunc(authId, authService)
}

// SelectUserByAuthIdCalls gets all the calls that were made to SelectUserByAuthId.
// Check the length with:
//     len(mockedRepository.SelectUserByAuthIdCalls())
func (mock *RepositoryMock) SelectUserByAuthIdCalls() []struct {
	AuthId      uint64
	AuthService string
} {
	var calls []struct {
		AuthId      uint64
		AuthService string
	}
	mock.lockSelectUserByAuthId.RLock()
	calls = mock.calls.SelectUserByAuthId
	mock.lockSelectUserByAuthId.RUnlock()
	return calls
}

// SelectUserByEmailOrLogin calls SelectUserByEmailOrLoginFunc.
func (mock *RepositoryMock) SelectUserByEmailOrLogin(emailOrLogin string) (*domain.UserData, error) {
	if mock.SelectUserByEmailOrLoginFunc == nil {
		panic("RepositoryMock.SelectUserByEmailOrLoginFunc: method is nil but Repository.SelectUserByEmailOrLogin was just called")
	}
	callInfo := struct {
		EmailOrLogin string
	}{
		EmailOrLogin: emailOrLogin,
	}
	mock.lockSelectUserByEmailOrLogin.Lock()
	mock.calls.SelectUserByEmailOrLogin = append(mock.calls.SelectUserByEmailOrLogin, callInfo)
	mock.lockSelectUserByEmailOrLogin.Unlock()
	return mock.SelectUserByEmailOrLoginFunc(emailOrLogin)
}

// SelectUserByEmailOrLoginCalls gets all the calls that were made to SelectUserByEmailOrLogin.
// Check the length with:
//     len(mockedRepository.SelectUserByEmailOrLoginCalls())
func (mock *RepositoryMock) SelectUserByEmailOrLoginCalls() []struct {
	EmailOrLogin string
} {
	var calls []struct {
		EmailOrLogin string
	}
	mock.lockSelectUserByEmailOrLogin.RLock()
	calls = mock.calls.SelectUserByEmailOrLogin
	mock.lockSelectUserByEmailOrLogin.RUnlock()
	return calls
}

// SelectUserById calls SelectUserByIdFunc.
func (mock *RepositoryMock) SelectUserById(userId uint64) (*domain.UserProfile, error) {
	if mock.SelectUserByIdFunc == nil {
		panic("RepositoryMock.SelectUserByIdFunc: method is nil but Repository.SelectUserById was just called")
	}
	callInfo := struct {
		UserId uint64
	}{
		UserId: userId,
	}
	mock.lockSelectUserById.Lock()
	mock.calls.SelectUserById = append(mock.calls.SelectUserById, callInfo)
	mock.lockSelectUserById.Unlock()
	return mock.SelectUserByIdFunc(userId)
}

// SelectUserByIdCalls gets all the calls that were made to SelectUserById.
// Check the length with:
//     len(mockedRepository.SelectUserByIdCalls())
func (mock *RepositoryMock) SelectUserByIdCalls() []struct {
	UserId uint64
} {
	var calls []struct {
		UserId uint64
	}
	mock.lockSelectUserById.RLock()
	calls = mock.calls.SelectUserById
	mock.lockSelectUserById.RUnlock()
	return calls
}
