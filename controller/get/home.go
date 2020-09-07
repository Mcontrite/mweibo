package get

import (
	"mweibo/model"
	"mweibo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	weiboList, _ := model.ListWeibos("")
	islogin := service.IsLogin(c)
	sessions := service.GetSessions(c)
	weibosNum := model.CountWeibos()
	usersNum := model.CountUsers()
	c.HTML(http.StatusOK, "home/home.html", gin.H{
		"weiboList":  weiboList,
		"islogin":    islogin,
		"sessions":   sessions,
		"weibos_num": weibosNum,
		"users_num":  usersNum,
	})
}
