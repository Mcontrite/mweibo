package get

import (
	"mweibo/model"
	userservice "mweibo/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateWeiboGET(c *gin.Context) {
	islogin := userservice.IsLogin(c)
	sessions := userservice.GetSessions(c)
	c.HTML(http.StatusOK, "weibo/create.html", gin.H{
		"islogin":  islogin,
		"sessions": sessions,
	})
}

func DisplayWeibo(c *gin.Context) {
	weiboID, _ := strconv.Atoi(c.Param("id"))
	weibo, err := model.GetWeiboByID(weiboID)
	if err != nil {
		c.HTML(http.StatusNotFound, "error/error.html", gin.H{})
		return
	}
	weibo.ViewsCnt++
	model.UpdateWeiboViewsCnt(&weibo)
	islogin := userservice.IsLogin(c)
	sessions := userservice.GetSessions(c)
	commentlist, _ := model.ListCommentsByWeiboID(weiboID, 500)
	c.HTML(http.StatusOK, "weibo/display.html", gin.H{
		"weibo": weibo,
		// "fcomment":         fcomment,
		"islogin":     islogin,
		"sessions":    sessions,
		"commentlist": commentlist,
		// "comment_list_len": commentlistLen,
		// "attachs":        attachs,
		// "isfav":          isfav,
	})
}

func UpdateWeiboGET(c *gin.Context) {
	weiboID, _ := strconv.Atoi(c.Param("id"))
	weibo, _ := model.GetWeiboByID(weiboID)
	// fcomment, _ := model.GetWeiboFirstCommentByTid(weiboID)
	islogin := userservice.IsLogin(c)
	sessions := userservice.GetSessions(c)
	// attachs, _ := model.GetAttachsByCommentID(int(fcomment.ID))
	c.HTML(http.StatusOK, "weibo/update.html", gin.H{
		"weibo": weibo,
		// "fcomment":   fcomment,
		"islogin":  islogin,
		"sessions": sessions,
		// "attachs":  attachs,
	})
}
