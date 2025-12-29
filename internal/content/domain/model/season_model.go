package model

import "time"

type SeasonModel struct {
	id           uint64
	seasonNumber uint8
	title        string
	createdAt    time.Time
	updatedAt    time.Time
}
