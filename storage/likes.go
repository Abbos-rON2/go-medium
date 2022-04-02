package storage

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
)

func (s *storage) CreateLike(ctx context.Context, like models.CreateLikeRequest) error {
	_, err := s.db.Exec(ctx, `
		INSERT INTO likes (
			post_id,
			user_id,
		) VALUES (
			$1,
			$2,
		)
	`, like.PostID, like.UserID)
	return err
}
func (s *storage) DeleteLike(ctx context.Context, postID, userID string) error {
	_, err := s.db.Exec(ctx, `
		DELETE FROM likes WHERE post_id = $1 AND user_id = $2
	`, postID, userID)
	return err
}
