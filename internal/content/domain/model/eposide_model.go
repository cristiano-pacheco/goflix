package model

type EpisodeModel struct {
	id          uint64
	title       string
	description string
	season      uint8
	number      uint8
	duration    uint64
	video       VideoModel
	thumbnail   ThumbnailModel
}
