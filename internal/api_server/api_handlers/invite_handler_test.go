package apiserver

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	invitecont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/invite_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	inviterepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/invite_repo"
	sessionrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/session_repo"
	testinghelp "github.com/aridae/web-dreamit-api-based-labs/internal/utils/testing"
	"github.com/stretchr/testify/assert"
)

var (
	COMMENT_ROUTE = "/comments"
	INVITE_ROUTE  = "/invites"
)

func setupInviteTest() *InviteHandler {
	h := &InviteHandler{
		InviteController:  invitecont.NewInviteController(inviterepo.NewRepositoryMock()),
		SessionController: sessioncont.NewSessionController(&sessionrepo.RepositoryMock{}),
	}
	return h
}

func TestInviteHandler_GetInvites(t *testing.T) {
	type testInput struct {
		inviteId string
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
				inviteId: "killme",
			},
			want: testWant{
				code: http.StatusBadRequest,
				body: fmt.Sprintf("{\"message\":\"%s\"}",
					fmt.Sprintf(FAILURE_GET_INVITE, "invalid parameters")),
			},
		},
		{
			input: testInput{
				inviteId: "0",
			},
			want: testWant{
				code: http.StatusOK,
				body: fmt.Sprintf(
					"{\"Id\":%d,\"EventId\":%d,\"ReceiverId\":%d,\"StatusId\":%d}",
					inviterepo.MockInvites[0].Id, inviterepo.MockInvites[0].EventId, inviterepo.MockInvites[0].ReceiverId, inviterepo.MockInvites[0].StatusId),
			},
		},
	}
	for _, test := range tests {
		reqBuilder := testinghelp.NewRequestBuilder()

		reqBuilder = reqBuilder.WithMethod(http.MethodGet).WithRoute(INVITE_ROUTE)
		reqBuilder = reqBuilder.WithRouteVar("id", test.input.inviteId)
		r, w := reqBuilder.Build()
		h := setupInviteTest()
		h.GetInvite(w, r)
		assert.Equal(t, test.want.code, w.Code)
		assert.Equal(t, test.want.body, w.Body.String())
	}
}

func TestInviteHandler_GetInvite(t *testing.T) {
	type testInput struct {
		inviteId string
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
				inviteId: "killme",
			},
			want: testWant{
				code: http.StatusBadRequest,
				body: fmt.Sprintf("{\"message\":\"%s\"}",
					fmt.Sprintf(FAILURE_GET_INVITE, "invalid parameters")),
			},
		},
		{
			input: testInput{
				inviteId: "0",
			},
			want: testWant{
				code: http.StatusOK,
				body: fmt.Sprintf(
					"{\"Id\":%d,\"EventId\":%d,\"ReceiverId\":%d,\"StatusId\":%d}",
					inviterepo.MockInvites[0].Id, inviterepo.MockInvites[0].EventId, inviterepo.MockInvites[0].ReceiverId, inviterepo.MockInvites[0].StatusId),
			},
		},
	}
	for _, test := range tests {
		reqBuilder := testinghelp.NewRequestBuilder()

		reqBuilder = reqBuilder.WithMethod(http.MethodGet).WithRoute(INVITE_ROUTE)
		reqBuilder = reqBuilder.WithRouteVar("id", test.input.inviteId)
		r, w := reqBuilder.Build()
		h := setupInviteTest()
		h.GetInvite(w, r)
		assert.Equal(t, test.want.code, w.Code)
		assert.Equal(t, test.want.body, w.Body.String())
	}
}

// func TestInviteHandler_AddInvite(t *testing.T) {
// 	type testInput struct {
// 		body apimodels.PostComment
// 	}
// 	type testWant struct {
// 		code int
// 		body string
// 	}
// 	tests := []struct {
// 		input testInput
// 		want  testWant
// 	}{
// 		{
// 			input: testInput{
// 				body: comment_repo.MockPostComment,
// 			},
// 			want: testWant{
// 				code: http.StatusCreated,
// 				body: fmt.Sprintf("{\"id\":%d}", comment_repo.CurrCommId),
// 			},
// 		},
// 	}
// 	for _, test := range tests {
// 		reqBuilder := testinghelp.NewRequestBuilder()
// 		r, w := reqBuilder.WithMethod(http.MethodPost).WithRoute(COMMENT_ROUTE).WithBody(test.input.body).Build()
// 		h := setupTest()
// 		h.AddNotifyComment(w, r)
// 		assert.Equal(t, test.want.code, w.Code)
// 		assert.Equal(t, test.want.body, w.Body.String())
// 	}
// }

func TestInviteHandler_DeleteInvite(t *testing.T) {
	type fields struct {
		InviteController  *invitecont.InviteController
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
			handler := InviteHandler{
				InviteController:  tt.fields.InviteController,
				SessionController: tt.fields.SessionController,
			}
			handler.DeleteInvite(tt.args.w, tt.args.r)
		})
	}
}

func TestInviteHandler_PatchEventInvitesStatus(t *testing.T) {
	type fields struct {
		InviteController  *invitecont.InviteController
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
			handler := InviteHandler{
				InviteController:  tt.fields.InviteController,
				SessionController: tt.fields.SessionController,
			}
			handler.PatchEventInvitesStatus(tt.args.w, tt.args.r)
		})
	}
}

func TestInviteHandler_PatchInviteStatus(t *testing.T) {
	type fields struct {
		InviteController  *invitecont.InviteController
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
			handler := InviteHandler{
				InviteController:  tt.fields.InviteController,
				SessionController: tt.fields.SessionController,
			}
			handler.PatchInviteStatus(tt.args.w, tt.args.r)
		})
	}
}

func TestNewInviteHandler(t *testing.T) {
	type args struct {
		InviteController  *invitecont.InviteController
		SessionController *sessioncont.SessionController
	}
	tests := []struct {
		name string
		args args
		want *InviteHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInviteHandler(tt.args.InviteController, tt.args.SessionController); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInviteHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
