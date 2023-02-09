package user

import "time"

type User struct {
	Id       int       `json:"id"`
	TgId     string    `json:"tg_id"`
	IsAdmin  bool      `json:"is_admin"`
	CreateAt time.Time `json:"create_at"`
}
