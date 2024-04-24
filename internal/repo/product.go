package repo

import "time"

type ProductRepo interface {
}

type Product struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Type        string    `db:"type"`
	Image       *string   `db:"image"`
	Description *string   `db:"description"`
	Price       float64   `db:"price"`
	Unit        string    `db:"unit"`
	Quantity    int       `db:"quantity"`
	CreatedAt   time.Time `db:"created_at"`
}
