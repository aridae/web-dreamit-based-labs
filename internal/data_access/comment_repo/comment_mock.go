package comment_repo

import (
	"sync"

	"github.com/aridae/web-dreamit-api-based-labs/internal/api_server/apimodels"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

var (
	CurrCommId   int64            = 3
	MockComments []domain.Comment = []domain.Comment{
		{
			Id:       0,
			NotifyId: 1,
			AuthorId: 2,
			Message:  "test message 1",
		},
		{
			Id:       1,
			NotifyId: 2,
			AuthorId: 1,
			Message:  "test message 2",
		},
	}
	MockPostComment apimodels.PostComment = apimodels.PostComment{
		NotifyId: 2,
		Message:  "new test message",
	}
)

type RepositoryMock struct {
	// CreateCommentFunc mocks the CreateComment method.
	CreateCommentFunc func(postComment domain.PostComment) (int64, error)

	// DeleteCommentFunc mocks the DeleteComment method.
	DeleteCommentFunc func(commId int64) error

	// DeleteNotifyCommentsFunc mocks the DeleteNotifyComments method.
	DeleteNotifyCommentsFunc func(notifyId int64) error

	// GetAuthorCommentsFunc mocks the GetAuthorComments method.
	GetAuthorCommentsFunc func(authorId uint64) ([]domain.Comment, error)

	// GetCommentFunc mocks the GetComment method.
	GetCommentFunc func(commId int64) (*domain.Comment, error)

	// GetNotifyCommentsFunc mocks the GetNotifyComments method.
	GetNotifyCommentsFunc func(notifyId int64) ([]domain.Comment, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateComment holds details about calls to the CreateComment method.
		CreateComment []struct {
			// PostComment is the postComment argument value.
			PostComment domain.PostComment
		}
		// DeleteComment holds details about calls to the DeleteComment method.
		DeleteComment []struct {
			// CommId is the commId argument value.
			CommId int64
		}
		// DeleteNotifyComments holds details about calls to the DeleteNotifyComments method.
		DeleteNotifyComments []struct {
			// NotifyId is the notifyId argument value.
			NotifyId int64
		}
		// GetAuthorComments holds details about calls to the GetAuthorComments method.
		GetAuthorComments []struct {
			// AuthorId is the authorId argument value.
			AuthorId uint64
		}
		// GetComment holds details about calls to the GetComment method.
		GetComment []struct {
			// CommId is the commId argument value.
			CommId int64
		}
		// GetNotifyComments holds details about calls to the GetNotifyComments method.
		GetNotifyComments []struct {
			// NotifyId is the notifyId argument value.
			NotifyId int64
		}
	}
	lockCreateComment        sync.RWMutex
	lockDeleteComment        sync.RWMutex
	lockDeleteNotifyComments sync.RWMutex
	lockGetAuthorComments    sync.RWMutex
	lockGetComment           sync.RWMutex
	lockGetNotifyComments    sync.RWMutex
}

func createCommentFunc(postComment domain.PostComment) (int64, error) {
	return CurrCommId, nil
}

func getCommentFunc(commId int64) (*domain.Comment, error) {
	return &MockComments[commId], nil
}

func getCommentsFunc(notifyId int64) ([]domain.Comment, error) {
	return []domain.Comment{MockComments[1]}, nil
}

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{
		GetCommentFunc:        getCommentFunc,
		GetNotifyCommentsFunc: getCommentsFunc,
		CreateCommentFunc:     createCommentFunc,
	}
}

// CreateComment calls CreateCommentFunc.
func (mock *RepositoryMock) CreateComment(postComment domain.PostComment) (int64, error) {
	if mock.CreateCommentFunc == nil {
		panic("RepositoryMock.CreateCommentFunc: method is nil but Repository.CreateComment was just called")
	}
	callInfo := struct {
		PostComment domain.PostComment
	}{
		PostComment: postComment,
	}
	mock.lockCreateComment.Lock()
	mock.calls.CreateComment = append(mock.calls.CreateComment, callInfo)
	mock.lockCreateComment.Unlock()
	return mock.CreateCommentFunc(postComment)
}

// CreateCommentCalls gets all the calls that were made to CreateComment.
// Check the length with:
//     len(mockedRepository.CreateCommentCalls())
func (mock *RepositoryMock) CreateCommentCalls() []struct {
	PostComment domain.PostComment
} {
	var calls []struct {
		PostComment domain.PostComment
	}
	mock.lockCreateComment.RLock()
	calls = mock.calls.CreateComment
	mock.lockCreateComment.RUnlock()
	return calls
}

// DeleteComment calls DeleteCommentFunc.
func (mock *RepositoryMock) DeleteComment(commId int64) error {
	if mock.DeleteCommentFunc == nil {
		panic("RepositoryMock.DeleteCommentFunc: method is nil but Repository.DeleteComment was just called")
	}
	callInfo := struct {
		CommId int64
	}{
		CommId: commId,
	}
	mock.lockDeleteComment.Lock()
	mock.calls.DeleteComment = append(mock.calls.DeleteComment, callInfo)
	mock.lockDeleteComment.Unlock()
	return mock.DeleteCommentFunc(commId)
}

