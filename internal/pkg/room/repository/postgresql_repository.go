package repository

import (
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/api_models"
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/room"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/errors"
	"github.com/jmoiron/sqlx"
)

type PostgresqlRepository struct {
	db *sqlx.DB
}

func (p PostgresqlRepository) DeleteRoomEvent(userId uint64, eventId int64) error {
	row := p.db.QueryRow("DELETE FROM calendar "+
		"WHERE author = $1 AND id = $2", userId, eventId)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (p PostgresqlRepository) GetRoomEventsByUserId(userId uint64) ([]api_models.Event, error) {
	rows, err := p.db.Query("SELECT id, roomId, title, start, \"end\" "+
		"FROM calendar WHERE author = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	event := api_models.Event{}
	result := make([]api_models.Event, 0)
	for rows.Next() {
		if err := rows.Scan(
			&event.Id,
			&event.RoomId,
			&event.Title,
			&event.Start,
			&event.End,
		); err != nil {
			return nil, err
		}
		result = append(result, event)
	}
	return result, nil
}

func (p PostgresqlRepository) AddRoomEvent(event api_models.Event) (int64, error) {
	check := p.db.QueryRow("SELECT COUNT(*) FROM calendar "+
		"WHERE roomId = $1 AND (start < $2 AND $2 < \"end\" OR \"end\" < $3 AND $3 < start )",
		event.RoomId, event.Start, event.End)
	var count int64
	if err := check.Err(); err != nil {
		return 0, err
	}
	if err := check.Scan(&count); err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, errors.ErrBadTime
	}

	row := p.db.QueryRow("INSERT INTO calendar(roomId, title, start, \"end\", author) "+
		"VALUES ($1, $2, $3, $4, $5) RETURNING id", event.RoomId, event.Title, event.Start, event.End, event.Author)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func (p PostgresqlRepository) GetAllRooms() ([]api_models.Room, error) {
	rows, err := p.db.Query("SELECT id, title " +
		"FROM rooms ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := make([]api_models.Room, 0)

	var roomInfo api_models.Room
	for rows.Next() {
		if err := rows.Scan(
			&roomInfo.Id,
			&roomInfo.Title,
		); err != nil {
			return nil, err
		}
		rooms = append(rooms, roomInfo)
	}
	return rooms, nil
}

func (p PostgresqlRepository) GetRoomEventsByRoomId(roomId int64) ([]api_models.Event, error) {
	rows, err := p.db.Query("SELECT id, roomId, title, start, \"end\" "+
		"FROM calendar "+
		"WHERE roomId = $1", roomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var event api_models.Event
	calendar := make([]api_models.Event, 0)
	for rows.Next() {
		if err := rows.Scan(
			&event.Id,
			&event.RoomId,
			&event.Title,
			&event.Start,
			&event.End,
		); err != nil {
			return nil, err
		}
		calendar = append(calendar, event)
	}

	return calendar, nil
}

func (p PostgresqlRepository) UpdateRoomEventsByRoomId(roomId int64, event api_models.Event) error {
	check := p.db.QueryRow("SELECT COUNT(*) FROM calendar "+
		"WHERE roomId = $1 AND (start < $2 AND $2 < \"end\" OR \"end\" < $3 AND $3 < start )",
		roomId, event.Start, event.End)
	var count int64
	if err := check.Err(); err != nil {
		return err
	}
	if err := check.Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return errors.ErrBadTime
	}

	row := p.db.QueryRow("UPDATE calendar "+
		"SET start = $1, \"end\" = $2 "+
		"WHERE roomId = $3 ", event.Start, event.End, roomId)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func NewSessionPostgresqlRepository(db *sqlx.DB) room.Repository {
	return &PostgresqlRepository{
		db: db,
	}
}
