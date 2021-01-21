package get

import (
	"mweibo2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func writeJSON(c *gin.Context, h gin.H) {
	if _, ok := h["success"]; !ok {
		h["success"] = false
	}
	c.JSON(http.StatusOK, h)
}

func RegisterGET(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register.html", nil)
}

func LoginGET(c *gin.Context) {
	islogin := service.IsLogin(c)
	usersession := service.GetUserSession(c)
	c.HTML(http.StatusOK, "user/login.html", gin.H{
		"islogin":     islogin,
		"usersession": usersession,
	})
}

func LogoutGET(c *gin.Context) {
	service.LogoutSession(c)
	// c.JSON(http.StatusOK,nil)
	c.Redirect(http.StatusSeeOther, "/")
}
