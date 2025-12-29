package model

type VideoModel struct {
	id       uint64
	url      string
	sizeInKb uint64
	duration uint64
	metadata VideoMetadataModel
}
