package service

import (
	"mweibo2/model"
	"mweibo2/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsLogin(c *gin.Context) bool {
	username := utils.GetSession(c, "username")
	if len(username) > 0 {
		return true
	}
	return false
}

func LoginSession(c *gin.Context, user model.User, ok chan int) {
	utils.SetSession(c, "username", user.Username)
	utils.SetSession(c, "userid", strconv.Itoa(int(user.ID)))
	ok <- 1
}

func LogoutSession(c *gin.Context) {
	utils.DelSession(c, "username")
	utils.DelSession(c, "userid")
}

type UserSession struct {
	Username string
	UserID   int
}

func GetUserSession(c *gin.Context) *UserSession {
	username := utils.GetSession(c, "username")
	userid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	return &UserSession{
		Username: username,
		UserID:   userid,
	}
}
