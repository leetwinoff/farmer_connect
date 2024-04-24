package repo

type ConsumerRepo interface {
}

type Consumer struct {
	ID            string    `db:"id"`
	TelegramID    *int64    `db:"telegram_id"`
	NickName      *string   `db:"nickname"`
	ProfilePic    *string   `db:"profile_pic"`
	Description   *string   `db:"description"`
	Location      string    `db:"location"`
	SavedFarmers  []Farmer  `db:"saved_farmers"`
	Reviews       []Review  `db:"reviews"`
	SavedProducts []Product `db:"saved_products"`
}
