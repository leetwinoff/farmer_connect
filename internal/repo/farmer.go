package repo

import "time"

type FarmersRepo interface {
}

type Farmer struct {
	ID          string    `db:"id"`
	TelegramID  *int64    `db:"telegram_id"`
	NickName    *string   `db:"nickname"`
	ProfilePic  *string   `db:"profile_pic"`
	Description *string   `db:"description"`
	Products    []Product `db:"products"`
	Rating      int       `db:"rating"`
	Reviews     []Review  `db:"reviews"`
	Location    string    `db:"location"`
	CreatedAt   time.Time `db:"created_at"`
}
