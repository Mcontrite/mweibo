package controller

import (
	"mweibo/model"

	"strconv"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	session := sessions.Default(c)
	sessionuserid := session.Get(SESSION_KEY)
	userid, _ := sessionuserid.(uint)
	code := c.PostForm("captcha_code")
	captchaid := session.Get(SESSION_CAPTCHA)
	capid := captchaid.(string)
	session.Delete(SESSION_CAPTCHA)
	if !captcha.VerifyString(capid, code) {
		res["message"] = "Captcha error"
		return
	}
	weiboid := c.PostForm("weiboid")
	content := c.PostForm("content")
	if len(content) == 0 {
		res["message"] = "content cannot be null"
		return
	}
	weiid, _ := strconv.ParseUint(weiboid, 10, 64)
	comment := &model.Comment{
		WeiboID: uint(weiid),
		Content: content,
		UserID:  userid,
	}
	err := model.CreateComment(comment)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func ReadComment(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	commentid := c.Param("id")
	comid, _ := strconv.ParseUint(commentid, 10, 64)
	comment := &model.Comment{}
	comment.ID = uint(comid)
	err := model.UpdateCommentIsRead(comment)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func ReadAllComments(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	err := model.SetAllCommentRead()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func DeleteComment(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	session := sessions.Default(c)
	sessionuserid := session.Get(SESSION_KEY)
	userid := sessionuserid.(uint)
	commentid := c.Param("id")
	comid, _ := strconv.ParseUint(commentid, 10, 64)
	comment := &model.Comment{
		UserID: uint(userid),
	}
	comment.ID = uint(comid)
	err := model.DeleteComment(comment)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}
