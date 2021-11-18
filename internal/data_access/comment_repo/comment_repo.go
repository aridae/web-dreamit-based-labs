package comment_repo

import "github.com/aridae/web-dreamit-api-based-labs/internal/domain"

type Repository interface {
	CreateComment(domain.PostComment) (int64, error)
	DeleteComment(commId int64) error
	GetComment(commId int64) (*domain.Comment, error)
	GetNotifyComments(notifyId int64) ([]domain.Comment, error)
	GetAuthorComments(authorId uint64) ([]domain.Comment, error)
	DeleteNotifyComments(notifyId int64) error
}
