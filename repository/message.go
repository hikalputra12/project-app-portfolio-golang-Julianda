package repository

import (
	"context"
	"project-portofolio/database"
	"project-portofolio/model"
	"time"
)

type MessageRepository struct {
	db database.PgxIface
}
type MessageRepositoryInterface interface {
	CreateMessage(message *model.Message) error
}

// constructor
func NewMessageRepository(db database.PgxIface) MessageRepository {
	return MessageRepository{
		db: db,
	}
}

func (r *MessageRepository) CreateMessage(message *model.Message) error {
	query := `INSERT INTO message (full_name, email, message, created_at)
	          VALUES ($1, $2, $3, $4) RETURNING message_id`

	now := time.Now()
	err := r.db.QueryRow(context.Background(), query,
		message.FullName,
		message.Email,
		message.Message,
		now,
	).Scan(&message.ID)

	if err != nil {
		return err
	}

	message.CreatedAt = now
	return nil
}
