package msg

type MusicMsg struct {
	MsgHaeder
	Title        string `xml:"MediaId"`
	Description  string `xml:"MediaId"`
	MusicURL     string `xml:"MediaId"`
	HQMusicUrl   string `xml:"MediaId"`
	ThumbMediaId string `xml:"MediaId"`
}

func NewMusic(title, des, music, hq, tm string) *MusicMsg {
	m := &MusicMsg{Title: title, Description: des, MusicURL: music, ThumbMediaId: tm}
	return m
}
