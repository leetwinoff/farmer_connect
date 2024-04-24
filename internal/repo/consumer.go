package repo

import (
	"context"
	"fmt"
)

type ConsumerRepo interface {
	Create(ctx context.Context, userID, fullName string) (*Consumer, error)
	Update(ctx context.Context, id string, updates ConsumerUpdates) (*Consumer, error)
}

type Consumer struct {
	ID            string         `db:"id"`
	UserID        string         `db:"user_id"`
	FullName      *string        `db:"full_name"`
	ProfilePic    *string        `db:"profile_pic"`
	Description   *string        `db:"description"`
	Location      string         `db:"location"`
	SavedFarmers  []Farmer       `db:"saved_farmers"`
	Reviews       []Review       `db:"reviews"`
	SavedProducts []SavedProduct `db:"saved_products"`
}

type ConsumerUpdates struct {
	FullName    *string
	ProfilePic  *string
	Description *string
}

type SavedProduct struct {
	ID        string `db:"id"`
	ProductID string `db:"product_id"`
	FarmerID  string `db:"farmer_id"`
}

type consumerRepo struct {
	db DB
}

func NewConsumerRepo(db DB) *consumerRepo {
	return &consumerRepo{db: db}
}

func (c *consumerRepo) Create(ctx context.Context, userID, fullName string) (*Consumer, error) {
	consumer := &Consumer{}
	err := c.db.GetContext(ctx, consumer, `
		insert into consumers (user_id, full_name) values ($1, $2) returning *
	`, userID, fullName)
	if err != nil {
		return nil, fmt.Errorf("failed to exec: %w", err)
	}
	return consumer, nil
}

func (c *consumerRepo) Update(ctx context.Context, id string, updates ConsumerUpdates) (*Consumer, error) {
	consumer := &Consumer{}
	err := c.db.GetContext(ctx, consumer, `
		update consumers
			set full_name = coalesce($2, full_name),
			profile_pic = coalesce($3, profile_pic),
			description = coalesce($4, description)
		where id = $1
		returning *
	`, id, updates.FullName, updates.ProfilePic, updates.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to exec: %w", err)
	}
	return consumer, nil
}
