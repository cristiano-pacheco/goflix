package model

type ThumbnailModel struct {
	id  uint64
	url string
}

func NewThumbnailModel(id uint64, url string) ThumbnailModel {
	return ThumbnailModel{
		id:  id,
		url: url,
	}
}

func (m *ThumbnailModel) ID() uint64 {
	return m.id
}

func (m *ThumbnailModel) URL() string {
	return m.url
}
