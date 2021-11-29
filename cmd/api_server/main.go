package main

import (
	"context"
	"log"
	"net/http"
	"time"

	commentrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/comment_repo"
	eventrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/event_repo"
	inviterepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/invite_repo"
	notifyrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/notify_repo"
	roomrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/room_repo"
	sessionrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/session_repo"
	userrepo "github.com/aridae/web-dreamit-api-based-labs/internal/data_access/user_repo"

	"github.com/aridae/web-dreamit-api-based-labs/internal/database"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/middleware"
	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/configer"

	apiserver "github.com/aridae/web-dreamit-api-based-labs/internal/api_server/api_handlers"
	middlewarev2 "github.com/aridae/web-dreamit-api-based-labs/internal/api_server/middleware"

	commentcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/comment_controller"
	eventcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/event_controller"
	invitecont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/invite_controller"
	notifycont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/notify_controller"
	roomcont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/room_controller"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	usercont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/user_controller"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	_ "github.com/aridae/web-dreamit-api-based-labs/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Dreamit Swagger API
// @version 2.0
// @description Swagger API for Dreamit-based web labs
// @termsOfService http://swagger.io/terms/

// @BasePath /api/v2
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configer.Init("configs/app/api_server.yaml")

	// СТАРАЯ ВЕРСИЯ АПИ БУДЕТ МАРШУТИЗИРОВАНА НА /ЛЕГАСИ НГИНХОМ??
	// если я правильно понимаю, мы документируем только v2 в мейне,
	// тк тут указывается basepath /api/v2 и version 2.0
	// **********************************************************
	//registerAPIv1(mainMux, roomHandler, userHandler, sessionHandler)

	postgresClient, err := database.NewPostgresClient(context.Background(), &database.Options{
		Host:     configer.AppConfig.Postgresql.Host,
		Port:     configer.AppConfig.Postgresql.Port,
		User:     configer.AppConfig.Postgresql.User,
		Password: configer.AppConfig.Postgresql.Password,
		DB:       configer.AppConfig.Postgresql.DBName,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer postgresClient.ClosePostgresClient()

	redisClient0 := database.NewRedisClient(&database.RedisOptions{
		Addr:     configer.AppConfig.Redis.Addr,
		Password: configer.AppConfig.Redis.Password,
		DB:       0,
	})
	if redisClient0 == nil {
		log.Fatal(err)
	}
	defer redisClient0.CloseRedisClient()

	redisClient1 := database.NewRedisClient(&database.RedisOptions{
		Addr:     configer.AppConfig.Redis.Addr,
		Password: configer.AppConfig.Redis.Password,
		DB:       0,
	})
	if redisClient1 == nil {
		log.Fatal(err)
	}
	defer redisClient1.CloseRedisClient()

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

	commentHandler := apiserver.NewCommentHandler(commentController, sessionController)
	notifyHandler := apiserver.NewNotifyHandler(notifyController, sessionController)
	inviteHandler := apiserver.NewInviteHandler(inviteController, sessionController)
	roomHandler := apiserver.NewRoomHandler(roomController, sessionController)
	userHandler := apiserver.NewUserHandler(userController, sessionController)
	eventHandler := apiserver.NewEventHandler(eventController, sessionController)
	jwtHandler := &middlewarev2.JWTHandler{
		SessionController: sessionController,
	}
	accessHandler := &middlewarev2.AccessHandler{
		SessionController: sessionController,
	}

	mainMux := mux.NewRouter()
	mainMux.Use(middleware.Cors)
	mainMux.Use(accessHandler.CheckAccess)
	mainMux.Use(jwtHandler.JWTAuth)

	// сваггером нужно аннотировать все ендпоинты:
	mainMux.HandleFunc("/api/v2/rooms", roomHandler.GetAllRooms).Methods("GET")
	mainMux.HandleFunc("/api/v2/rooms/{id:[0-9]+}", roomHandler.GetRoom).Methods("GET")

	mainMux.HandleFunc("/api/v2/events", eventHandler.GetEventsCollection).Methods("GET")
	mainMux.HandleFunc("/api/v2/events", eventHandler.PostEvent).Methods("POST")
	mainMux.HandleFunc("/api/v2/events/{id:[0-9]+}", eventHandler.DeleteEvent).Methods("DELETE")
	mainMux.HandleFunc("/api/v2/events/{id:[0-9]+}", eventHandler.PatchEvent).Methods("PATCH")
	mainMux.HandleFunc("/api/v2/events/{id:[0-9]+}", eventHandler.GetEvent).Methods("GET")

	mainMux.HandleFunc("/api/v2/invites", inviteHandler.GetInvites).Methods("GET")
	mainMux.HandleFunc("/api/v2/invites", inviteHandler.AddInvite).Methods("POST")
	mainMux.HandleFunc("/api/v2/invites", inviteHandler.PatchEventInvitesStatus).Methods("PATCH")
	mainMux.HandleFunc("/api/v2/invites/{id:[0-9]+}", inviteHandler.DeleteInvite).Methods("DELETE")
	mainMux.HandleFunc("/api/v2/invites/{id:[0-9]+}", inviteHandler.GetInvite).Methods("GET")
	mainMux.HandleFunc("/api/v2/invites/{id:[0-9]+}", inviteHandler.PatchInviteStatus).Methods("PATCH")

	mainMux.HandleFunc("/api/v2/notifies", notifyHandler.GetNotifies).Methods("GET")
	mainMux.HandleFunc("/api/v2/notifies", notifyHandler.AddNotify).Methods("POST")
	mainMux.HandleFunc("/api/v2/notifies/{id:[0-9]+}", notifyHandler.DeleteNotify).Methods("DELETE")
	mainMux.HandleFunc("/api/v2/notifies/{id:[0-9]+}", notifyHandler.GetNotify).Methods("GET")

	mainMux.HandleFunc("/api/v2/comments", commentHandler.GetNotifyComments).Methods("GET")
	mainMux.HandleFunc("/api/v2/comments/{id:[0-9]+}", commentHandler.GetNotifyComment).Methods("GET")
	mainMux.HandleFunc("/api/v2/comments/{id:[0-9]+}", commentHandler.AddNotifyComment).Methods("POST")
	mainMux.HandleFunc("/api/v2/comments/{id:[0-9]+}", commentHandler.DeleteNotifyComment).Methods("DELETE")

	mainMux.HandleFunc("/api/v2/users", userHandler.GetUsers).Methods("GET")
	mainMux.HandleFunc("/api/v2/users/{id:[0-9]+}", userHandler.GetUser).Methods("GET")

	mainMux.HandleFunc("/api/v2/users/signup", userHandler.SignUp).Methods("POST")
	mainMux.HandleFunc("/api/v2/users/login", userHandler.LogIn).Methods("POST")
	mainMux.HandleFunc("/api/v2/users/logout", userHandler.Logout).Methods("POST")

	// враппер генерит index.html
	mainMux.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mainMux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// func registerAPIv1(mainMux *mux.Router, roomHandler room.Handler, userHandler user.Handler, sessionHandler session.Handler) {
// 	mainMux.HandleFunc("/api/v1/room", roomHandler.GetAllRooms).Methods("GET")
// 	mainMux.HandleFunc("/api/v1/room/my", roomHandler.MyRoomEvents).Methods("GET")
// 	mainMux.HandleFunc("/api/v1/room/{id:[0-9]+}", roomHandler.GetRoomEvents).Methods("GET")
// 	mainMux.HandleFunc("/api/v1/room/{id:[0-9]+}", roomHandler.AddRoomEvent).Methods("POST")
// 	mainMux.HandleFunc("/api/v1/room/{bookingId}", roomHandler.DeleteRoomEvent).Methods("DELETE")
// 	mainMux.HandleFunc("/api/v1/room/{id:[0-9]+}/{date}", roomHandler.UpdateRoomEvent).Methods("POST")

// 	mainMux.HandleFunc("/api/v1/user/signup", userHandler.SignUp).Methods("POST")
// 	mainMux.HandleFunc("/api/v1/user/login", userHandler.LogIn).Methods("POST")
// 	mainMux.HandleFunc("/api/v1/user/oauth/keycloak", userHandler.LogInKeycloak).Methods("GET")

// 	mainMux.HandleFunc("/api/v1/session/refresh", sessionHandler.RefreshSession).Methods("GET")
// 	mainMux.HandleFunc("/api/v1/session/check", sessionHandler.CheckSession).Methods("GET")
// }
