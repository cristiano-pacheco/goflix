package model

type VideoModel struct {
	id       uint64
	url      string
	sizeInKb uint64
	duration uint64
	metadata VideoMetadataModel
}

func (v VideoModel) ID() uint64 {
	return v.id
}

func (v VideoModel) URL() string {
	return v.url
}

func (v VideoModel) SizeInKb() uint64 {
	return v.sizeInKb
}

func (v VideoModel) Duration() uint64 {
	return v.duration
}

func (v VideoModel) Metadata() VideoMetadataModel {
	return v.metadata
}
