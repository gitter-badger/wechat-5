package msg

type ImageMsg struct {
	MsgHaeder
	MediaId string `xml:"MediaId"`
}

func NewImage(id string) *ImageMsg {
	image := &ImageMsg{MediaId: id}
	return image
}
