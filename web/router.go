package web

import (
	"github.com/gin-gonic/gin"
	"nicosoft.org/wechat/web/controller"
)

type Router struct {
	Addr string
}

func initRouteList() *gin.Engine {

	router := gin.Default()

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
