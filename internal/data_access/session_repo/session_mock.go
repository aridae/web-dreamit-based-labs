// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sessionrepo

import (
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked Repository
// 		mockedRepository := &RepositoryMock{
// 			DeleteAccessTokenFunc: func(Uuid string) error {
// 				panic("mock out the DeleteAccessToken method")
// 			},
// 			DeleteRefreshTokenFunc: func(Uuid string) error {
// 				panic("mock out the DeleteRefreshToken method")
// 			},
// 			InsertTokenFunc: func(userId uint64, token domain.TokenDetails) error {
// 				panic("mock out the InsertToken method")
// 			},
// 			SelectUserIdByAccessTokenFunc: func(Uuid string) (uint64, error) {
// 				panic("mock out the SelectUserIdByAccessToken method")
// 			},
// 			SelectUserIdByRefreshTokenFunc: func(Uuid string) (uint64, error) {
// 				panic("mock out the SelectUserIdByRefreshToken method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// DeleteAccessTokenFunc mocks the DeleteAccessToken method.
	DeleteAccessTokenFunc func(Uuid string) error

	// DeleteRefreshTokenFunc mocks the DeleteRefreshToken method.
	DeleteRefreshTokenFunc func(Uuid string) error

	// InsertTokenFunc mocks the InsertToken method.
	InsertTokenFunc func(userId uint64, token domain.TokenDetails) error

	// SelectUserIdByAccessTokenFunc mocks the SelectUserIdByAccessToken method.
	SelectUserIdByAccessTokenFunc func(Uuid string) (uint64, error)

	// SelectUserIdByRefreshTokenFunc mocks the SelectUserIdByRefreshToken method.
	SelectUserIdByRefreshTokenFunc func(Uuid string) (uint64, error)

	// calls tracks calls to the methods.
	calls struct {
		// DeleteAccessToken holds details about calls to the DeleteAccessToken method.
		DeleteAccessToken []struct {
			// UUID is the Uuid argument value.
			UUID string
		}
		// DeleteRefreshToken holds details about calls to the DeleteRefreshToken method.
		DeleteRefreshToken []struct {
			// UUID is the Uuid argument value.
			UUID string
		}
		// InsertToken holds details about calls to the InsertToken method.
		InsertToken []struct {
			// UserId is the userId argument value.
			UserId uint64
			// Token is the token argument value.
			Token domain.TokenDetails
		}
		// SelectUserIdByAccessToken holds details about calls to the SelectUserIdByAccessToken method.
		SelectUserIdByAccessToken []struct {
			// UUID is the Uuid argument value.
			UUID string
		}
		// SelectUserIdByRefreshToken holds details about calls to the SelectUserIdByRefreshToken method.
		SelectUserIdByRefreshToken []struct {
			// UUID is the Uuid argument value.
			UUID string
		}
	}
	lockDeleteAccessToken          sync.RWMutex
	lockDeleteRefreshToken         sync.RWMutex
	lockInsertToken                sync.RWMutex
	lockSelectUserIdByAccessToken  sync.RWMutex
	lockSelectUserIdByRefreshToken sync.RWMutex
}

// DeleteAccessToken calls DeleteAccessTokenFunc.
func (mock *RepositoryMock) DeleteAccessToken(Uuid string) error {
	if mock.DeleteAccessTokenFunc == nil {
		panic("RepositoryMock.DeleteAccessTokenFunc: method is nil but Repository.DeleteAccessToken was just called")
	}
	callInfo := struct {
		UUID string
	}{
		UUID: Uuid,
	}
	mock.lockDeleteAccessToken.Lock()
	mock.calls.DeleteAccessToken = append(mock.calls.DeleteAccessToken, callInfo)
	mock.lockDeleteAccessToken.Unlock()
	return mock.DeleteAccessTokenFunc(Uuid)
}

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{
	}
}

// DeleteAccessTokenCalls gets all the calls that were made to DeleteAccessToken.
// Check the length with:
//     len(mockedRepository.DeleteAccessTokenCalls())
func (mock *RepositoryMock) DeleteAccessTokenCalls() []struct {
	UUID string
} {
	var calls []struct {
		UUID string
	}
	mock.lockDeleteAccessToken.RLock()
	calls = mock.calls.DeleteAccessToken
	mock.lockDeleteAccessToken.RUnlock()
	return calls
}

// DeleteRefreshToken calls DeleteRefreshTokenFunc.
func (mock *RepositoryMock) DeleteRefreshToken(Uuid string) error {
	if mock.DeleteRefreshTokenFunc == nil {
		panic("RepositoryMock.DeleteRefreshTokenFunc: method is nil but Repository.DeleteRefreshToken was just called")
	}
	callInfo := struct {
		UUID string
	}{
		UUID: Uuid,
	}
	mock.lockDeleteRefreshToken.Lock()
	mock.calls.DeleteRefreshToken = append(mock.calls.DeleteRefreshToken, callInfo)
	mock.lockDeleteRefreshToken.Unlock()
	return mock.DeleteRefreshTokenFunc(Uuid)
}

