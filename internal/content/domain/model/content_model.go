package model

import (
	"time"

	"github.com/cristiano-pacheco/goflix/internal/content/domain/enum"
)

type ContentModel struct {
	id                uint64
	contentType       enum.ContentTypeEnum
	title             string
	description       string
	ageRecommendation uint8
	releaseDate       time.Time
}

func (c ContentModel) ID() uint64 {
	return c.id
}

func (c ContentModel) ContentType() enum.ContentTypeEnum {
	return c.contentType
}

func (c ContentModel) Title() string {
	return c.title
}

func (c ContentModel) Description() string {
	return c.description
}

func (c ContentModel) AgeRecommendation() uint8 {
	return c.ageRecommendation
}

func (c ContentModel) ReleaseDate() time.Time {
	return c.releaseDate
}
