package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	ReplyID   int       `json:"reply_id,omitempty"`
	AuthorID  int       `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateCommentRequest struct {
	PostID   int    `json:"post_id"`
	ReplyID  int    `json:"reply_id,omitempty"`
	AuthorID int    `json:"-"`
	Content  string `json:"content"`
}
