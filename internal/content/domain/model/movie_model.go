package model

import "time"

type MovieModel struct {
	id             uint64
	externalRating float64
	content        ContentModel
	video          VideoModel
	thumbnail      ThumbnailModel
	createdAt      time.Time
	updatedAt      time.Time
}

func NewMovieModel(
	externalRating float64,
	content ContentModel,
	video VideoModel,
	thumbnail ThumbnailModel,
	createdAt time.Time,
	updatedAt time.Time,
) MovieModel {
	return MovieModel{
		content:   content,
		video:     video,
		thumbnail: thumbnail,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func RestoreMovieModel(
	id uint64,
	externalRating float64,
	content ContentModel,
	video VideoModel,
	thumbnail ThumbnailModel,
	createdAt time.Time,
	updatedAt time.Time,
) MovieModel {
	return MovieModel{
		id:        id,
		content:   content,
		video:     video,
		thumbnail: thumbnail,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (m *MovieModel) ID() uint64 {
	return m.id
}

func (m *MovieModel) ExternalRating() float64 {
	return m.externalRating
}

func (m *MovieModel) Content() ContentModel {
	return m.content
}

func (m *MovieModel) Video() VideoModel {
	return m.video
}

func (m *MovieModel) Thumbnail() ThumbnailModel {
	return m.thumbnail
}

func (m *MovieModel) CreatedAt() time.Time {
	return m.createdAt
}

func (m *MovieModel) UpdatedAt() time.Time {
	return m.updatedAt
}
