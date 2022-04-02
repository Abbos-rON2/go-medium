package repo

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
	"github.com/jackc/pgx/v4"
)

type chatRepo struct {
	db *pgx.Conn
}

type ChatI interface {
	CreateChat(ctx context.Context, chat models.CreateChatRequest) error
	GetChat(ctx context.Context, id string, chat *models.Chat) error
	GetAllChats(ctx context.Context, chats *[]models.Chat) error

	GetUserChats(ctx context.Context, userID string, chats *[]models.Chat) error
	GetChatUsers(ctx context.Context, chatID string, users *[]models.UserDTO) error
	AddUserToChat(ctx context.Context, chatID, userID string) error
	RemoveUserFromChat(ctx context.Context, chatID string, userID string) error
}

func NewChatRepo(db *pgx.Conn) ChatI {
	return &chatRepo{db: db}
}

func (r *chatRepo) CreateChat(ctx context.Context, chat models.CreateChatRequest) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO chats (user_id, content) VALUES ($1, $2)`,
		chat.UserID,
		chat.Content,
	)
	return err
}

func (r *chatRepo) GetChat(ctx context.Context, id string, chat *models.Chat) error {
	return r.db.QueryRow(
		ctx,
		`SELECT id, user_id, content, created_at FROM chats WHERE id = $1`,
		id,
	).Scan(
		&chat.ID,
		&chat.UserID,
		&chat.Content,
		&chat.CreatedAt,
	)
}
func (r *chatRepo) GetAllChats(ctx context.Context, chats *[]models.Chat) error {
	rows, err := r.db.Query(
		ctx,
		`SELECT id, user_id, content, created_at FROM chats`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var chat models.Chat
		if err := rows.Scan(
			&chat.ID,
			&chat.UserID,
			&chat.Content,
			&chat.CreatedAt,
		); err != nil {
			return err
		}
		*chats = append(*chats, chat)
	}
	return nil
}
func (r *chatRepo) GetUserChats(ctx context.Context, userID string, chats *[]models.Chat) error {
	rows, err := r.db.Query(
		ctx,
		`SELECT id, user_id, content, created_at FROM chats WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var chat models.Chat
		if err := rows.Scan(
			&chat.ID,
			&chat.UserID,
			&chat.Content,
			&chat.CreatedAt,
		); err != nil {
			return err
		}
		*chats = append(*chats, chat)
	}
	return nil
}

func (r *chatRepo) AddUserToChat(ctx context.Context, chatID, userID string) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO chats_users (chat_id, user_id) VALUES ($1, $2)`,
		chatID,
		userID,
	)
	return err
}
func (r *chatRepo) RemoveUserFromChat(ctx context.Context, userID, chatID string) error {
	_, err := r.db.Exec(
		ctx,
		`DELETE FROM chats_users WHERE chat_id = $1 AND user_id = $2`,
		chatID,
		userID,
	)
	return err
}
func (r *chatRepo) GetChatUsers(ctx context.Context, chatID string, users *[]models.UserDTO) error {
	rows, err := r.db.Query(
		ctx,
		`SELECT users.id, users.name, users.email, users.created_at FROM chats_users
			INNER JOIN users ON chats_users.user_id = users.id
			WHERE chats_users.chat_id = $1`,
		chatID,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UserDTO
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return err
		}
		*users = append(*users, user)
	}
	return nil
}
