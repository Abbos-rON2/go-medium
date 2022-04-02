package storage

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
)

func (s *storage) CreateComment(ctx context.Context, comment models.CreateCommentRequest) error {
	_, err := s.db.Exec(ctx, `
		INSERT INTO comments (
			post_id,
			user_id,
			content,
			reply_id,
		) VALUES (
			$1,
			$2,
			$3,
			$4, 
			$5
		)
	`, comment.PostID, comment.UserID, comment.Content, comment.ReplyID)
	return err
}
