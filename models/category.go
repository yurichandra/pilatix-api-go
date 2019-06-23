package models

import "time"

//Category represent category model
type Category struct {
	id        uint
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}
