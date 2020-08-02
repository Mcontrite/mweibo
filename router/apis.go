package router

import (
	ctr "mweibo/controller"

	"github.com/gin-gonic/gin"
)

func registerApis(g *gin.Engine) {
	g.GET("/", ctr.Home)
}
