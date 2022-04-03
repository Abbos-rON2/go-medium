package repo

import (
	"context"

	"github.com/abbos-ron2/go-medium/models"
	"github.com/jackc/pgx/v4"
)

type subscriptionRepo struct {
	db *pgx.Conn
}
type SubscriptionI interface {
	Subscribe(ctx context.Context, userID string, targetID string) error
	Unsubscribe(ctx context.Context, userID string, targetID string) error

	GetSubscribers(ctx context.Context, userID string, subscribers *[]models.Subscription) error
	GetSubscribersCount(ctx context.Context, userID string, count *int) error

	GetSubscriptionsCount(ctx context.Context, userID string, count *int) error
	GetSubscriptions(ctx context.Context, userID string, subscriptions *[]models.Subscription) error
}

func NewSubscriptionRepo(db *pgx.Conn) SubscriptionI {
	return &subscriptionRepo{db: db}
}

func (r *subscriptionRepo) Subscribe(ctx context.Context, userID string, targetID string) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO subscriptions (user_id, subscriber_id) VALUES ($1, $2)`,
		targetID,
		userID,
	)
	return err
}

func (r *subscriptionRepo) Unsubscribe(ctx context.Context, userID string, targetID string) error {
	_, err := r.db.Exec(
		ctx,
		`DELETE FROM subscriptions WHERE user_id = $1 AND subscriber_id = $2`,
		targetID,
		userID,
	)
	return err
}

func (r *subscriptionRepo) GetSubscribers(ctx context.Context, userID string, subscribers *[]models.Subscription) error {
	rows, err := r.db.Query(
		ctx,
		`SELECT id, subscriber_id, name, email, created_at FROM subscriptions WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var subscriber models.Subscription
		if err := rows.Scan(
			&subscriber.ID,
			&subscriber.UserID,
			&subscriber.Name,
			&subscriber.Email,
			&subscriber.CreatedAt,
		); err != nil {
			return err
		}
		*subscribers = append(*subscribers, subscriber)
	}
	return rows.Err()
}

func (r *subscriptionRepo) GetSubscribersCount(ctx context.Context, userID string, count *int) error {
	err := r.db.QueryRow(
		ctx,
		`SELECT COUNT(*) FROM subscriptions WHERE user_id = $1`,
		userID,
	).Scan(
		&count,
	)
	return err
}

func (r *subscriptionRepo) GetSubscriptions(ctx context.Context, userID string, subscriptions *[]models.Subscription) error {
	rows, err := r.db.Query(
		ctx,
		`SELECT id, user_id, name, email, created_at FROM subscriptions WHERE subscriber_id = $1`,
		userID,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var subscription models.Subscription
		if err := rows.Scan(
			&subscription.ID,
			&subscription.UserID,
			&subscription.Name,
			&subscription.Email,
			&subscription.CreatedAt,
		); err != nil {
			return err
		}
		*subscriptions = append(*subscriptions, subscription)
	}
	return rows.Err()
}

func (r *subscriptionRepo) GetSubscriptionsCount(ctx context.Context, userID string, count *int) error {
	err := r.db.QueryRow(
		ctx,
		`SELECT COUNT(*) FROM subscriptions WHERE subscriber_id = $1`,
		userID,
	).Scan(
		&count,
	)
	return err
}
