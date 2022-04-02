package repo

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
)

type UserI interface {
	CreateUser(ctx context.Context, user models.CreateUserRequest) error
	GetUser(ctx context.Context, id string, user *models.User) error
	GetUserByEmail(ctx context.Context, email string, user *models.User) error
	GetAllUsers(ctx context.Context, users *[]models.User) error
}
