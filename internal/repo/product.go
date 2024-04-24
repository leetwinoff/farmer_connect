package repo

import "time"

type ProductRepo interface {
}

type Product struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Type        string    `db:"type"`
	ProductPic  *string   `db:"product_pic"`
	Description *string   `db:"description"`
	Price       int       `db:"price"`
	Unit        string    `db:"unit"`
	Quantity    int       `db:"quantity"`
	CreatedAt   time.Time `db:"created_at"`
}
