package apiserver

import (
	"net/http"
	"reflect"
	"testing"

	commentcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/comment_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
)

func TestCommentHandler_GetNotifyComments(t *testing.T) {
	type fields struct {
		CommentController *commentcont.CommentController
		SessionController *sessioncont.SessionController
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := CommentHandler{
				CommentController: tt.fields.CommentController,
				SessionController: tt.fields.SessionController,
			}
			handler.GetNotifyComments(tt.args.w, tt.args.r)
		})
	}
}

func TestCommentHandler_GetNotifyComment(t *testing.T) {
	type fields struct {
		CommentController *commentcont.CommentController
		SessionController *sessioncont.SessionController
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := CommentHandler{
				CommentController: tt.fields.CommentController,
				SessionController: tt.fields.SessionController,
			}
			handler.GetNotifyComment(tt.args.w, tt.args.r)
		})
	}
}

func TestCommentHandler_AddNotifyComment(t *testing.T) {
	type fields struct {
		CommentController *commentcont.CommentController
		SessionController *sessioncont.SessionController
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := CommentHandler{
				CommentController: tt.fields.CommentController,
				SessionController: tt.fields.SessionController,
			}
			handler.AddNotifyComment(tt.args.w, tt.args.r)
		})
	}
}

func TestCommentHandler_DeleteNotifyComment(t *testing.T) {
	type fields struct {
		CommentController *commentcont.CommentController
		SessionController *sessioncont.SessionController
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := CommentHandler{
				CommentController: tt.fields.CommentController,
				SessionController: tt.fields.SessionController,
			}
			handler.DeleteNotifyComment(tt.args.w, tt.args.r)
		})
	}
}

func TestNewCommentHandler(t *testing.T) {
	type args struct {
		CommentController *commentcont.CommentController
		SessionController *sessioncont.SessionController
	}
	tests := []struct {
		name string
		args args
		want *CommentHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommentHandler(tt.args.CommentController, tt.args.SessionController); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommentHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
