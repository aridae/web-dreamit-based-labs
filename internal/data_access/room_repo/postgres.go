package roomrepo

import (
	db "github.com/aridae/web-dreamit-api-based-labs/internal/database"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

type PostgresqlRepository struct {
	db *db.PostgresClient
}

func (p PostgresqlRepository) GetRoom(roomId int64) (*domain.Room, error) {
	row := p.db.Client.QueryRow("SELECT id, title FROM rooms WHERE id = $1", roomId)
	if row.Err() != nil {
		return nil, row.Err()
	}

	room := new(domain.Room)
	if err := row.Scan(
		&room.Id,
		&room.Title,
	); err != nil {
		return nil, err
	}

	return room, nil
}

func (p PostgresqlRepository) GetAllRooms() ([]domain.Room, error) {
	rows, err := p.db.Client.Query("SELECT id, title " +
		"FROM rooms ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := make([]domain.Room, 0)

	var roomInfo Room
	for rows.Next() {
		if err := rows.Scan(
			&roomInfo.Id,
			&roomInfo.Title,
		); err != nil {
			return nil, err
		}
		rooms = append(rooms, domain.Room{
			Id:    roomInfo.Id,
			Title: roomInfo.Title,
		})
	}
	return rooms, nil
}

func NewSessionPostgresqlRepository(client *db.PostgresClient) Repository {
	return &PostgresqlRepository{
		db: client,
	}
}