// DeleteRefreshTokenCalls gets all the calls that were made to DeleteRefreshToken.
// Check the length with:
//     len(mockedRepository.DeleteRefreshTokenCalls())
func (mock *RepositoryMock) DeleteRefreshTokenCalls() []struct {
	UUID string
} {
	var calls []struct {
		UUID string
	}
	mock.lockDeleteRefreshToken.RLock()
	calls = mock.calls.DeleteRefreshToken
	mock.lockDeleteRefreshToken.RUnlock()
	return calls
}

// InsertToken calls InsertTokenFunc.
func (mock *RepositoryMock) InsertToken(userId uint64, token domain.TokenDetails) error {
	if mock.InsertTokenFunc == nil {
		panic("RepositoryMock.InsertTokenFunc: method is nil but Repository.InsertToken was just called")
	}
	callInfo := struct {
		UserId uint64
		Token  domain.TokenDetails
	}{
		UserId: userId,
		Token:  token,
	}
	mock.lockInsertToken.Lock()
	mock.calls.InsertToken = append(mock.calls.InsertToken, callInfo)
	mock.lockInsertToken.Unlock()
	return mock.InsertTokenFunc(userId, token)
}

// InsertTokenCalls gets all the calls that were made to InsertToken.
// Check the length with:
//     len(mockedRepository.InsertTokenCalls())
func (mock *RepositoryMock) InsertTokenCalls() []struct {
	UserId uint64
	Token  domain.TokenDetails
} {
	var calls []struct {
		UserId uint64
		Token  domain.TokenDetails
	}
	mock.lockInsertToken.RLock()
	calls = mock.calls.InsertToken
	mock.lockInsertToken.RUnlock()
	return calls
}

// SelectUserIdByAccessToken calls SelectUserIdByAccessTokenFunc.
func (mock *RepositoryMock) SelectUserIdByAccessToken(Uuid string) (uint64, error) {
	if mock.SelectUserIdByAccessTokenFunc == nil {
		panic("RepositoryMock.SelectUserIdByAccessTokenFunc: method is nil but Repository.SelectUserIdByAccessToken was just called")
	}
	callInfo := struct {
		UUID string
	}{
		UUID: Uuid,
	}
	mock.lockSelectUserIdByAccessToken.Lock()
	mock.calls.SelectUserIdByAccessToken = append(mock.calls.SelectUserIdByAccessToken, callInfo)
	mock.lockSelectUserIdByAccessToken.Unlock()
	return mock.SelectUserIdByAccessTokenFunc(Uuid)
}

// SelectUserIdByAccessTokenCalls gets all the calls that were made to SelectUserIdByAccessToken.
// Check the length with:
//     len(mockedRepository.SelectUserIdByAccessTokenCalls())
func (mock *RepositoryMock) SelectUserIdByAccessTokenCalls() []struct {
	UUID string
} {
	var calls []struct {
		UUID string
	}
	mock.lockSelectUserIdByAccessToken.RLock()
	calls = mock.calls.SelectUserIdByAccessToken
	mock.lockSelectUserIdByAccessToken.RUnlock()
	return calls
}

// SelectUserIdByRefreshToken calls SelectUserIdByRefreshTokenFunc.
func (mock *RepositoryMock) SelectUserIdByRefreshToken(Uuid string) (uint64, error) {
	if mock.SelectUserIdByRefreshTokenFunc == nil {
		panic("RepositoryMock.SelectUserIdByRefreshTokenFunc: method is nil but Repository.SelectUserIdByRefreshToken was just called")
	}
	callInfo := struct {
		UUID string
	}{
		UUID: Uuid,
	}
	mock.lockSelectUserIdByRefreshToken.Lock()
	mock.calls.SelectUserIdByRefreshToken = append(mock.calls.SelectUserIdByRefreshToken, callInfo)
	mock.lockSelectUserIdByRefreshToken.Unlock()
	return mock.SelectUserIdByRefreshTokenFunc(Uuid)
}

// SelectUserIdByRefreshTokenCalls gets all the calls that were made to SelectUserIdByRefreshToken.
// Check the length with:
//     len(mockedRepository.SelectUserIdByRefreshTokenCalls())
func (mock *RepositoryMock) SelectUserIdByRefreshTokenCalls() []struct {
	UUID string
} {
	var calls []struct {
		UUID string
	}
	mock.lockSelectUserIdByRefreshToken.RLock()
	calls = mock.calls.SelectUserIdByRefreshToken
	mock.lockSelectUserIdByRefreshToken.RUnlock()
	return calls
}
