package storage

import (
	"github.com/abbos-ron2/go-medium/config"
	"github.com/abbos-ron2/go-medium/storage/repo"
	"github.com/jackc/pgx/v4"
)

type storage struct {
	cfg         config.Config
	db          *pgx.Conn
	userRepo    repo.UserI
	postRepo    repo.PostI
	commentRepo repo.CommentI
	likeRepo    repo.LikeI
	chatRepo    repo.ChatI
	messageRepo repo.MessageI
}

type StorageI interface {
	User() repo.UserI
	Post() repo.PostI
	Comment() repo.CommentI
	Like() repo.LikeI
	Chat() repo.ChatI
	Message() repo.MessageI
}

func New(cfg config.Config, db *pgx.Conn) *storage {
	return &storage{
		cfg:         cfg,
		db:          db,
		userRepo:    repo.NewUserRepo(db),
		postRepo:    repo.NewPostRepo(db),
		commentRepo: repo.NewCommentRepo(db),
		likeRepo:    repo.NewLikeRepo(db),
		chatRepo:    repo.NewChatRepo(db),
		messageRepo: repo.NewMessageRepo(db),
	}
}

func (s *storage) User() repo.UserI {
	return s.userRepo
}
func (s *storage) Post() repo.PostI {
	return s.postRepo
}
func (s *storage) Comment() repo.CommentI {
	return s.commentRepo
}
func (s *storage) Like() repo.LikeI {
	return s.likeRepo
}
func (s *storage) Chat() repo.ChatI {
	return s.chatRepo
}
func (s *storage) Message() repo.MessageI {
	return s.messageRepo
}
