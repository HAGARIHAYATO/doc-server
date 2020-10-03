package model

import "time"

type Doc struct {
	ID int64 `db:"id"`
	Title string `db:"title"`
	Text string `db:"text"`
	UserID int64 `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
