package router

import (
	ctr "mweibo/controller"

	"github.com/gin-gonic/gin"
)

func registerApis(g *gin.Engine) {
	g.GET("/", ctr.Home)
	g.GET("/register", ctr.RegisterGet)
	g.POST("/register", ctr.RegisterPost)
	g.GET("/login", ctr.LoginGet)
	g.POST("/login", ctr.LoginPost)
	g.GET("/createweibo", ctr.CreateWeiboGet)
	g.POST("/createweibo", ctr.CreateWeiboPost)
}
