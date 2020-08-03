package controller

import (
	"mweibo/model"
	"mweibo/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func writeJSON(c *gin.Context, h gin.H) {
	if _, ok := h["succeed"]; !ok {
		h["succeed"] = false
	}
	c.JSON(http.StatusOK, h)
}

func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register.html", nil)
}

func RegisterPost(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		res["message"] = "Register message can not be null."
		return
	}
	user := &model.User{
		Username: username,
		Password: password,
	}
	//user.Password = utils.MD5(user.Username + user.Password)
	user.Password = utils.MD5(username + password)
	if err := model.CreateUser(user); err != nil {
		res["message"] = "Username already exists."
		return
	}
	res["succeed"] = true
	c.Redirect(http.StatusSeeOther, "/login")
}

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login.html", nil)
}

func LoginPost(c *gin.Context) {
	var user *model.User
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.HTML(http.StatusOK, "user/login.html", gin.H{
			"message": "Login message can not be null.",
		})
		return
	}
	user, err := model.GetUserByUsername(username)
	if user.Password != utils.MD5(username+password) || err != nil {
		c.HTML(http.StatusOK, "user/login.html", gin.H{
			"message": "Username or password error.",
		})
		return
	}
	session := sessions.Default(c)
	session.Clear()
	session.Set("UserID", user.ID)
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
}
