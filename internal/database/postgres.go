package database

import (
	"context"
	"fmt"
	"log"

	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/configer"
	"github.com/jmoiron/sqlx"
)

type PostgresClient struct {
	Client *sqlx.DB
}

var (
	postgresClient *PostgresClient
)

type Options struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

func NewPostgresClient(ctx context.Context, options *Options) (*PostgresClient, error) {
	if postgresClient == nil {
		client, err := getPostgresClient(ctx, options)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize postgres client, error is: %s", err)
		}
		postgresClient = &PostgresClient{
			Client: client,
		}
		return postgresClient, nil
	}

	return postgresClient, nil
}

func (client *PostgresClient) ClosePostgresClient() {
	client.Client.Close()
}

func getPostgresClient(ctx context.Context, options *Options) (*sqlx.DB, error) {

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
		return nil, err
	}
	if err := postgreSqlConn.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return postgreSqlConn, nil
}
