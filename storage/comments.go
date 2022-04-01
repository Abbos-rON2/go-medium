package storage

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
)

func (s *storage) CreateComment(ctx context.Context, comment models.Comment) error {
	_, err := s.db.Exec(ctx, `
		INSERT INTO comments (
			post_id,
			user_id,
			content,
			reply_id,
			created_at
		) VALUES (
			$1,
			$2,
			$3,
			$4, 
			$5
		)
	`, comment.PostID, comment.UserID, comment.Content, comment.ReplyID, comment.CreatedAt)
	return err
}
