package model

import "time"

type TvShowModel struct {
	id        uint64
	content   ContentModel
	createdAt time.Time
	updatedAt time.Time
}
