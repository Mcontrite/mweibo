package router

import (
	ctr "mweibo/controller"

	"github.com/gin-gonic/gin"
)

func registerApis(g *gin.Engine) {
	g.GET("/", ctr.Home)
	g.GET("/captcha", GetCaptcha)

	g.GET("/register", ctr.RegisterGet)
	g.POST("/register", ctr.RegisterPost)
	g.GET("/login", ctr.LoginGet)
	g.POST("/login", ctr.LoginPost)
	g.GET("/logout", ctr.Logout)

	g.GET("weibo/:id", ctr.DisplayWeibo)
	g.GET("tag/:tag", ctr.DisplayTag)

	auth := g.Group("/auth")
	auth.Use(Auth())
	{
		auth.PUT("/user", ctr.UpdateUserAvatar)
		auth.POST("/bindemail", ctr.BindUserEmail)
		auth.POST("/unbindemail", ctr.UnbindUserEmail)

		auth.GET("/weibo", ctr.CreateWeiboGet)
		auth.POST("/weibo", ctr.CreateWeiboPost)
		auth.GET("/weibo/:id", ctr.UpdateWeiboGet)
		auth.POST("/weibo/:id", ctr.UpdateWeibo)

		auth.POST("/comment", ctr.CreateComment)
		auth.POST("/comment/:id", ctr.ReadComment)
		auth.POST("/comments", ctr.ReadAllComments)

		auth.POST("/tag", ctr.CreateTag)
	}
	admin := g.Group("/admin")
	admin.Use(Admin())
	{
		admin.GET("/users", ctr.ListUsers)
		admin.GET("/weibos", ctr.ListWeibos)

		admin.DELETE("/weibo/:id", ctr.DeleteWeibo)
		admin.DELETE("/comment/:id", ctr.DeleteComment)
	}
}
