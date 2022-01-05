package apiserver

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/aridae/web-dreamit-api-based-labs/internal/api_server/apimodels"
	commentcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/comment_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	"github.com/aridae/web-dreamit-api-based-labs/internal/data_access/comment_repo"
	sessionrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/session_repo"
	testinghelp "github.com/aridae/web-dreamit-api-based-labs/internal/utils/testing"
	"github.com/stretchr/testify/assert"
)

var (
	ROUTE = "/comments"
)

func setupTest() *CommentHandler {
	h := &CommentHandler{
		CommentController: commentcont.NewCommentController(comment_repo.NewRepositoryMock()),
		SessionController: sessioncont.NewSessionController(&sessionrepo.RepositoryMock{}),
	}
	return h
}

func TestCommentHandler_GetNotifyComments(t *testing.T) {
	type testInput struct {
		notifyId string
	}
	type testWant struct {
		code int
		body string
	}
	tests := []struct {
		input testInput
		want  testWant
	}{
		{
			input: testInput{
				notifyId: "killme",
			},
			want: testWant{
				code: http.StatusBadRequest,
				body: fmt.Sprintf("{\"message\":\"%s\"}",
					fmt.Sprintf(FAILURE_COMMENTS, fmt.Errorf("invalid query paramete notifyId"))),
			},
		},
		{
			input: testInput{
				notifyId: "2",
			},
			want: testWant{
				code: http.StatusOK,
				body: fmt.Sprintf(
					"[{\"Id\":%d,\"NotifyId\":%d,\"AuthorId\":%d,\"Message\":\"%s\"}]",
					comment_repo.MockComments[1].Id, comment_repo.MockComments[1].NotifyId,
					comment_repo.MockComments[1].AuthorId, comment_repo.MockComments[1].Message,
				),
			},
		},
	}
	for _, test := range tests {
		reqBuilder := testinghelp.NewRequestBuilder()
		r, w := reqBuilder.WithParameter("notifyId", test.input.notifyId).WithMethod(http.MethodGet).WithRoute(ROUTE).Build()
		h := setupTest()
		h.GetNotifyComments(w, r)
		assert.Equal(t, test.want.code, w.Code)
		assert.Equal(t, test.want.body, w.Body.String())
	}
}

func TestCommentHandler_GetNotifyComment(t *testing.T) {
	type testInput struct {
		id string
	}
	type testWant struct {
		code int
		body string
	}

	tests := []struct {
		input testInput
		want  testWant
	}{
		{
			input: testInput{
				id: "",
			},
			want: testWant{
				code: http.StatusBadRequest,
				body: fmt.Sprintf("{\"message\":\"%s\"}",
					fmt.Sprintf(FAILURE_GET_COMMENT, fmt.Errorf("invalid path variable id"))),
			},
		},
		{
			input: testInput{
				id: "killme",
			},
			want: testWant{
				code: http.StatusBadRequest,
				body: fmt.Sprintf("{\"message\":\"%s\"}",
					fmt.Sprintf(FAILURE_GET_COMMENT, fmt.Errorf("invalid path variable id"))),
			},
		},
		{
			input: testInput{
				id: "1",
			},
			want: testWant{
				code: http.StatusOK,
				body: fmt.Sprintf(
					"{\"Id\":%d,\"NotifyId\":%d,\"AuthorId\":%d,\"Message\":\"%s\"}",
					comment_repo.MockComments[1].Id, comment_repo.MockComments[1].NotifyId,
					comment_repo.MockComments[1].AuthorId, comment_repo.MockComments[1].Message,
				),
			},
		},
	}
	for _, test := range tests {
		reqBuilder := testinghelp.NewRequestBuilder()
		r, w := reqBuilder.WithRouteVar("id", test.input.id).WithMethod(http.MethodGet).WithRoute(ROUTE).Build()
		h := setupTest()
		h.GetNotifyComment(w, r)
		assert.Equal(t, test.want.code, w.Code)
		assert.Equal(t, test.want.body, w.Body.String())
	}
}

func TestCommentHandler_AddNotifyComment(t *testing.T) {
	type testInput struct {
		body apimodels.PostComment
	}
	type testWant struct {
		code int
		body string
	}
	tests := []struct {
		input testInput
		want  testWant
	}{
		{
			input: testInput{
				body: comment_repo.MockPostComment,
			},
			want: testWant{
				code: http.StatusCreated,
				body: fmt.Sprintf("{\"id\":%d}", comment_repo.CurrCommId),
			},
		},
	}
	for _, test := range tests {
		reqBuilder := testinghelp.NewRequestBuilder()
		r, w := reqBuilder.WithMethod(http.MethodPost).WithRoute(ROUTE).WithBody(test.input.body).Build()
		h := setupTest()
		h.AddNotifyComment(w, r)
		assert.Equal(t, test.want.code, w.Code)
		assert.Equal(t, test.want.body, w.Body.String())
	}
}
