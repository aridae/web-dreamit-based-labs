//go:build integration
// +build integration

package integration_tests

import (
	"context"
	"log"
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/suite"

	commentrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/comment_repo"
	eventrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/event_repo"
	notifyrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/notify_repo"
	roomrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/room_repo"
	userrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/user_repo"
	"github.com/aridae/web-dreamit-api-based-labs/internal/database"

	commentcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/comment_controller"
	eventcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/event_controller"
	invitecont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/invite_controller"
	notifycont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/notify_controller"
	roomcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/room_controller"
	usercont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/user_controller"

	domain "github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

// тестим слой контроллеров, так как тут
// уже не мокаем все внешние зависимости
// цель - протестировать внешние зависимости,
// в т.ч. взаимодействие с бд
type EventXInviteTestSuite struct {
	suite.Suite
	postgresClient *database.PostgresClient

	EventController   *eventcont.EventController
	NotifyController  *notifycont.NotifyController
	RoomController    *roomcont.RoomController
	UserController    *usercont.UserController
	CommentController *commentcont.CommentController
	InviteController  *invitecont.InviteController
}

func TestEventXInviteTestSuite(t *testing.T) {
	suite.Run(t, &EventXInviteTestSuite{})
}

func (s *EventXInviteTestSuite) SetupSuite() {
	//configer.Init("./test_api_server.yaml")
	postgresClient, err := database.NewPostgresClient(context.Background(), &database.Options{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DB:       "dreamit_api_db",
	})
	if err != nil {
		log.Fatal(err)
	}
	s.postgresClient = postgresClient

	//sessionRepo := sessionrepo.NewRepositoryMock() // мокаем
	eventRepo := eventrepo.NewSessionPostgresqlRepository(postgresClient)
	userRepo := userrepo.NewSessionPostgresqlRepository(postgresClient)
	roomRepo := roomrepo.NewSessionPostgresqlRepository(postgresClient)
	notifyRepo := notifyrepo.NewSessionPostgresqlRepository(postgresClient)
	commentRepo := commentrepo.NewSessionPostgresqlRepository(postgresClient)

	// sessionController := sessioncont.NewSessionController(sessionRepo)
	eventController := eventcont.NewEventController(eventRepo)
	userController := usercont.NewUserController(userRepo)
	roomController := roomcont.NewRoomController(roomRepo)
	notifyController := notifycont.NewNotifyController(notifyRepo)
	commentController := commentcont.NewCommentController(commentRepo)

	s.EventController = eventController
	s.UserController = userController
	s.RoomController = roomController
	s.NotifyController = notifyController
	s.CommentController = commentController
}

func (s *EventXInviteTestSuite) TearDownSuite() {
	s.postgresClient.ClosePostgresClient()
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *EventXInviteTestSuite) SetupTest() {
}

func (s *EventXInviteTestSuite) TearDownTest() {
}

func (s *EventXInviteTestSuite) Test_EventXNotify_CreateNotify() {
	// создаем ивент
	testEvent := domain.PostEvent{
		AuthorId: 1,
		End:      "2021-12-03 12:00",
		RoomId:   1,
		Start:    "2021-12-03 11:00",
		Title:    "Confa!",
	}
	eventId, err := s.EventController.AddRoomEvent(testEvent)
	s.Assert().NoError(err)

	// создаем нотифай
	testNotify := domain.PostNotify{
		Subject: "test subject",
		EventId: eventId,
		Message: "Upd: все приходят в пижамах",
	}
	_, err = s.NotifyController.CreateNotify(testNotify)
	s.Assert().NoError(err)
}

func (s *EventXInviteTestSuite) Test_EventXInvite_CreateInvite() {
	// создаем ивент
	testEvent := domain.PostEvent{
		End:      "2021-12-03 12:00",
		RoomId:   1,
		Start:    "2021-12-03 11:00",
		Title:    "Confa!",
		AuthorId: 1,
	}
	eventId, err := s.EventController.AddRoomEvent(testEvent)
	s.Assert().NoError(err)

	// создаем инвайт
	testInvite := domain.PostInvite{
		ReceiverId: 2,
		EventId:    eventId,
	}
	_, err = s.InviteController.CreateInvite(testInvite)
	s.Assert().NoError(err)

	// принимаем инвайт
	//err = s.InviteController.AcceptInvite(inviteId)
	//s.Assert().NoError(err)
}
