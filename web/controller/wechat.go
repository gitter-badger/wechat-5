package controller

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
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
	c.String(http.StatusOK, "%v", "Request Bad.")
}

func WeChatMenuCreate(c *gin.Context) {

	data := c.Param("menu")

	if data == "" {
		c.String(http.StatusOK, "%v", "Request Bad.")
		return
	}

	menu := new(wechat.Menu)
	err := json.Unmarshal([]byte(data), &menu)

	if err != nil {
		log4go.Error(err)
		c.String(http.StatusInternalServerError, "%v", err.Error())
	} else {
		res := menu.Create()
		c.JSON(http.StatusOK, res)
	}

}
