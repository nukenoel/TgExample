package division

import (
	"time"
)

type Division struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Active   bool      `json:"active"`
	CreatAt  time.Time `json:"creat_at"`
	UpdateAt time.Time `json:"update_at"`
	DeleteAt time.Time `json:"delete_at"`
}
