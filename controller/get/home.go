package get

import (
	"mweibo2/model"
	"mweibo2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	weibos, _ := model.ListWeibosObject()
	islogin := service.IsLogin(c)
	usersession := service.GetUserSession(c)
	user, _ := model.GetUserObjectByID(usersession.UserID)
	c.HTML(http.StatusOK, "home/home.html", gin.H{
		"islogin":     islogin,
		"usersession": usersession,
		"user":        user,
		"weibos":      weibos,
	})
}

func ErrorPage(c *gin.Context) {
	islogin := service.IsLogin(c)
	usersession := service.GetUserSession(c)
	user, _ := model.GetUserObjectByID(usersession.UserID)
	c.HTML(http.StatusNotFound, "home/error.html", gin.H{
		"islogin":     islogin,
		"usersession": usersession,
		"user":        user,
	})
}
