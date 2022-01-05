package controllers

import (
	"reflect"
	"testing"

	commentrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/comment_repo"
	domain "github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

func TestCommentController_CreateComment(t *testing.T) {
	type fields struct {
		CommentRepo commentrepo.Repository
	}
	type args struct {
		comment domain.PostComment
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommentController{
				CommentRepo: tt.fields.CommentRepo,
			}
			got, err := c.CreateComment(tt.args.comment)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentController.CreateComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CommentController.CreateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentController_DeleteComment(t *testing.T) {
	type fields struct {
		CommentRepo commentrepo.Repository
	}
	type args struct {
		commId int64
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
			c := &CommentController{
				CommentRepo: tt.fields.CommentRepo,
			}
			if err := c.DeleteComment(tt.args.commId); (err != nil) != tt.wantErr {
				t.Errorf("CommentController.DeleteComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommentController_GetComment(t *testing.T) {
	type fields struct {
		CommentRepo commentrepo.Repository
	}
	type args struct {
		commId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommentController{
				CommentRepo: tt.fields.CommentRepo,
			}
			got, err := c.GetComment(tt.args.commId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentController.GetComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentController.GetComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentController_GetNotifyComments(t *testing.T) {
	type fields struct {
		CommentRepo commentrepo.Repository
	}
	type args struct {
		notifyId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommentController{
				CommentRepo: tt.fields.CommentRepo,
			}
			got, err := c.GetNotifyComments(tt.args.notifyId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentController.GetNotifyComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentController.GetNotifyComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentController_GetAuthorComments(t *testing.T) {
	type fields struct {
		CommentRepo commentrepo.Repository
	}
	type args struct {
		authorId uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommentController{
				CommentRepo: tt.fields.CommentRepo,
			}
			got, err := c.GetAuthorComments(tt.args.authorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentController.GetAuthorComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentController.GetAuthorComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentController_DeleteNotifyComments(t *testing.T) {
	type fields struct {
		CommentRepo commentrepo.Repository
	}
	type args struct {
		notifyId int64
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
			c := &CommentController{
				CommentRepo: tt.fields.CommentRepo,
			}
			if err := c.DeleteNotifyComments(tt.args.notifyId); (err != nil) != tt.wantErr {
				t.Errorf("CommentController.DeleteNotifyComments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCommentController(t *testing.T) {
	type args struct {
		CommentRepo commentrepo.Repository
	}
	tests := []struct {
		name string
		args args
		want *CommentController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommentController(tt.args.CommentRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommentController() = %v, want %v", got, tt.want)
			}
		})
	}
}
