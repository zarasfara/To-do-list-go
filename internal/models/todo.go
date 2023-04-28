package models

import "time"

type Task struct {
	id        uint64
	title     string
	category  string
	completed bool
	createdAt time.Time
}
