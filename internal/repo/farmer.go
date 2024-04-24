package repo

import (
	"context"
	"fmt"
	"time"
)

type FarmersRepo interface {
}

type Farmer struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	FullName    *string   `db:"full_name"`
	CompanyName *string   `db:"company_name"`
	ProfilePic  *string   `db:"profile_pic"`
	Description *string   `db:"description"`
	Products    []Product `db:"products"`
	Rating      int       `db:"rating"`
	Reviews     []Review  `db:"reviews"`
	Location    string    `db:"location"`
	CreatedAt   time.Time `db:"created_at"`
}

type farmersRepo struct {
	db DB
}

func NewFarmersRepo(db DB) *farmersRepo {
	return &farmersRepo{db: db}
}

func (f *farmersRepo) Create(ctx context.Context, userID, fullName string) (*Farmer, error) {
	farmer := &Farmer{}
	err := f.db.GetContext(ctx, farmer, `
		insert into farmers (user_id, full_name) values ($1, $2) returning *
	`, userID, fullName)
	if err != nil {
		return nil, fmt.Errorf("failed to exec: %w", err)
	}
	return farmer, nil
}
