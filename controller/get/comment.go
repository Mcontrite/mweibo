package get

import (
	"mweibo2/model"
	"mweibo2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCommentGET(c *gin.Context) {
	weiboid, _ := strconv.Atoi(c.Param("id"))
	islogin := service.IsLogin(c)
	usersession := service.GetUserSession(c)
	user, _ := model.GetUserObjectByID(usersession.UserID)
	c.HTML(http.StatusOK, "comment/create.html", gin.H{
		"islogin":     islogin,
		"usersession": usersession,
		"user":        user,
		"weiboid":     weiboid,
	})
}

func UpdateCommentGET(c *gin.Context) {
	commentid, _ := strconv.Atoi(c.Param("id"))
	comment, _ := model.GetCommentByID(commentid)
	islogin := service.IsLogin(c)
	usersession := service.GetUserSession(c)
	user, _ := model.GetUserObjectByID(usersession.UserID)
	c.HTML(http.StatusOK, "comment/update.html", gin.H{
		"islogin":     islogin,
		"usersession": usersession,
		"user":        user,
		"comment":     comment,
	})
}
