package inviterepo

import (
	db "github.com/aridae/web-dreamit-api-based-labs/internal/database"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

type PostgresqlRepository struct {
	db *db.PostgresClient
}

func (r *PostgresqlRepository) CreateInvite(invite domain.PostInvite) (int64, error) {
	queryRow := "INSERT INTO invites(eventId, receiverId, statusId) " +
		"VALUES ($1, $2, $3) RETURNING id"

	row := r.db.Client.QueryRow(queryRow, invite.EventId, invite.ReceiverId, domain.PENDING_STATUS)
	if err := row.Err(); err != nil {
		return -1, err
	}

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func (r *PostgresqlRepository) UpdateInviteStatusById(inviteId int64, status int64) error {
	row := r.db.Client.QueryRow(
		"UPDATE invites "+
			"SET statusId = $1 "+
			"WHERE id = $2 ",
		status, inviteId,
	)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (r *PostgresqlRepository) GetInviteById(inviteId int64) (*domain.Invite, error) {
	row := r.db.Client.QueryRow(
		"SELECT id, eventId, receiverId, statusId FROM invites "+
			"WHERE id = $1", inviteId,
	)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var invite domain.Invite
	row.Scan(
		&invite.Id,
		&invite.ReceiverId,
		&invite.StatusId,
	)
	return &invite, nil
}

func (r *PostgresqlRepository) DeleteInviteById(inviteId int64) error {
	row := r.db.Client.QueryRow(
		"DELETE FROM invites "+
			"WHERE id = $1", inviteId,
	)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (r *PostgresqlRepository) GetInvitesByEventId(eventId int64) ([]domain.Invite, error) {
	rows, err := r.db.Client.Query(
		"SELECT id, eventId, receiverId, statusId "+
			"FROM invites "+
			"WHERE eventId = $1 ",
		eventId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invites := make([]domain.Invite, 0)

	inviteDB := Invite{}
	for rows.Next() {
		if err := rows.Scan(
			&inviteDB.Id,
			&inviteDB.EventId,
			&inviteDB.ReceiverId,
			&inviteDB.StatusId,
		); err != nil {
			return nil, err
		}
		invites = append(invites, domain.Invite{
			Id:         inviteDB.Id,
			EventId:    inviteDB.EventId,
			ReceiverId: inviteDB.ReceiverId,
			StatusId:   uint64(inviteDB.StatusId),
		})
	}
	return invites, nil
}

func (r *PostgresqlRepository) GetStatusInvitesByEventId(eventId int64, statusId int64) ([]domain.Invite, error) {
	rows, err := r.db.Client.Query(
		"SELECT id, eventId, receiverId, statusId "+
			"FROM invites "+
			"WHERE eventId = $1 and statusId = $2",
		eventId, statusId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invites := make([]domain.Invite, 0)

	inviteDB := Invite{}
	for rows.Next() {
		if err := rows.Scan(
			&inviteDB.Id,
			&inviteDB.EventId,
			&inviteDB.ReceiverId,
			&inviteDB.StatusId,
		); err != nil {
			return nil, err
		}
		invites = append(invites, domain.Invite{
			Id:         inviteDB.Id,
			EventId:    inviteDB.EventId,
			ReceiverId: inviteDB.ReceiverId,
			StatusId:   uint64(inviteDB.StatusId),
		})
	}
	return invites, nil
}

func (r *PostgresqlRepository) GetInvitesByReceiverId(recId uint64) ([]domain.Invite, error) {
	rows, err := r.db.Client.Query(
		"SELECT id, eventId, receiverId, statusId "+
			"FROM invites "+
			"WHERE receiverId = $1 ",
		recId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invites := make([]domain.Invite, 0)

	inviteDB := Invite{}
	for rows.Next() {
		if err := rows.Scan(
			&inviteDB.Id,
			&inviteDB.EventId,
			&inviteDB.ReceiverId,
			&inviteDB.StatusId,
		); err != nil {
			return nil, err
		}
		invites = append(invites, domain.Invite{
			Id:         inviteDB.Id,
			EventId:    inviteDB.EventId,
			ReceiverId: inviteDB.ReceiverId,
			StatusId:   uint64(inviteDB.StatusId),
		})
	}
	return invites, nil
}

func (r *PostgresqlRepository) UpdateInvitesStatusByEventId(eventId int64, status int64) error {
	row := r.db.Client.QueryRow(
		"UPDATE invites "+
			"SET statusId = $1 "+
			"WHERE eventId = $2 ",
		status, eventId,
	)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (r *PostgresqlRepository) DeleteInvitesByEventId(eventId int64) error {
	row := r.db.Client.QueryRow(
		"DELETE FROM invites "+
			"WHERE eventId = $1", eventId,
	)
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
