package project

import "time"

//TODO: add table in database
type Project struct {
	Id         int       `json:"id"`
	DivisionId int       `json:"division_id"`
	Name       string    `json:"name"`
	Active     bool      `json:"active"`
	CreatAt    time.Time `json:"creat_at"`
	DeleteAt   time.Time `json:"delete_at"`
}
