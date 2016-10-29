package controller

import (
	"github.com/alecthomas/log4go"
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

	code := c.Param("code")
	state := c.Param("state")

	log4go.Debug("OAuth2 code: %v,state: %v", code, state)

	if code == "" {
		c.String(http.StatusOK, "%v", "Request Bad.")
		return
	}

	user := new(wechat.UserInfo)
	token := user.GetToken(code)
	result := user.GetUserInfo(token.OpenId, token.Token)

	c.JSON(http.StatusOK, result)
}
