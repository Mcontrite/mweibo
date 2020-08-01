package router

import (
	"github.com/gin-gonic/gin"
)

func registerRouters(g *gin.Engine) {
	g.Get("/")
}
