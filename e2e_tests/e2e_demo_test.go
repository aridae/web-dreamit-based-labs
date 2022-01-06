package e2e

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/suite"

	apiserver "github.com/aridae/web-dreamit-api-based-labs/internal/api_server/api_handlers"
	"github.com/aridae/web-dreamit-api-based-labs/internal/api_server/apimodels"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/configer"

	commentrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/comment_repo"
	eventrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/event_repo"
	inviterepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/invite_repo"
	notifyrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/notify_repo"
	roomrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/room_repo"
	sessionrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/session_repo"
	userrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/user_repo"

	"github.com/aridae/web-dreamit-api-based-labs/internal/database"

	commentcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/comment_controller"
	eventcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/event_controller"
	invitecont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/invite_controller"
	notifycont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/notify_controller"
	roomcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/room_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	usercont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/user_controller"

	testinghelp "github.com/aridae/web-dreamit-api-based-labs/internal/utils/testing"
)

type E2ETestSuite struct {
	suite.Suite

	postgresClient *database.PostgresClient
	redis_01       *database.RedisClient
	redis_02       *database.RedisClient

	EventHandler   *apiserver.EventHandler
	NotifyHandler  *apiserver.NotifyHandler
	RoomHandler    *apiserver.RoomHandler
	UserHandler    *apiserver.UserHandler
	CommentHandler *apiserver.CommentHandler
	InviteHandler  *apiserver.InviteHandler

	EventController   *eventcont.EventController
	NotifyController  *notifycont.NotifyController
	RoomController    *roomcont.RoomController
	UserController    *usercont.UserController
	CommentController *commentcont.CommentController
	InviteController  *invitecont.InviteController
}

func Test_E2ETestSuite(t *testing.T) {
	suite.Run(t, &E2ETestSuite{})
}

func (s *E2ETestSuite) SetupSuite() {
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

	redisClient0 := database.NewRedisClient(&database.RedisOptions{
		Addr:     configer.AppConfig.Redis.Addr,
		Password: configer.AppConfig.Redis.Password,
		DB:       0,
	})
	if redisClient0 == nil {
		log.Fatal(err)
	}

	redisClient1 := database.NewRedisClient(&database.RedisOptions{
		Addr:     configer.AppConfig.Redis.Addr,
		Password: configer.AppConfig.Redis.Password,
		DB:       0,
	})
	if redisClient1 == nil {
		log.Fatal(err)
	}

	s.postgresClient = postgresClient
	s.redis_01 = redisClient0
	s.redis_02 = redisClient1

	sessionRepo := sessionrepo.NewSessionRedisRepository(redisClient0, redisClient1)
	eventRepo := eventrepo.NewSessionPostgresqlRepository(postgresClient)
	userRepo := userrepo.NewSessionPostgresqlRepository(postgresClient)
	roomRepo := roomrepo.NewSessionPostgresqlRepository(postgresClient)
	inviteRepo := inviterepo.NewSessionPostgresqlRepository(postgresClient)
	notifyRepo := notifyrepo.NewSessionPostgresqlRepository(postgresClient)
	commentRepo := commentrepo.NewSessionPostgresqlRepository(postgresClient)

	sessionController := sessioncont.NewSessionController(sessionRepo)
	eventController := eventcont.NewEventController(eventRepo)
	userController := usercont.NewUserController(userRepo)
	roomController := roomcont.NewRoomController(roomRepo)
	inviteController := invitecont.NewInviteController(inviteRepo)
	notifyController := notifycont.NewNotifyController(notifyRepo)
	commentController := commentcont.NewCommentController(commentRepo)

	s.EventController = eventController
	s.NotifyController = notifyController
	s.RoomController = roomController
	s.UserController = userController
	s.CommentController = commentController
	s.InviteController = inviteController

	commentHandler := apiserver.NewCommentHandler(commentController, sessionController)
	notifyHandler := apiserver.NewNotifyHandler(notifyController, sessionController)
	inviteHandler := apiserver.NewInviteHandler(inviteController, sessionController)
	roomHandler := apiserver.NewRoomHandler(roomController, sessionController)
	userHandler := apiserver.NewUserHandler(userController, sessionController)
	eventHandler := apiserver.NewEventHandler(eventController, sessionController)

	s.CommentHandler = commentHandler
	s.NotifyHandler = notifyHandler
	s.InviteHandler = inviteHandler
	s.RoomHandler = roomHandler
	s.UserHandler = userHandler
	s.EventHandler = eventHandler
}

func (s *E2ETestSuite) TearDownSuit() {
	fmt.Println("closing test databases...")
	s.postgresClient.ClosePostgresClient()
	s.redis_01.CloseRedisClient()
	s.redis_02.CloseRedisClient()

	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *E2ETestSuite) SetupTest() {
}

func (s *E2ETestSuite) TearDownTest() {
}

func (s *E2ETestSuite) RegisterUser(usr apimodels.SignupUserRequest) string {
	reqBuilder := testinghelp.NewRequestBuilder()
	r, w := reqBuilder.WithMethod(http.MethodPost).WithRoute("/api/v2/users/signup").WithBody(usr).Build()
	s.UserHandler.SignUp(w, r)
	s.Assert().Equal(w.Code, http.StatusCreated)

	tokenBytes := w.Body.Bytes()
	var token apimodels.Token
	err := json.Unmarshal(tokenBytes, &token)
	s.Assert().NoError(err)
	return token.AccessToken
}

func (s *E2ETestSuite) Test_E2EDemo() {
	// регистрируем двух пользователей
	testAuthor_01 := apimodels.SignupUserRequest{
		Email:    "tRRETERET6585656@mail.ru",
		Login:    "Fasfssxdhfsgydifksdfeses",
		Password: "HHHfTAAAfg6333",
	}
	token_01 := s.RegisterUser(testAuthor_01)
	s.Assert().NotEmpty(token_01)
	log.Printf("registered user with token: %s\n", token_01)

	testAuthor_02 := apimodels.SignupUserRequest{
		Email:    "tETSTTSTT333@mail.ru",
		Login:    "Fahzfgysryghdfygdhfies",
		Password: "HHHfTAAAfgwetwetwet6333",
	}
	token_02 := s.RegisterUser(testAuthor_02)
	s.Assert().NotEmpty(token_02)
	log.Printf("registered user with token: %s\n", token_02)

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

	// берем токен одного пользователя и
	// от его имени создаем
	// инвайт другому пользователю
	testInvite := domain.PostInvite{
		ReceiverId: 2,
		EventId:    eventId,
	}
	_, err = s.InviteController.CreateInvite(testInvite)
	s.Assert().NoError(err)

	// создаем нотифай
	testNotify := domain.PostNotify{
		Subject: "test subject",
		EventId: eventId,
		Message: "Upd: все приходят в пижамах",
	}
	_, err = s.NotifyController.CreateNotify(testNotify)
	s.Assert().NoError(err)

	// создаем коммент
	testComment := domain.PostComment{
		AuthorId: 1,
		NotifyId: 1,
		Message:  "класс люблю пижамы",
	}
	_, err = s.CommentController.CreateComment(testComment)
	s.Assert().NoError(err)
}
