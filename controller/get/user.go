package get

import (
	"mweibo/pkgs"
	"mweibo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterGET(c *gin.Context) {
	islogin := service.IsLogin(c)
	sessions := service.GetSessions(c)
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title":    "用户注册",
		"islogin":  islogin,
		"sessions": sessions,
	})
}

func LoginGET(c *gin.Context) {
	islogin := service.IsLogin(c)
	sessions := service.GetSessions(c)
	c.HTML(http.StatusOK, "user/login.html", gin.H{
		"title":    "Home Page",
		"islogin":  islogin,
		"sessions": sessions,
	})
}

func LogoutGET(c *gin.Context) {
	code := pkgs.SUCCESS
	service.LogoutSession(c)
	pkgs.ResponseJSONError(c, code)
}
