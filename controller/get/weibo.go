package get

import (
	"fmt"
	"mweibo2/model"
	"mweibo2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadWeibo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	weibo, err := model.GetWeiboObjectByID(id)
	if err != nil {
		fmt.Println("Get weibo err: ", err)
		return
	}
	weibo.ViewsCnt++
	model.UpdateWeiboViewsCnt(&weibo)
	islogin := service.IsLogin(c)
	usersession := service.GetUserSession(c)
	comments, _ := model.ListCommentsByWeiboID(id, 10)
	c.HTML(http.StatusOK, "weibo/read.html", gin.H{
		"islogin":     islogin,
		"usersession": usersession,
		"weibo":       weibo,
		"comments":    comments,
	})
}

func CreateWeiboGET(c *gin.Context) {
	islogin := service.IsLogin(c)
	usersession := service.GetUserSession(c)
	c.HTML(http.StatusOK, "weibo/create.html", gin.H{
		"islogin":     islogin,
		"usersession": usersession,
	})
}

func UpdateWeiboGET(c *gin.Context) {
	weiboid, _ := strconv.Atoi(c.Param("id"))
	weibo, _ := model.GetWeiboObjectByID(weiboid)
	islogin := service.IsLogin(c)
	usersession := service.GetUserSession(c)
	c.HTML(http.StatusOK, "weibo/update.html", gin.H{
		"islogin":     islogin,
		"usersession": usersession,
		"weibo":       weibo,
	})
}
