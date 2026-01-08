package model

import "time"

type SeasonModel struct {
	id           uint64
	seasonNumber uint8
	title        string
	createdAt    time.Time
	updatedAt    time.Time
}

func (s SeasonModel) ID() uint64 {
	return s.id
}

func (s SeasonModel) SeasonNumber() uint8 {
	return s.seasonNumber
}

func (s SeasonModel) Title() string {
	return s.title
}

func (s SeasonModel) CreatedAt() time.Time {
	return s.createdAt
}

func (s SeasonModel) UpdatedAt() time.Time {
	return s.updatedAt
}
