package eventrepo

import (
	"fmt"

	db "github.com/aridae/web-dreamit-api-based-labs/internal/database"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/errors"
)

type PostgresqlRepository struct {
	db *db.PostgresClient
}

func (p PostgresqlRepository) GetEvent(eventId int64) (*domain.Event, error) {
	row := p.db.Client.QueryRow("SELECT id, roomId, title, start, \"end\" FROM calendar WHERE id=$1;", eventId)
	if row.Err() != nil {
		return nil, row.Err()
	}

	event := new(domain.Event)
	if err := row.Scan(
		&event.Id,
		&event.RoomId,
		&event.Title,
		&event.Start,
		&event.End,
	); err != nil {
		return nil, err
	}

	return event, nil
}

func (p PostgresqlRepository) DeleteRoomEvent(eventId int64) error {
	row := p.db.Client.QueryRow("DELETE FROM calendar "+
		"WHERE id = $1", eventId)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (p PostgresqlRepository) GetEvents() ([]domain.Event, error) {
	rows, err := p.db.Client.Query("SELECT id, roomId, title, start, \"end\" FROM calendar;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]domain.Event, 0)

	eventDB := Event{}
	for rows.Next() {
		if err := rows.Scan(
			&eventDB.Id,
			&eventDB.RoomId,
			&eventDB.Title,
			&eventDB.Start,
			&eventDB.End,
		); err != nil {
			return nil, err
		}
		events = append(events, domain.Event{
			Id:       eventDB.Id,
			RoomId:   eventDB.RoomId,
			Title:    eventDB.Title,
			Start:    eventDB.Start,
			End:      eventDB.End,
			AuthorId: eventDB.AuthorId,
		})
	}
	return events, nil
}

func (p PostgresqlRepository) GetRoomEventsByUserId(userId uint64) ([]domain.Event, error) {
	rows, err := p.db.Client.Query("SELECT id, roomId, title, start, \"end\" "+
		"FROM calendar WHERE author = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]domain.Event, 0)

	eventDB := Event{}
	for rows.Next() {
		if err := rows.Scan(
			&eventDB.Id,
			&eventDB.RoomId,
			&eventDB.Title,
			&eventDB.Start,
			&eventDB.End,
		); err != nil {
			return nil, err
		}
		events = append(events, domain.Event{
			Id:       eventDB.Id,
			RoomId:   eventDB.RoomId,
			Title:    eventDB.Title,
			Start:    eventDB.Start,
			End:      eventDB.End,
			AuthorId: eventDB.AuthorId,
		})
	}
	return events, nil
}

func (p PostgresqlRepository) AddRoomEvent(event domain.PostEvent) (int64, error) {

	fmt.Println("in post event Repository")

	check := p.db.Client.QueryRow("SELECT COUNT(*) FROM calendar "+
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

	row := p.db.Client.QueryRow("INSERT INTO calendar(roomId, title, start, \"end\", author) "+
		"VALUES ($1, $2, $3, $4, $5) RETURNING id", event.RoomId, event.Title, event.Start, event.End, event.AuthorId)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	fmt.Printf("in post event id=%d\n", id)
	return id, nil
}

func (p PostgresqlRepository) GetRoomEventsByRoomId(roomId int64) ([]domain.Event, error) {
	rows, err := p.db.Client.Query("SELECT id, roomId, title, start, \"end\" "+
		"FROM calendar "+
		"WHERE roomId = $1", roomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]domain.Event, 0)

	eventDB := Event{}
	for rows.Next() {
		if err := rows.Scan(
			&eventDB.Id,
			&eventDB.RoomId,
			&eventDB.Title,
			&eventDB.Start,
			&eventDB.End,
		); err != nil {
			return nil, err
		}
		events = append(events, domain.Event{
			Id:       eventDB.Id,
			RoomId:   eventDB.RoomId,
			Title:    eventDB.Title,
			Start:    eventDB.Start,
			End:      eventDB.End,
			AuthorId: eventDB.AuthorId,
		})
	}

	return events, nil
}

func (p PostgresqlRepository) RescheduleRoomEvent(eventId int64, event domain.PatchEvent) error {
	check := p.db.Client.QueryRow("SELECT COUNT(*) FROM calendar "+
		"WHERE roomId = $1 AND (start < $2 AND $2 < \"end\" OR \"end\" < $3 AND $3 < start )",
		event.RoomId, event.Start, event.End)
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

	row := p.db.Client.QueryRow("UPDATE calendar "+
		"SET start = $1, \"end\" = $2, roomId = $3 "+
		"WHERE id = $4  ", event.Start, event.End, event.RoomId, event.Id)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func NewSessionPostgresqlRepository(client *db.PostgresClient) Repository {
	return &PostgresqlRepository{
		db: client,
	}
}
