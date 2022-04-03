package repo

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
	"github.com/jackc/pgx/v4"
)

type userRepo struct {
	db *pgx.Conn
}
type UserI interface {
	CreateUser(ctx context.Context, user models.CreateUserRequest) error
	GetUser(ctx context.Context, id string, user *models.User) error
	GetUserByEmail(ctx context.Context, email string, user *models.User) error
	GetAllUsers(ctx context.Context, users *[]models.User) error
}

func NewUserRepo(db *pgx.Conn) UserI {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(ctx context.Context, user models.CreateUserRequest) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`,
		user.Name,
		user.Email,
		user.Password,
	)
	return err
}

func (r *userRepo) GetUser(ctx context.Context, id string, user *models.User) error {
	return r.db.QueryRow(
		ctx,
		`SELECT id, name, email, password, created_at FROM users WHERE id = $1`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string, user *models.User) error {
	return r.db.QueryRow(
		ctx,
		`SELECT id, name, email, password, created_at FROM users WHERE email = $1`,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
}

func (r *userRepo) GetAllUsers(ctx context.Context, users *[]models.User) error {
	rows, err := r.db.Query(
		ctx,
		`SELECT id, name, email, password, created_at FROM users`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		); err != nil {
			return err
		}
		*users = append(*users, user)
	}
	return rows.Err()
}
