package repo

import (
	"context"

	"github.com/abbos-ron2/go-medium/errs"
	"github.com/abbos-ron2/go-medium/models"
	"github.com/jackc/pgx/v4"
)

type likeRepo struct {
	db *pgx.Conn
}
type LikeI interface {
	CreateLike(ctx context.Context, like models.CreateLikeRequest) error
	DeleteLike(ctx context.Context, postID, userID string) error
}

func NewLikeRepo(db *pgx.Conn) LikeI {
	return &likeRepo{db: db}
}
func (r *likeRepo) CreateLike(ctx context.Context, like models.CreateLikeRequest) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO likes (
			post_id,
			user_id
		) VALUES (
			$1,
			$2
		)
	`, like.PostID, like.UserID)
	if err != nil {
		return errs.Errf(errs.ErrLikeAleadyExists, err)
	}
	return err
}
func (r *likeRepo) DeleteLike(ctx context.Context, postID, userID string) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM likes WHERE post_id = $1 AND user_id = $2
	`, postID, userID)
	return err
}
