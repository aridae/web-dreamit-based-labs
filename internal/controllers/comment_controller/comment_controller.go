package controllers

import (
	commentrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/comment_repo"
	domain "github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

type CommentController struct {
	CommentRepo commentrepo.Repository
}

func (c *CommentController) CreateComment(comment domain.PostComment) (int64, error) {
	return c.CommentRepo.CreateComment(comment)
}

func (c *CommentController) DeleteComment(commId int64) error {
	return c.CommentRepo.DeleteComment(commId)
}

func (c *CommentController) GetComment(commId int64) (*domain.Comment, error) {
	return c.CommentRepo.GetComment(commId)
}

func (c *CommentController) GetNotifyComments(notifyId int64) ([]domain.Comment, error) {
	return c.CommentRepo.GetNotifyComments(notifyId)
}

func (c *CommentController) GetAuthorComments(authorId uint64) ([]domain.Comment, error) {
	return c.CommentRepo.GetAuthorComments(authorId)
}

func (c *CommentController) DeleteNotifyComments(notifyId int64) error {
	return c.CommentRepo.DeleteNotifyComments(notifyId)
}

func NewCommentController(CommentRepo commentrepo.Repository) *CommentController {
	return &CommentController{
		CommentRepo: CommentRepo,
	}
}
