package repo

import (
	"time"
)

type ReviewRepo interface {
}

type Review struct {
	ID         string    `db:"id"`
	ConsumerID string    `db:"consumer_id"`
	FarmerID   string    `db:"farmer_id"`
	Comment    string    `db:"comment"`
	CreateAt   time.Time `db:"created_at"`
}
