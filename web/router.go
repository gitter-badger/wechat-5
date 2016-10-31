package web

import (
	"github.com/gin-gonic/gin"
	"nicosoft.org/wechat/web/controller"
	"os"
)

type Router struct {
	Addr string
}

func initRouteList() *gin.Engine {

	viewpath, _ := os.Getwd()

	router := gin.Default()
	router.LoadHTMLGlob(viewpath + "/web/views/*")

	router.GET("/login", controller.LoginPage)
	router.POST("/login", controller.LoginHandler)

	router.GET("/core", controller.WeChatInit)
	router.POST("/core", controller.WeChatService)

	router.GET("/oauth2/init", controller.WeChatOauth2)
	router.GET("/oauth2/handler", controller.OAuth2Handler)

	router.POST("/manager/menu", controller.MenuCreate)

	return router
}

func (r *Router) Start() {

	if r.Addr != "" {
		initRouteList().Run(r.Addr)
	} else {
		initRouteList().Run()
	}

}
