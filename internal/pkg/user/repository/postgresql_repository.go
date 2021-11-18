package repository

import (
	"database/sql"
	"fmt"

	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/user"
	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/db_sql"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	dbConn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) user.Repository {
	return &UserRepository{
		dbConn: conn,
	}
}

func (r *UserRepository) InsertUser(user *api_models.UserData) (uint64, error) {
	var userId uint64
	err := r.dbConn.Get(
		&userId,
		"INSERT INTO users(first_name, last_name, login, email, password) "+
			"VALUES ($1, $2, $3, $4, $5) RETURNING id",
		db_sql.NewNullString(user.FirstName),
		db_sql.NewNullString(user.LastName),
		user.Login,
		db_sql.NewNullString(user.Email),
		user.Password,
	)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *UserRepository) SelectUserByEmailOrLogin(emailOrLogin string) (*api_models.UserData, error) {
	row := r.dbConn.QueryRow(
		"SELECT id, first_name, last_name, login, email, password "+
			"FROM users "+
			"WHERE email = $1 or login = $1",
		emailOrLogin,
	)

	userData := &api_models.UserData{}

	firstName := sql.NullString{}
	lastName := sql.NullString{}
	email := sql.NullString{}
	err := row.Scan(
		&userData.Id,
		&firstName,
		&lastName,
		&userData.Login,
		&email,
		&userData.Password,
	)
	userData.FirstName = firstName.String
	userData.LastName = lastName.String
	userData.Email = email.String

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (r *UserRepository) SelectUserById(userId uint64) (*api_models.UserData, error) {
	row := r.dbConn.QueryRow(
		"SELECT id, first_name, last_name, login, email, password, avatar, background "+
			"FROM users "+
			"WHERE id = $1",
		userId,
	)

	userData := &api_models.UserData{}

	firstName := sql.NullString{}
	lastName := sql.NullString{}
	email := sql.NullString{}
	err := row.Scan(
		&userData.Id,
		&firstName,
		&lastName,
		&userData.Login,
		&email,
		&userData.Password,
		&userData.Avatar,
		&userData.Background,
	)
	userData.FirstName = firstName.String
	userData.LastName = lastName.String
	userData.Email = email.String

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (r *UserRepository) SelectUserByAuthId(authId uint64, authService string) (*api_models.UserData, error) {
	row := r.dbConn.QueryRow(
		"SELECT u.id, u.first_name, u.last_name, u.login, u.email, u.password "+
			"FROM users u "+
			"JOIN auth_tokens at ON u.id = at.user_id "+
			"JOIN auth_services asv ON at.service_id = asv.id "+
			"WHERE at.auth_id = $1 AND asv.service = $2",
		authId,
		authService,
	)

	userData := &api_models.UserData{}

	firstName := sql.NullString{}
	lastName := sql.NullString{}
	email := sql.NullString{}
	err := row.Scan(
		&userData.Id,
		&firstName,
		&lastName,
		&userData.Login,
		&email,
		&userData.Password,
	)
	userData.FirstName = firstName.String
	userData.LastName = lastName.String
	userData.Email = email.String

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (r *UserRepository) InsertAuthUser(user *api_models.AuthUserData) (uint64, error) {
	var serviceId uint64
	err := r.dbConn.Get(
		&serviceId,
		"SELECT id "+
			"FROM auth_services "+
			"WHERE service = $1",
		user.ServiceType,
	)

	if err != nil {
		return 0, err
	}

	var userId uint64
	tx := r.dbConn.MustBegin()
	err = tx.Get(
		&userId,
		"INSERT INTO users(first_name, last_name, login, password) "+
			"VALUES ($1, $2, $3, $4) RETURNING id",
		db_sql.NewNullString(user.FirstName),
		db_sql.NewNullString(user.LastName),
		user.Login,
		user.Password,
	)
	if err != nil {
		return 0, err
	}

	tx.MustExec(
		"INSERT INTO auth_tokens(auth_id, service_id, access_token, user_id, refresh_token) "+
			"VALUES ($1, $2, $3, $4, $5)",
		user.AuthId,
		serviceId,
		user.AccessToken,
		userId,
		db_sql.NewNullString(user.RefreshToken),
	)
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *UserRepository) SelectNewUniqLogin(login string) (string, error) {
	var countLogins uint64
	err := r.dbConn.Get(
		&countLogins,
		"SELECT COUNT(*) "+
			"FROM users "+
			"WHERE login LIKE $1",
		login+"#%",
	)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s#%d", login, countLogins+1), nil
}

func (r *UserRepository) DeleteSelfProfile(userId uint64) error {
	_, err := r.dbConn.Exec(
		"DELETE FROM users "+
			"WHERE id = $1",
		userId,
	)

	if err != nil {
		return err
	}

	// TODO: NEED CASCADE DELETE or mb 1 tr
	_, _ = r.dbConn.Exec(
		"DELETE FROM auth_tokens "+
			"WHERE auth_id = $1",
		userId,
	)

	return nil
}
