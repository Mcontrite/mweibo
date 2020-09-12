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
	user.ActiveToken = string(utils.CreateRandomBytes(30))
	user.RememberMeToken = string(utils.CreateRandomBytes(10))
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
	session.Set(SESSION_KEY, user.ID)
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusSeeOther, "/login")
}

func ListUsers(c *gin.Context) {
	user, _ := c.Get(CONTEXT_USER_KEY)
	users, _ := model.ListUsers()
	c.HTML(http.StatusOK, "admin/users.html", gin.H{
		"user":  user,
		"users": users,
	})
}

func UpdateUserAvatar(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	avatarurl := c.PostForm("avatarurl")
	ctxuser, ok := c.Get(CONTEXT_USER_KEY)
	if !ok {
		res["message"] = "Get user error"
		return
	}
	user := new(model.User)
	user, ok = ctxuser.(*model.User) // 接口类型转换为实体类型
	err := model.UpdateUserAvatar(user, avatarurl)
	if err != nil {
		res["message"] = "Update avatar error " + err.Error()
		return
	}
	res["secceed"] = true
	res["user"] = model.User{
		Avatar: avatarurl,
	}
}

func BindUserEmail(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	email := c.PostForm("email")
	ctxuser, ok := c.Get(CONTEXT_USER_KEY)
	if !ok {
		res["message"] = "Get user error"
		return
	}
	user, _ := ctxuser.(*model.User)
	if len(user.Email) > 0 {
		res["message"] = "Email can't be null"
		return
	}
	err := model.UpdateUserEmail(user, email)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["secceed"] = true
}

func UnbindUserEmail(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	ctxuser, ok := c.Get(CONTEXT_USER_KEY)
	if !ok {
		res["message"] = "Get user error"
		return
	}
	user, _ := ctxuser.(*model.User)
	if user.Email == "" {
		res["message"] = "Email is null"
		return
	}
	err := model.UpdateUserEmail(user, "")
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["secceed"] = true
}
