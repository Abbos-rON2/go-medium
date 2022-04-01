package repo

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
)

type CommentI interface {
	CreateComment(ctx context.Context, comment models.Comment) error
}
