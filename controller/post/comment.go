package post

import (
	"fmt"
	"mweibo2/model"
	"mweibo2/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCommentPOST(c *gin.Context) {
	weiboid, _ := strconv.Atoi(c.DefaultPostForm("weiboid", "1"))
	content := c.DefaultPostForm("content", "")
	userid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	comment := &model.Comment{
		WeiboID: uint(weiboid),
		UserID:  uint(userid),
		Content: content,
	}
	err := model.CreateComment(comment)
	if err != nil {
		fmt.Println("Create comment err: ", err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/weibo/"+c.Param("id"))
}

func UpdateCommentPOST(c *gin.Context) {
	commentid, _ := strconv.Atoi(c.DefaultPostForm("commentid", "1"))
	content := c.DefaultPostForm("content", "")
	userid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	oldcomment, _ := model.GetCommentByID(commentid)
	if uint(userid) != oldcomment.UserID {
		fmt.Println("Current User is not the comment user")
		return
	}
	comment := model.Comment{Content: content}
	model.UpdateCommentByID(comment, commentid)
	c.Redirect(http.StatusSeeOther, "/weibo/"+c.Param("id"))
}
