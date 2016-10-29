package msg

type VideoMsg struct {
	MsgHaeder
	MediaId     string `xml:"MediaId"`
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
}

func NewVideo(id, t, d string) *VideoMsg {
	v := &VideoMsg{MediaId: id, Title: t, Description: d}
	return v
}
