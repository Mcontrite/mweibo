package service

import (
	"mweibo/model"
	"mweibo/pkgs"
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
	pkgs.SetSession(c, "username", user.Username)
	pkgs.SetSession(c, "userid", strconv.Itoa(int(user.ID)))
	pkgs.SetSession(c, "useravatar", user.Avatar)
	// pkgs.SetSession(c, "userarticlecnt", strconv.Itoa(user.ArticlesCnt))
	// pkgs.SetSession(c, "userreplycnt", strconv.Itoa(user.ReplysCnt))
	// pkgs.SetSession(c, "usersayingcnt", strconv.Itoa(user.SayingsCnt))
	// pkgs.SetSession(c, "usercommentcnt", strconv.Itoa(user.CommentsCnt))
	// pkgs.SetSession(c, "isadmin", IsAdmin(user.GroupID))
	sok <- 1
}

func GetSessions(c *gin.Context) (sessions *UserSession) {
	username := pkgs.GetSession(c, "username")
	userid, _ := strconv.Atoi(pkgs.GetSession(c, "userid"))
	useravatar := pkgs.GetSession(c, "useravatar")
	// userarticlecnt, _ := strconv.Atoi(pkgs.GetSession(c, "userarticlecnt"))
	// userreplycnt, _ := strconv.Atoi(pkgs.GetSession(c, "userreplycnt"))
	// usersayingcnt, _ := strconv.Atoi(pkgs.GetSession(c, "usersayingcnt"))
	// usercommentcnt, _ := strconv.Atoi(pkgs.GetSession(c, "usercommentcnt"))
	isadmin := pkgs.GetSession(c, "isadmin")
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
	pkgs.DeleteSession(c, "username")
	pkgs.DeleteSession(c, "userid")
	pkgs.DeleteSession(c, "useravatar")
	// pkgs.DeleteSession(c, "userarticlecnt")
	// pkgs.DeleteSession(c, "userreplycnt")
	// pkgs.DeleteSession(c, "usersayingcnt")
	// pkgs.DeleteSession(c, "usercommentcnt")
	pkgs.DeleteSession(c, "isadmin")
}

func IsLogin(c *gin.Context) bool {
	username := pkgs.GetSession(c, "username")
	if len(username) > 0 {
		return true
	}
	return false
}
