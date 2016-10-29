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

	router.GET("/push", controller.WeChatInit)
	router.POST("/push", controller.WeChatService)

	router.GET("/menu", controller.WeChatMenuCreate)

	return router
}

func (r *Router) Start() {

	if r.Addr != "" {
		initRouteList().Run(r.Addr)
	} else {
		initRouteList().Run()
	}

}
