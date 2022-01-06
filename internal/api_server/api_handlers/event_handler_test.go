//go:build unit
// +build unit

package apiserver

import (
	"fmt"
	"net/http"
	"testing"

	eventcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/event_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	eventrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/event_repo"
	sessionrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/session_repo"
	testinghelp "github.com/aridae/web-dreamit-api-based-labs/internal/utils/testing"
	"github.com/stretchr/testify/assert"
)

var (
	EVENT_ROUTE = "/events"
)

func setupEventTest() *EventHandler {
	h := &EventHandler{
		EventController:   eventcont.NewEventController(eventrepo.NewRepositoryMock()),
		SessionController: sessioncont.NewSessionController(&sessionrepo.RepositoryMock{}),
	}
	return h
}

func TestEventHandler_GetEventsCollection(t *testing.T) {
	type testInput struct {
		authorId string
		roomId   string
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
				authorId: "killme",
			},
			want: testWant{
				code: http.StatusBadRequest,
				body: fmt.Sprintf("{\"message\":\"%s\"}",
					fmt.Sprintf(FAILURE_AUTHOR_EVENTS, "invalid parameters")),
			},
		},
		{
			input: testInput{
				roomId: "1",
			},
			want: testWant{
				code: http.StatusOK,
				body: fmt.Sprintf(
					"[{\"Id\":%d,\"RoomId\":%d,\"Title\":\"%s\",\"Start\":\"%s\",\"End\":\"%s\",\"AuthorId\":%d}]",
					eventrepo.MockEvents[0].Id, eventrepo.MockEvents[0].RoomId, eventrepo.MockEvents[0].Title, eventrepo.MockEvents[0].Start, eventrepo.MockEvents[0].End, eventrepo.MockEvents[0].AuthorId),
			},
		},
		{
			input: testInput{
				roomId: "killme",
			},
			want: testWant{
				code: http.StatusBadRequest,
				body: fmt.Sprintf("{\"message\":\"%s\"}",
					fmt.Sprintf(FAILURE_ROOM_EVENTS, "invalid parameters")),
			},
		},
		{
			input: testInput{
				authorId: "1",
			},
			want: testWant{
				code: http.StatusOK,
				body: fmt.Sprintf(
					"[{\"Id\":%d,\"RoomId\":%d,\"Title\":\"%s\",\"Start\":\"%s\",\"End\":\"%s\",\"AuthorId\":%d}]",
					eventrepo.MockEvents[0].Id, eventrepo.MockEvents[0].RoomId, eventrepo.MockEvents[0].Title, eventrepo.MockEvents[0].Start, eventrepo.MockEvents[0].End, eventrepo.MockEvents[0].AuthorId),
			},
		},
		{
			input: testInput{},
			want: testWant{
				code: http.StatusOK,
				body: fmt.Sprintf(
					"[{\"Id\":%d,\"RoomId\":%d,\"Title\":\"%s\",\"Start\":\"%s\",\"End\":\"%s\",\"AuthorId\":%d},{\"Id\":%d,\"RoomId\":%d,\"Title\":\"%s\",\"Start\":\"%s\",\"End\":\"%s\",\"AuthorId\":%d},{\"Id\":%d,\"RoomId\":%d,\"Title\":\"%s\",\"Start\":\"%s\",\"End\":\"%s\",\"AuthorId\":%d}]",
					eventrepo.MockEvents[0].Id, eventrepo.MockEvents[0].RoomId, eventrepo.MockEvents[0].Title, eventrepo.MockEvents[0].Start, eventrepo.MockEvents[0].End, eventrepo.MockEvents[0].AuthorId,
					eventrepo.MockEvents[1].Id, eventrepo.MockEvents[1].RoomId, eventrepo.MockEvents[1].Title, eventrepo.MockEvents[1].Start, eventrepo.MockEvents[1].End, eventrepo.MockEvents[1].AuthorId,
					eventrepo.MockEvents[2].Id, eventrepo.MockEvents[2].RoomId, eventrepo.MockEvents[2].Title, eventrepo.MockEvents[2].Start, eventrepo.MockEvents[2].End, eventrepo.MockEvents[2].AuthorId,
				),
			},
		},
	}
	for _, test := range tests {
		reqBuilder := testinghelp.NewRequestBuilder()

		reqBuilder = reqBuilder.WithMethod(http.MethodGet).WithRoute(EVENT_ROUTE)
		if test.input.authorId != "" {
			reqBuilder = reqBuilder.WithParameter("authorId", test.input.authorId)
		}
		if test.input.roomId != "" {
			reqBuilder = reqBuilder.WithParameter("roomId", test.input.roomId)
		}
		r, w := reqBuilder.Build()
		h := setupEventTest()
		h.GetEventsCollection(w, r)
		assert.Equal(t, test.want.code, w.Code)
		assert.Equal(t, test.want.body, w.Body.String())
	}
}
