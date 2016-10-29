package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nicosoft.org/wechat/wechat"
)

func WeChatInit(c *gin.Context) {

	data, signature, echoStr := wechat.CheckData(c)

	if wechat.Check(data, signature) {
		c.String(http.StatusOK, "%v", echoStr)
	} else {
		at := new(wechat.AccessToken).Get()

		c.JSON(http.StatusOK, at)
	}
}

func WeChatService(c *gin.Context) {
	wechat.HandlerMsg(c)
}