// DeleteCommentCalls gets all the calls that were made to DeleteComment.
// Check the length with:
//     len(mockedRepository.DeleteCommentCalls())
func (mock *RepositoryMock) DeleteCommentCalls() []struct {
	CommId int64
} {
	var calls []struct {
		CommId int64
	}
	mock.lockDeleteComment.RLock()
	calls = mock.calls.DeleteComment
	mock.lockDeleteComment.RUnlock()
	return calls
}

// DeleteNotifyComments calls DeleteNotifyCommentsFunc.
func (mock *RepositoryMock) DeleteNotifyComments(notifyId int64) error {
	if mock.DeleteNotifyCommentsFunc == nil {
		panic("RepositoryMock.DeleteNotifyCommentsFunc: method is nil but Repository.DeleteNotifyComments was just called")
	}
	callInfo := struct {
		NotifyId int64
	}{
		NotifyId: notifyId,
	}
	mock.lockDeleteNotifyComments.Lock()
	mock.calls.DeleteNotifyComments = append(mock.calls.DeleteNotifyComments, callInfo)
	mock.lockDeleteNotifyComments.Unlock()
	return mock.DeleteNotifyCommentsFunc(notifyId)
}

// DeleteNotifyCommentsCalls gets all the calls that were made to DeleteNotifyComments.
// Check the length with:
//     len(mockedRepository.DeleteNotifyCommentsCalls())
func (mock *RepositoryMock) DeleteNotifyCommentsCalls() []struct {
	NotifyId int64
} {
	var calls []struct {
		NotifyId int64
	}
	mock.lockDeleteNotifyComments.RLock()
	calls = mock.calls.DeleteNotifyComments
	mock.lockDeleteNotifyComments.RUnlock()
	return calls
}

// GetAuthorComments calls GetAuthorCommentsFunc.
func (mock *RepositoryMock) GetAuthorComments(authorId uint64) ([]domain.Comment, error) {
	if mock.GetAuthorCommentsFunc == nil {
		panic("RepositoryMock.GetAuthorCommentsFunc: method is nil but Repository.GetAuthorComments was just called")
	}
	callInfo := struct {
		AuthorId uint64
	}{
		AuthorId: authorId,
	}
	mock.lockGetAuthorComments.Lock()
	mock.calls.GetAuthorComments = append(mock.calls.GetAuthorComments, callInfo)
	mock.lockGetAuthorComments.Unlock()
	return mock.GetAuthorCommentsFunc(authorId)
}

// GetAuthorCommentsCalls gets all the calls that were made to GetAuthorComments.
// Check the length with:
//     len(mockedRepository.GetAuthorCommentsCalls())
func (mock *RepositoryMock) GetAuthorCommentsCalls() []struct {
	AuthorId uint64
} {
	var calls []struct {
		AuthorId uint64
	}
	mock.lockGetAuthorComments.RLock()
	calls = mock.calls.GetAuthorComments
	mock.lockGetAuthorComments.RUnlock()
	return calls
}

// GetComment calls GetCommentFunc.
func (mock *RepositoryMock) GetComment(commId int64) (*domain.Comment, error) {
	if mock.GetCommentFunc == nil {
		panic("RepositoryMock.GetCommentFunc: method is nil but Repository.GetComment was just called")
	}
	callInfo := struct {
		CommId int64
	}{
		CommId: commId,
	}
	mock.lockGetComment.Lock()
	mock.calls.GetComment = append(mock.calls.GetComment, callInfo)
	mock.lockGetComment.Unlock()
	return mock.GetCommentFunc(commId)
}

// GetCommentCalls gets all the calls that were made to GetComment.
// Check the length with:
//     len(mockedRepository.GetCommentCalls())
func (mock *RepositoryMock) GetCommentCalls() []struct {
	CommId int64
} {
	var calls []struct {
		CommId int64
	}
	mock.lockGetComment.RLock()
	calls = mock.calls.GetComment
	mock.lockGetComment.RUnlock()
	return calls
}

// GetNotifyComments calls GetNotifyCommentsFunc.
func (mock *RepositoryMock) GetNotifyComments(notifyId int64) ([]domain.Comment, error) {
	if mock.GetNotifyCommentsFunc == nil {
		panic("RepositoryMock.GetNotifyCommentsFunc: method is nil but Repository.GetNotifyComments was just called")
	}
	callInfo := struct {
		NotifyId int64
	}{
		NotifyId: notifyId,
	}
	mock.lockGetNotifyComments.Lock()
	mock.calls.GetNotifyComments = append(mock.calls.GetNotifyComments, callInfo)
	mock.lockGetNotifyComments.Unlock()
	return mock.GetNotifyCommentsFunc(notifyId)
}

// GetNotifyCommentsCalls gets all the calls that were made to GetNotifyComments.
// Check the length with:
//     len(mockedRepository.GetNotifyCommentsCalls())
func (mock *RepositoryMock) GetNotifyCommentsCalls() []struct {
	NotifyId int64
} {
	var calls []struct {
		NotifyId int64
	}
	mock.lockGetNotifyComments.RLock()
	calls = mock.calls.GetNotifyComments
	mock.lockGetNotifyComments.RUnlock()
	return calls
}
