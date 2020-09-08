package service

import (
	"mweibo/model"
	"mweibo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserSession struct {
	Username   string `json:"username"`
	Userid     int    `json:"userid"`
	Useravatar string `json:"useravatar"`
	// Userarticlecnt int    `json:"userarticlecnt"`
	// Userreplycnt   int    `json:"userreplycnt"`
	// Usersayingcnt  int    `json:"usersayingcnt"`
	// Usercommentcnt int    `json:"usercommentcnt"`
	Isadmin string `json:"isadmin"`
}

func LoginSession(c *gin.Context, user model.User, sok chan int) {
	utils.SetSession(c, "username", user.Username)
	utils.SetSession(c, "userid", strconv.Itoa(int(user.ID)))
	utils.SetSession(c, "useravatar", user.Avatar)
	// utils.SetSession(c, "userarticlecnt", strconv.Itoa(user.ArticlesCnt))
	// utils.SetSession(c, "userreplycnt", strconv.Itoa(user.ReplysCnt))
	// utils.SetSession(c, "usersayingcnt", strconv.Itoa(user.SayingsCnt))
	// utils.SetSession(c, "usercommentcnt", strconv.Itoa(user.CommentsCnt))
	// utils.SetSession(c, "isadmin", IsAdmin(user.GroupID))
	sok <- 1
}

func GetSessions(c *gin.Context) (sessions *UserSession) {
	username := utils.GetSession(c, "username")
	userid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	useravatar := utils.GetSession(c, "useravatar")
	// userarticlecnt, _ := strconv.Atoi(utils.GetSession(c, "userarticlecnt"))
	// userreplycnt, _ := strconv.Atoi(utils.GetSession(c, "userreplycnt"))
	// usersayingcnt, _ := strconv.Atoi(utils.GetSession(c, "usersayingcnt"))
	// usercommentcnt, _ := strconv.Atoi(utils.GetSession(c, "usercommentcnt"))
	isadmin := utils.GetSession(c, "isadmin")
	sessions = &UserSession{
		Username:   username,
		Userid:     userid,
		Useravatar: useravatar,
		// Userarticlecnt: userarticlecnt,
		// Userreplycnt:   userreplycnt,
		// Usersayingcnt:  usersayingcnt,
		// Usercommentcnt: usercommentcnt,
		Isadmin: isadmin,
	}
	return
}

func LogoutSession(c *gin.Context) {
	utils.DeleteSession(c, "username")
	utils.DeleteSession(c, "userid")
	utils.DeleteSession(c, "useravatar")
	// utils.DeleteSession(c, "userarticlecnt")
	// utils.DeleteSession(c, "userreplycnt")
	// utils.DeleteSession(c, "usersayingcnt")
	// utils.DeleteSession(c, "usercommentcnt")
	utils.DeleteSession(c, "isadmin")
}

func IsLogin(c *gin.Context) bool {
	username := utils.GetSession(c, "username")
	if len(username) > 0 {
		return true
	}
	return false
}
