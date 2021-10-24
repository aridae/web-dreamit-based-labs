package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/middleware"
	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/configer"
	"log"
	"net/http"
	"time"

	room_delivery "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/room/handler"
	room_repo "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/room/repository"
	room_usecase "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/room/usecase"
	session_delivery "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/session/handler"
	session_repo "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/session/repository"
	session_usecase "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/session/usecase"
	user_delivery "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/user/handler"
	user_repo "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/user/repository"
	user_usecase "github.com/aridae/web-dreamit-api-based-labs/internal/pkg/user/usecase"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	configer.Init("configs/app/api_server.yaml")

	postgreSqlConn, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
			configer.AppConfig.Postgresql.User,
			configer.AppConfig.Postgresql.Password,
			configer.AppConfig.Postgresql.DBName,
			configer.AppConfig.Postgresql.Host,
			configer.AppConfig.Postgresql.Port,
			configer.AppConfig.Postgresql.Sslmode,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer postgreSqlConn.Close()
	if err := postgreSqlConn.Ping(); err != nil {
		log.Fatal(err)
	}

	// Connect to redis db #0
	redisConnDB0 := redis.NewClient(&redis.Options{
		Addr:     configer.AppConfig.Redis.Addr,
		Password: configer.AppConfig.Redis.Password,
		DB:       0,
	})
	if redisConnDB0 == nil {
		log.Fatal(err)
	}
	defer redisConnDB0.Close()

	// Connect to redis db #1
	redisConnDB1 := redis.NewClient(&redis.Options{
		Addr:     configer.AppConfig.Redis.Addr,
		Password: configer.AppConfig.Redis.Password,
		DB:       1,
	})
	if redisConnDB1 == nil {
		log.Fatal(err)
	}
	defer redisConnDB1.Close()



	sessionRepo := session_repo.NewSessionRedisRepository(redisConnDB0, redisConnDB1)
	sessionUCase := session_usecase.NewUseCase(sessionRepo)
	sessionHandler := session_delivery.NewHandler(sessionUCase)

	roomRepo := room_repo.NewSessionPostgresqlRepository(postgreSqlConn)
	roomUCase := room_usecase.NewUseCase(roomRepo)
	roomHandler := room_delivery.NewHandler(roomUCase, sessionUCase)

	userRepo := user_repo.NewRepository(postgreSqlConn)
	userUCase := user_usecase.NewUseCase(userRepo)
	userHandler := user_delivery.NewHandler(userUCase, sessionUCase)

	mainMux := mux.NewRouter()
	mainMux.Use(middleware.Cors)
	mainMux.HandleFunc("/api/v1/room", roomHandler.GetAllRooms).Methods("GET")
	mainMux.HandleFunc("/api/v1/room/my", roomHandler.MyRoomBooking).Methods("GET")
	mainMux.HandleFunc("/api/v1/room/{id:[0-9]+}", roomHandler.GetRoomCalendar).Methods("GET")
	mainMux.HandleFunc("/api/v1/room/{id:[0-9]+}", roomHandler.AddRoomBooking).Methods("POST")
	mainMux.HandleFunc("/api/v1/room/{bookingId}", roomHandler.DeleteRoomBooking).Methods("DELETE")
	mainMux.HandleFunc("/api/v1/room/{id:[0-9]+}/{date}", roomHandler.GetRoomSchedule).Methods("GET")
	mainMux.HandleFunc("/api/v1/room/{id:[0-9]+}/{date}", roomHandler.UpdateRoomBooking).Methods("POST")

	mainMux.HandleFunc("/api/v1/user/signup", userHandler.SignUp).Methods("POST")
	mainMux.HandleFunc("/api/v1/user/login", userHandler.LogIn).Methods("POST")
	mainMux.HandleFunc("/api/v1/user/oauth/keycloak", userHandler.LogInKeycloak).Methods("GET")

	mainMux.HandleFunc("/api/v1/session/refresh", sessionHandler.RefreshSession).Methods("GET")
	mainMux.HandleFunc("/api/v1/session/check", sessionHandler.CheckSession).Methods("GET")

	server := &http.Server{
		Addr:         ":80",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mainMux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
