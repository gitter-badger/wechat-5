package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Sign in",
	})
}

func LoginHandler(c *gin.Context) {
	c.Request.ParseForm()

	c.JSON(http.StatusOK, c.Request.Form)
}
