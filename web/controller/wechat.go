package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nicosoft.org/wechat/web/service"
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

func WeChatOauth2(c *gin.Context) {
	redirect := service.GetOAath2Url()
	c.Redirect(http.StatusOK, redirect)
}

func OAuth2Handler(c *gin.Context) {

}
