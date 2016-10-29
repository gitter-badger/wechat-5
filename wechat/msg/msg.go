package msg

import "encoding/xml"

type MsgHaeder struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
}

const (
	MsgTypeText       = "text"
	MsgTypeImage      = "image"
	MsgTypeVoice      = "voice"
	MsgTypeVideo      = "video"
	MsgTypeShortVideo = "shortvideo"
	MsgTypeLocation   = "location"
	MsgTypeLink       = "link"
	MsgTypeMusic      = "music"
	MsgTypeNews       = "news"
	MsgTypeTransfer   = "transfer_customer_service"
	MsgTypeEvent      = "event"
)

const (
	EventSubscribe       = "subscribe"
	EventUnsubscribe     = "unsubscribe"
	EventScan            = "SCAN"
	EventLocation        = "LOCATION"
	EventClick           = "CLICK"
	EventView            = "VIEW"
	EventScancodePush    = "scancode_push"
	EventScancodeWaitmsg = "scancode_waitmsg"
	EventPicSysphoto     = "pic_sysphoto"
	EventPicPhotoOrAlbum = "pic_photo_or_album"
	EventPicWeixin       = "pic_weixin"
	EventLocationSelect  = "location_select"
)

type ReceiveMsg struct {
	MsgHaeder

	MsgID        int64   `xml:"MsgId"`
	Content      string  `xml:"Content"`
	PicURL       string  `xml:"PicUrl"`
	MediaID      string  `xml:"MediaId"`
	Format       string  `xml:"Format"`
	ThumbMediaID string  `xml:"ThumbMediaId"`
	LocationX    float64 `xml:"Location_X"`
	LocationY    float64 `xml:"Location_Y"`
	Scale        float64 `xml:"Scale"`
	Label        string  `xml:"Label"`
	Title        string  `xml:"Title"`
	Description  string  `xml:"Description"`
	URL          string  `xml:"Url"`
}
