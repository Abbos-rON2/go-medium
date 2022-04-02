package repo

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
	"github.com/jackc/pgx/v4"
)

type messageRepo struct {
	db *pgx.Conn
}
type MessageI interface {
	CreateMessage(ctx context.Context, message models.CreateMessageRequest) error
	GetMessage(ctx context.Context, id string, message *models.Message) error
	GetAllMessages(ctx context.Context, messages *[]models.Message) error
	GetChatMessages(ctx context.Context, chat_id string, messages *[]models.Message) error
}

func NewMessageRepo(db *pgx.Conn) MessageI {
	return &messageRepo{db: db}
}

func (r *messageRepo) CreateMessage(ctx context.Context, message models.CreateMessageRequest) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO messages (chat_id, user_id, content) VALUES ($1, $2, $3)`,
		message.ChatID,
		message.UserID,
		message.Content,
	)
	return err
}

func (r *messageRepo) GetMessage(ctx context.Context, id string, message *models.Message) error {
	return r.db.QueryRow(
		ctx,
		`SELECT id, chat_id, user_id, content, created_at FROM messages WHERE id = $1`,
		id,
	).Scan(
		&message.ID,
		&message.ChatID,
		&message.UserID,
		&message.Content,
		&message.CreatedAt,
	)
}

func (r *messageRepo) GetAllMessages(ctx context.Context, messages *[]models.Message) error {
	rows, err := r.db.Query(
		ctx,
		`SELECT id, chat_id, user_id, content, created_at FROM messages`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Message
		if err := rows.Scan(
			&message.ID,
			&message.ChatID,
			&message.UserID,
			&message.Content,
			&message.CreatedAt,
		); err != nil {
			return err
		}
		*messages = append(*messages, message)
	}
	return nil
}

func (r *messageRepo) GetChatMessages(ctx context.Context, chatID string, messages *[]models.Message) error {
	rows, err := r.db.Query(
		ctx,
		`SELECT id, chat_id, user_id, content, created_at FROM messages WHERE chat_id = $1`,
		chatID,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Message
		if err := rows.Scan(
			&message.ID,
			&message.ChatID,
			&message.UserID,
			&message.Content,
			&message.CreatedAt,
		); err != nil {
			return err
		}
		*messages = append(*messages, message)
	}
	return nil
}
