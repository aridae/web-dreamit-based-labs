package comment_repo

import (
	db "github.com/aridae/web-dreamit-api-based-labs/internal/database"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
)

type PostgresqlRepository struct {
	db *db.PostgresClient
}

func (r *PostgresqlRepository) CreateComment(comment domain.PostComment) (int64, error) {
	queryRow := "INSERT INTO comments(notifyId, authorId, message) " +
		"VALUES ($1, $2, $3) RETURNING id"

	row := r.db.Client.QueryRow(queryRow, comment.NotifyId, comment.AuthorId, comment.Message)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func (r *PostgresqlRepository) DeleteComment(commId int64) error {
	row := r.db.Client.QueryRow(
		"DELETE FROM comments "+
			"WHERE id = $1", commId,
	)
	if err := row.Err(); err != nil {
		return err
	}
	return nil
}

func (r *PostgresqlRepository) GetComment(commId int64) (*domain.Comment, error) {
	row := r.db.Client.QueryRow(
		"SELECT id, notifyId, authorid, message FROM comments "+
			"WHERE id = $1", commId,
	)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var comment domain.Comment
	row.Scan(
		&comment.Id,
		&comment.AuthorId,
		&comment.Message,
	)
	return &comment, nil
}

func (r *PostgresqlRepository) GetNotifyComments(notifyId int64) ([]domain.Comment, error) {
	rows, err := r.db.Client.Query(
		"SELECT id, notifyId, authorId, message "+
			"FROM comments "+
			"WHERE notifyId = $1 ",
		notifyId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]domain.Comment, 0)

	commentDB := Comment{}
	for rows.Next() {
		if err := rows.Scan(
			&commentDB.Id,
			&commentDB.NotifyId,
			&commentDB.AuthorId,
			&commentDB.Message,
		); err != nil {
			return nil, err
		}
		comments = append(comments, domain.Comment{
			Id:       commentDB.Id,
			NotifyId: commentDB.NotifyId,
			AuthorId: commentDB.AuthorId,
			Message:  commentDB.Message,
		})
	}
	return comments, nil
}

func (r *PostgresqlRepository) GetAuthorComments(authorId uint64) ([]domain.Comment, error) {
	rows, err := r.db.Client.Query(
		"SELECT id, notifyId, authorId, message "+
			"FROM comments "+
			"WHERE authorId = $1 ",
		authorId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]domain.Comment, 0)

	commentDB := Comment{}
	for rows.Next() {
		if err := rows.Scan(
			&commentDB.Id,
			&commentDB.NotifyId,
			&commentDB.AuthorId,
			&commentDB.Message,
		); err != nil {
			return nil, err
		}
		comments = append(comments, domain.Comment{
			Id:       commentDB.Id,
			NotifyId: commentDB.NotifyId,
			AuthorId: commentDB.AuthorId,
			Message:  commentDB.Message,
		})
	}
	return comments, nil
}

func (r *PostgresqlRepository) DeleteNotifyComments(notifyId int64) error {
	row := r.db.Client.QueryRow(
		"DELETE FROM comments "+
			"WHERE notifyId = $1", notifyId,
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
