package msg

type VoiceMsg struct {
	MsgHaeder
	MediaId string `xml:"MediaId"`
}

func NewVoice(id string) *VoiceMsg {
	v := &VoiceMsg{MediaId: id}
	return v
}
