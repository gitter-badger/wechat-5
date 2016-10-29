package wechat

import (
	"encoding/xml"
	"github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"nicosoft.org/wechat/wechat/msg"
	"time"
)

func HandlerMsg(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log4go.Error(err)
	}

	log4go.Debug(string(body))

	msgData := new(msg.ReceiveMsg)
	xml.Unmarshal(body, msgData)

	switch {

	case msgData.MsgType == msg.MsgTypeText:
		go PushTextMsg(c, msgData, "Test Message.")
	case msgData.MsgType == msg.MsgTypeImage:
		//TODO Image message handle.
	case msgData.MsgType == msg.MsgTypeVoice:
		//TODO Voice message handle.
	case msgData.MsgType == msg.MsgTypeVideo:
		//TODO Video message handle.
	case msgData.MsgType == msg.MsgTypeShortVideo:
		//TODO ShortVideo message handle.
	case msgData.MsgType == msg.MsgTypeLocation:
		//TODO Location message handle.
	case msgData.MsgType == msg.MsgTypeLink:
		//TODO Link message handle.
	default:
		//TODO Default message handle.
	}

}

func PushTextMsg(c *gin.Context, msgData *msg.ReceiveMsg, content string) {

	text := new(msg.TextMsg)
	text.MsgType = msg.MsgTypeText
	text.Content = content
	text.CreateTime = time.Now().Unix()
	text.FromUserName = msgData.FromUserName
	text.ToUserName = msgData.ToUserName

	xmlByte, err := xml.Marshal(text)

	if err != nil {
		log4go.Error(err)
	}

	log4go.Debug(string(xmlByte))

	c.XML(http.StatusOK, string(xmlByte))
}
