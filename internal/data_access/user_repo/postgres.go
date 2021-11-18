package userrepo

import (
	"database/sql"
	"fmt"

	db "github.com/aridae/web-dreamit-api-based-labs/internal/database"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"github.com/aridae/web-dreamit-api-based-labs/pkg/tools/db_sql"
)

type PostgresqlRepository struct {
	db *db.PostgresClient
}

func dbToDomainUserData(data *UserData) *domain.UserData {
	return &domain.UserData{
		Id:         data.Id,
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		Email:      data.Email,
		Login:      data.Login,
		Avatar:     data.Avatar,
		Background: data.Background,
		Password:   data.Password,
	}
}

func (r *PostgresqlRepository) InsertUser(user *domain.UserData) (uint64, error) {
	var userId uint64
	err := r.db.Client.Get(
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

func (r *PostgresqlRepository) GetUsers() ([]domain.UserProfile, error) {
	rows, err := r.db.Client.Query("SELECT id, first_name, last_name, login, email, avatar, background FROM users ")
	if err != nil {
		return nil, err
	}

	user := new(domain.UserProfile)
	users := make([]domain.UserProfile, 0)
	for rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Login,
			&user.Email,
			&user.Avatar,
			&user.Background,
		); err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return users, nil
}

func (r *PostgresqlRepository) SelectUserByEmailOrLogin(emailOrLogin string) (*domain.UserData, error) {
	row := r.db.Client.QueryRow(
		"SELECT id, first_name, last_name, login, email, password "+
			"FROM users "+
			"WHERE email = $1 or login = $1",
		emailOrLogin,
	)

	userData := &UserData{}

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

	return dbToDomainUserData(userData), nil
}

func (r *PostgresqlRepository) SelectUserById(userId uint64) (*domain.UserProfile, error) {
	row := r.db.Client.QueryRow(
		"SELECT id, first_name, last_name, login, email, avatar, background "+
			"FROM users "+
			"WHERE id = $1",
		userId,
	)

	userData := &domain.UserProfile{}

	firstName := sql.NullString{}
	lastName := sql.NullString{}
	email := sql.NullString{}
	err := row.Scan(
		&userData.Id,
		&firstName,
		&lastName,
		&userData.Login,
		&email,
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

func (r *PostgresqlRepository) SelectUserByAuthId(authId uint64, authService string) (*domain.UserData, error) {
	row := r.db.Client.QueryRow(
		"SELECT u.id, u.first_name, u.last_name, u.login, u.email, u.password "+
			"FROM users u "+
			"JOIN auth_tokens at ON u.id = at.user_id "+
			"JOIN auth_services asv ON at.service_id = asv.id "+
			"WHERE at.auth_id = $1 AND asv.service = $2",
		authId,
		authService,
	)

	userData := &UserData{}

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

	return dbToDomainUserData(userData), nil
}

func (r *PostgresqlRepository) InsertAuthUser(user *domain.AuthUserData) (uint64, error) {
	var serviceId uint64
	err := r.db.Client.Get(
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
	tx := r.db.Client.MustBegin()
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

func (r *PostgresqlRepository) SelectNewUniqLogin(login string) (string, error) {
	var countLogins uint64
	err := r.db.Client.Get(
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

func (r *PostgresqlRepository) DeleteSelfProfile(userId uint64) error {
	_, err := r.db.Client.Exec(
		"DELETE FROM users "+
			"WHERE id = $1",
		userId,
	)

	if err != nil {
		return err
	}

	// TODO: NEED CASCADE DELETE or mb 1 tr
	_, _ = r.db.Client.Exec(
		"DELETE FROM auth_tokens "+
			"WHERE auth_id = $1",
		userId,
	)

	return nil
}

func NewSessionPostgresqlRepository(client *db.PostgresClient) Repository {
	return &PostgresqlRepository{
		db: client,
	}
}
