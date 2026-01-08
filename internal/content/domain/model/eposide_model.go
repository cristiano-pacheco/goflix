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

func (e EpisodeModel) ID() uint64 {
	return e.id
}

func (e EpisodeModel) Title() string {
	return e.title
}

func (e EpisodeModel) Description() string {
	return e.description
}

func (e EpisodeModel) Season() uint8 {
	return e.season
}

func (e EpisodeModel) Number() uint8 {
	return e.number
}

func (e EpisodeModel) Duration() uint64 {
	return e.duration
}

func (e EpisodeModel) Video() VideoModel {
	return e.video
}

func (e EpisodeModel) Thumbnail() ThumbnailModel {
	return e.thumbnail
}
