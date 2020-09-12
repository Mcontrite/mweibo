package get

import (
	"mweibo/model"
	userservice "mweibo/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCommentGET(c *gin.Context) {
	weiboID, _ := strconv.Atoi(c.Param("id"))
	sessions := userservice.GetSessions(c)
	islogin := userservice.IsLogin(c)
	c.HTML(http.StatusOK, "comment/create.html", gin.H{
		"sessions": sessions,
		"islogin":  islogin,
		"weibo_id": weiboID,
	})
}

func UpdateCommentGET(c *gin.Context) {
	commentid, _ := strconv.Atoi(c.Param("id"))
	comment, _ := model.GetCommentByID(commentid)
	sessions := userservice.GetSessions(c)
	islogin := userservice.IsLogin(c)
	c.HTML(http.StatusOK, "comment/update.html", gin.H{
		"sessions": sessions,
		"islogin":  islogin,
		"comment":  comment,
	})
}
