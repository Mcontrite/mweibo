package route

import (
	ctrget "mweibo2/controller/get"
	ctrpost "mweibo2/controller/post"

	"github.com/gin-gonic/gin"
)

func initAPI(e *gin.Engine) {
	e.GET("/", ctrget.Home)
	e.GET("/register", ctrget.RegisterGET)
	e.POST("/register", ctrpost.RegisterPOST)
	e.GET("/login", ctrget.LoginGET)
	e.POST("/login", ctrpost.LoginPOST)
	e.GET("/logout", ctrget.LogoutGET)
}
