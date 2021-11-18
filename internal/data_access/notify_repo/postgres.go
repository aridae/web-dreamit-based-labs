package notifyrepo

import (
	"fmt"

	db "github.com/aridae/web-dreamit-api-based-labs/internal/database"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"github.com/lib/pq"
)

type PostgresqlRepository struct {
	db *db.PostgresClient
}

func (r *PostgresqlRepository) FilterNotifies(filter domain.OptionalNotifyFilter) ([]domain.Notify, error) {

	queryRow := "SELECT id from notifies JOIN notify_tags on notifies.id = notify_tags.notifyId WHERE notifyId = $1"
	if filter.Subject != "" {
		queryRow += " AND subject = $2"
	}
	if len(filter.Tags) > 0 {
		queryRow += " AND tag = ANY($3)"
	}

	rows, err := r.db.Client.Query(
		queryRow,
		filter.EventId,
		filter.Subject,
		pq.Array(filter.Tags),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifies := make([]domain.Notify, 0)

	notifyDB := Notify{}
	for rows.Next() {
		if err := rows.Scan(
			&notifyDB.Id,
			&notifyDB.EventId,
			&notifyDB.Subject,
			&notifyDB.Message,
		); err != nil {
			return nil, err
		}

		tags, err := r.GetTagsForNotify(notifyDB.Id)
		if err != nil {
			return nil, err
		}
		notifyDB.Tags = tags

		notifies = append(notifies, domain.Notify{
			Id:      notifyDB.Id,
			EventId: notifyDB.EventId,
			Subject: notifyDB.Subject,
			Tags:    notifyDB.Tags,
			Message: notifyDB.Message,
		})
	}
	return notifies, nil
}

func (r *PostgresqlRepository) GetTagsForNotify(notifyId int64) ([]string, error) {
	rows, err := r.db.Client.Query(
		"SELECT tag FROM notify_tags "+
			"WHERE notifyId = $1", notifyId,
	)
	if err != nil {
		return nil, err
	}

	var tag string
	tags := make([]string, 0)
	for rows.Next() {
		if err := rows.Scan(
			&tag,
		); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *PostgresqlRepository) CreateNotify(notify domain.PostNotify) (int64, error) {
	queryRow := "INSERT INTO notifies(eventId, message, subject) " +
		"VALUES ($1, $2, $3) RETURNING id"
	row := r.db.Client.QueryRow(queryRow, notify.EventId, notify.Message)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	queryRow = fmt.Sprintf("INSERT INTO notify_tags(notifyId, tag) "+
		"VALUES (%d, %s), ", id, notify.Tags[0])
	for i := 1; i < len(notify.Tags); i++ {
		queryRow += fmt.Sprintf("VALUES (%d, %s), ", id, notify.Tags[i])
	}

	row = r.db.Client.QueryRow(queryRow)
	if err := row.Err(); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresqlRepository) GetNotify(notifyId int64) (*domain.Notify, error) {
	row := r.db.Client.QueryRow(
		"SELECT id, eventId, subject, message FROM notifies "+
			"WHERE id = $1", notifyId,
	)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var notify domain.Notify
	row.Scan(
		&notify.Id,
		&notify.EventId,
		&notify.Subject,
		&notify.Message,
	)
	tags, err := r.GetTagsForNotify(notifyId)
	if err != nil {
		return nil, err
	}
	notify.Tags = tags
	return &notify, nil
}

func (r *PostgresqlRepository) DeleteNotify(notifyId int64) error {
	row := r.db.Client.QueryRow(
		"DELETE FROM notifies "+
			"WHERE id = $1", notifyId,
	)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (r *PostgresqlRepository) GetNotifiesByEventId(eventId int64) ([]domain.Notify, error) {
	rows, err := r.db.Client.Query(
		"SELECT id, eventId, subject, message "+
			"FROM notifies "+
			"WHERE eventId = $1 ",
		eventId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifies := make([]domain.Notify, 0)

	notifyDB := Notify{}
	for rows.Next() {
		if err := rows.Scan(
			&notifyDB.Id,
			&notifyDB.EventId,
			&notifyDB.Subject,
			&notifyDB.Message,
		); err != nil {
			return nil, err
		}

		tags, err := r.GetTagsForNotify(notifyDB.Id)
		if err != nil {
			return nil, err
		}
		notifyDB.Tags = tags

		notifies = append(notifies, domain.Notify{
			Id:      notifyDB.Id,
			EventId: notifyDB.EventId,
			Subject: notifyDB.Subject,
			Tags:    notifyDB.Tags,
			Message: notifyDB.Message,
		})
	}
	return notifies, nil
}

func (r *PostgresqlRepository) GetEventNotifiesWithTag(eventId int64, tag string) ([]domain.Notify, error) {
	rows, err := r.db.Client.Query(
		"SELECT id, eventId, subject, message "+
			"FROM notifies join notify_tags on notifies.id = notify_tags.notifyId "+
			"WHERE eventId = $1 AND tag='$2'",
		eventId, tag,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifies := make([]domain.Notify, 0)

	notifyDB := Notify{}
	for rows.Next() {
		if err := rows.Scan(
			&notifyDB.Id,
			&notifyDB.EventId,
			&notifyDB.Subject,
			&notifyDB.Message,
		); err != nil {
			return nil, err
		}

		tags, err := r.GetTagsForNotify(notifyDB.Id)
		if err != nil {
			return nil, err
		}
		notifyDB.Tags = tags

		notifies = append(notifies, domain.Notify{
			Id:      notifyDB.Id,
			EventId: notifyDB.EventId,
			Subject: notifyDB.Subject,
			Tags:    notifyDB.Tags,
			Message: notifyDB.Message,
		})
	}
	return notifies, nil
}

func (r *PostgresqlRepository) GetEventNotifiesWithSubject(eventId int64, subject string) ([]domain.Notify, error) {
	rows, err := r.db.Client.Query(
		"SELECT id, eventId, subject, message "+
			"FROM notifies "+
			"WHERE eventId = $1 AND subject='$2'",
		eventId, subject,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifies := make([]domain.Notify, 0)

	notifyDB := Notify{}
	for rows.Next() {
		if err := rows.Scan(
			&notifyDB.Id,
			&notifyDB.EventId,
			&notifyDB.Subject,
			&notifyDB.Message,
		); err != nil {
			return nil, err
		}

		tags, err := r.GetTagsForNotify(notifyDB.Id)
		if err != nil {
			return nil, err
		}
		notifyDB.Tags = tags

		notifies = append(notifies, domain.Notify{
			Id:      notifyDB.Id,
			EventId: notifyDB.EventId,
			Subject: notifyDB.Subject,
			Tags:    notifyDB.Tags,
			Message: notifyDB.Message,
		})
	}
	return notifies, nil
}

func (r *PostgresqlRepository) GetEventNotifiesWithTagSubject(eventId int64, subject, tag string) ([]domain.Notify, error) {
	rows, err := r.db.Client.Query(
		"SELECT id, eventId, subject, message "+
			"FROM notifies join notify_tags on notifies.id = notify_tags.notifyId "+
			"WHERE eventId = $1 AND subject='$2' AND tag='$3'",
		eventId, subject, tag,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifies := make([]domain.Notify, 0)

	notifyDB := Notify{}
	for rows.Next() {
		if err := rows.Scan(
			&notifyDB.Id,
			&notifyDB.EventId,
			&notifyDB.Subject,
			&notifyDB.Message,
		); err != nil {
			return nil, err
		}

		tags, err := r.GetTagsForNotify(notifyDB.Id)
		if err != nil {
			return nil, err
		}
		notifyDB.Tags = tags

		notifies = append(notifies, domain.Notify{
			Id:      notifyDB.Id,
			EventId: notifyDB.EventId,
			Subject: notifyDB.Subject,
			Tags:    notifyDB.Tags,
			Message: notifyDB.Message,
		})
	}
	return notifies, nil
}

func NewSessionPostgresqlRepository(client *db.PostgresClient) Repository {
	return &PostgresqlRepository{
		db: client,
	}
}
