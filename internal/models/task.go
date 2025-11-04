package models

import "time"

type Task struct {
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}
