package storage

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
)

func (s *storage) CreateUser(ctx context.Context, user models.User) error {
	_, err := s.db.Exec(
		ctx,
		`INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`,
		user.Name,
		user.Email,
		user.Password,
	)
	return err
}
func (s *storage) GetUser(ctx context.Context, id string, user *models.User) error {
	return s.db.QueryRow(
		ctx,
		`SELECT id, name, email, password FROM users WHERE id = $1`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)
}
func (s *storage) GetUserByEmail(ctx context.Context, email string, user *models.User) error {
	return s.db.QueryRow(
		ctx,
		`SELECT id, name, email, password FROM users WHERE email = $1`,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)
}
func (s *storage) GetAllUsers(ctx context.Context, users *[]models.User) error {
	rows, err := s.db.Query(
		ctx,
		`SELECT id, name, email, password FROM users`,
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
		); err != nil {
			return err
		}
		*users = append(*users, user)
	}
	return rows.Err()
}
