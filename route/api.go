package route

import (
	ctrget "mweibo2/controller/get"

	"github.com/gin-gonic/gin"
)

func initAPI(e *gin.Engine) {
	e.GET("/", ctrget.Home)
}
