package get

import (
	"mweibo/model"
	userservice "mweibo/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	weiboList, _ := model.ListWeibosByTag("")
	islogin := userservice.IsLogin(c)
	sessions := userservice.GetSessions(c)
	user, _ := model.GetUserByID(sessions.Userid)
	tags, _ := model.ListTags()
	weibosNum := model.CountWeibos()
	usersNum := model.CountUsers()
	c.HTML(http.StatusOK, "home/home.html", gin.H{
		"weiboList":  weiboList,
		"islogin":    islogin,
		"sessions":   sessions,
		"user":       user,
		"tags":       tags,
		"weibos_num": weibosNum,
		"users_num":  usersNum,
	})
}

func Handle404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error/error.html", gin.H{
		"message": "404 not found...",
	})
	// c.HTML(http.StatusOK, "error/error.html", nil)
}
