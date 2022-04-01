package storage

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
)

func (s *storage) CreateLike(ctx context.Context, like models.Like) error {
	_, err := s.db.Exec(ctx, `
		INSERT INTO likes (
			post_id,
			user_id,
			created_at
		) VALUES (
			$1,
			$2,
			$3
		)
	`, like.PostID, like.UserID, like.CreatedAt)
	return err
}
func (s *storage) DeleteLike(ctx context.Context, postID, userID string) error {
	_, err := s.db.Exec(ctx, `
		DELETE FROM likes WHERE post_id = $1 AND user_id = $2
	`, postID, userID)
	return err
}
