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
