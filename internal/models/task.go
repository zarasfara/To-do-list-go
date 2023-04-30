package models

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
}
