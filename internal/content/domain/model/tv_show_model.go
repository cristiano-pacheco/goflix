package model

import "time"

type TvShowModel struct {
	id        uint64
	content   ContentModel
	createdAt time.Time
	updatedAt time.Time
}

func (ts TvShowModel) ID() uint64 {
	return ts.id
}

func (ts TvShowModel) Content() ContentModel {
	return ts.content
}

func (ts TvShowModel) CreatedAt() time.Time {
	return ts.createdAt
}

func (ts TvShowModel) UpdatedAt() time.Time {
	return ts.updatedAt
}
