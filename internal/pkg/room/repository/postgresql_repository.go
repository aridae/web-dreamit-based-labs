package repository

import (
	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/internal/server/errors"
	"github.com/jmoiron/sqlx"
	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/internal/pkg/models"
	"lab.qoollo.com/practice/2021/dreamit/dreamit-api/internal/pkg/room"
	"time"
)

type PostgresqlRepository struct {
	db *sqlx.DB
}

func (p PostgresqlRepository) DeleteRoomBooking(userId uint64, eventId int64) error {
	row := p.db.QueryRow("DELETE FROM calendar " +
		"WHERE author = $1 AND id = $2", userId, eventId)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (p PostgresqlRepository) MyRoomBooking(userId uint64) ([]models.Booking, error) {
	rows, err := p.db.Query("SELECT id, roomId, title, start, \"end\" " +
		"FROM calendar WHERE author = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	booking := models.Booking{}
	result := make([]models.Booking, 0)
	for rows.Next() {
		if err := rows.Scan(
			&booking.Booking.Id,
			&booking.RoomId,
			&booking.Booking.Title,
			&booking.Booking.Start,
			&booking.Booking.End,
		); err != nil {
			return nil, err
		}
		result = append(result, booking)
	}
	return result, nil
}

func (p PostgresqlRepository) AddRoomBooking(roomId int64, event models.Event) (int64, error) {
	check := p.db.QueryRow("SELECT COUNT(*) FROM calendar " +
		"WHERE roomId = $1 AND (start < $2 AND $2 < \"end\" OR \"end\" < $3 AND $3 < start )",
		roomId, event.Start, event.End)
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

	row := p.db.QueryRow("INSERT INTO calendar(roomId, title, start, \"end\", author) " +
		"VALUES ($1, $2, $3, $4, $5) RETURNING id", roomId, event.Title, event.Start, event.End, event.Author)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func (p PostgresqlRepository) GetAllRooms() ([]models.Room, error) {
	rows, err := p.db.Query("SELECT id, title " +
		"FROM rooms ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := make([]models.Room, 0)

	var roomInfo models.Room
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

func (p PostgresqlRepository) GetRoomCalendarById(roomId int64) ([]models.Event, error) {
	rows, err := p.db.Query("SELECT id, title, start, \"end\" "+
		"FROM calendar "+
		"WHERE roomId = $1", roomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var event models.Event
	calendar := make([]models.Event, 0)
	for rows.Next() {
		if err := rows.Scan(
			&event.Id,
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

func (p PostgresqlRepository) GetRoomScheduleByIdAndDate(roomId int64, date time.Time) (*models.Schedule, error) {
	panic("implement me")
}

func (p PostgresqlRepository) UpdateRoomBookingById(roomId int64, event models.Event) error {
	check := p.db.QueryRow("SELECT COUNT(*) FROM calendar " +
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

	row := p.db.QueryRow("UPDATE calendar " +
		"SET start = $1, \"end\" = $2 " +
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
