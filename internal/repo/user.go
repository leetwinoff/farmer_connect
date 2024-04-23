package repo

import (
	"context"
	"fmt"
	"time"
)

//Basic User Created when start Telegram Bot
//User can have Farmer or Consumer profile later

type UsersRepo interface {
	Create(ctx context.Context, telegramID *int64) (*User, error)
	GetByTelegramID(ctx context.Context, telegramID int64) (*User, error)
}

type User struct {
	ID            string     `db:"id"`
	OAuthProvider *string    `db:"oauth_provider"`
	Username      *string    `db:"username"`
	TelegramID    *int64     `db:"telegram_id"`
	Role          UserRole   `db:"role"`
	CreatedAt     *time.Time `db:"created_at"`
}

type UserRole int

const (
	Farmer UserRole = iota + 1
	Consumer
)

type usersRepo struct {
	db DB
}

func NewUsersRepo(db DB) *usersRepo {
	return &usersRepo{db: db}
}

func (u *usersRepo) Create(ctx context.Context, telegramID int64) (*User, error) {
	user := &User{}
	err := u.db.GetContext(ctx, user, `
		insert into users (telegram_id) values ($1) returning *
`, telegramID)
	if err != nil {
		return nil, fmt.Errorf("failed to exec: %w", err)
	}
	return user, nil
}

func (u *usersRepo) GetByTelegramID(ctx context.Context, telegramID int64) (*User, error) {
	user := &User{}
	err := u.db.GetContext(ctx, user, `
		select * from users where telegram_id = $1
`, telegramID)
	if err != nil {
		return nil, fmt.Errorf("failed to exec: %w", err)
	}
	return user, nil
}

func (u *usersRepo) UpdateRoleByTelegramID(ctx context.Context, telegramID int64, role UserRole) error {
	_, err := u.db.ExecContext(ctx, `
		update users set role = $1 where telegram_id = $2
`, role, telegramID)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}
	return nil
}
