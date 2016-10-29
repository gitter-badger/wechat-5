package msg

type TextMsg struct {
	MsgHaeder
	Content string `xml:"Content"`
}

func NewText(text string) *TextMsg {
	t := &TextMsg{Content: text}
	return t
}
