package question

import "time"

type Question struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	ProjectId int       `json:"project_id"`
	CreatAt   time.Time `json:"creat_at"`
	Active    bool      `json:"active"`
}
