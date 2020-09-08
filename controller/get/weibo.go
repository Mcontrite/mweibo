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
	// fcomment, _ := model.GetWeiboFirstCommentByTid(weiboID)
	// fcomment.MessageFmt = html.UnescapeString(fcomment.MessageFmt)
	islogin := userservice.IsLogin(c)
	sessions := userservice.GetSessions(c)
	// commentlist, _ := model.GetWeiboCommentListByTid(weiboID, 500, 1)
	// commentlistLen := len(commentlist)
	// attachs, _ := model.GetAttachsByCommentID(int(fcomment.ID))
	// isfav, _ := model.CheckFavourite(sessions.Userid, weiboID)
	// model.UpdateWeiboViewsCnt(weiboID)
	c.HTML(http.StatusOK, "weibo/display.html", gin.H{
		"weibo": weibo,
		// "fcomment":         fcomment,
		"islogin":  islogin,
		"sessions": sessions,
		// "commentlist":      commentlist,
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

func WeiboAddCommentGET(c *gin.Context) {
	weiboID, _ := strconv.Atoi(c.Param("id"))
	sessions := userservice.GetSessions(c)
	islogin := userservice.IsLogin(c)
	c.HTML(http.StatusOK, "advance_comment.html", gin.H{
		"sessions": sessions,
		"islogin":  islogin,
		"weibo_id": weiboID,
	})
}
