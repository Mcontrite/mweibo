package get

import (
	userservice "mweibo/service/user"
	"mweibo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterGET(c *gin.Context) {
	islogin := userservice.IsLogin(c)
	sessions := userservice.GetSessions(c)
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title":    "用户注册",
		"islogin":  islogin,
		"sessions": sessions,
	})
}

func LoginGET(c *gin.Context) {
	islogin := userservice.IsLogin(c)
	sessions := userservice.GetSessions(c)
	c.HTML(http.StatusOK, "user/login.html", gin.H{
		"title":    "Home Page",
		"islogin":  islogin,
		"sessions": sessions,
	})
}

func LogoutGET(c *gin.Context) {
	code := utils.SUCCESS
	userservice.LogoutSession(c)
	utils.ResponseJSONError(c, code)
}
