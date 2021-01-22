package post

import (
	"fmt"
	"mweibo2/model"
	"mweibo2/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateWeiboPOST(c *gin.Context) {
	content := c.PostForm("content")
	userid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	weibo := &model.Weibo{
		UserID:  uint(userid),
		Content: content,
	}
	err := model.CreateWeibo(weibo)
	if err != nil {
		fmt.Println("Create weibo err: ", err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}

func UpdateWeiboPOST(c *gin.Context) {
	weiboid, _ := strconv.Atoi(c.Param("id"))
	content := c.PostForm("content")
	userid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	oldweibo, _ := model.GetWeiboObjectByID(weiboid)
	if uint(userid) != oldweibo.UserID {
		fmt.Println("Current user is not the weibo user")
		return
	}
	weibo := &model.Weibo{Content: content}
	model.UpdateWeibo(weibo, weiboid)
	c.Redirect(http.StatusSeeOther, "/weibo/"+c.Param("id"))
}

func DeleteWeibo(c *gin.Context) {
	weiboid, _ := strconv.Atoi(c.Param("id"))
	weibo := &model.Weibo{}
	weibo.ID = uint(weiboid)
	userid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	oldweibo, _ := model.GetWeiboObjectByID(weiboid)
	if uint(userid) != oldweibo.UserID {
		fmt.Println("Current user is not the weibo user")
		return
	}
	err := model.DeleteWeibo(weibo)
	if err != nil {
		fmt.Println("Delete weibo err: ", err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}
