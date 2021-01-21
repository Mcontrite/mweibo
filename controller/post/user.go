package post

import (
	"fmt"
	"mweibo2/model"
	"mweibo2/service"
	"mweibo2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPOST(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		fmt.Println("Username or password can not be empty")
		return
	}
	user := &model.User{
		Username: username,
		Password: password,
	}
	user.Password = utils.MD5(username + password)
	if err := model.UserCreate(user); err != nil {
		fmt.Println("Register error: ", err)
		return
	}
	// c.JSON(http.StatusOK, nil)
	c.Redirect(http.StatusSeeOther, "/login")
}

func LoginPOST(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	m := make(map[string]interface{})
	m["username"] = username
	user, err := model.GetUserObjectByMap(m)
	if err != nil {
		fmt.Println("get user err: ", err)
		return
	}
	if user.Password != utils.MD5(username+password) {
		fmt.Println("Password err...")
		return
	}
	ok := make(chan int, 1)
	go service.LoginSession(c, user, ok)
	<-ok
	// c.JSON(http.StatusNotFound, nil)
	c.Redirect(http.StatusSeeOther, "/")
}
