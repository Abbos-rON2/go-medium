package repo

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
)

type LikeI interface {
	CreateLike(ctx context.Context, like models.Like) error
	DeleteLike(ctx context.Context, postID, userID string) error
}
