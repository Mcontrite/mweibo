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

	e.GET("/weibo/:id", ctrget.ReadWeibo)

	auth := e.Group("/auth")
	{
		auth.GET("/create", ctrget.CreateWeiboGET)
		auth.POST("/create", ctrpost.CreateWeiboPOST)
		auth.GET("/update/:id", ctrget.UpdateWeiboGET)
		auth.POST("/update/:id", ctrpost.UpdateWeiboPOST)

		auth.GET("/comment", ctrget.CreateCommentGET)
		auth.POST("/comment", ctrpost.CreateCommentPOST)
		auth.GET("/comment/:id", ctrget.UpdateCommentGET)
		auth.POST("/comment/:id", ctrpost.UpdateCommentPOST)
	}
}
