package controller

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
	"net/http"
	"nicosoft.org/wechat/wechat"
)

func MenuCreate(c *gin.Context) {

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
