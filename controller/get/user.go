package get

import (
	ctr "mweibo/controller"
	"mweibo/middleware/jwt"
	"mweibo/middleware/redis"
	"mweibo/model"
	userservice "mweibo/service/user"
	"mweibo/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterGET(c *gin.Context) {
	islogin := userservice.IsLogin(c)
	sessions := userservice.GetSessions(c)
	c.HTML(http.StatusOK, "user/register.html", gin.H{
		//"title":    "用户注册",
		"islogin":  islogin,
		"sessions": sessions,
	})
}

func LoginGET(c *gin.Context) {
	islogin := userservice.IsLogin(c)
	sessions := userservice.GetSessions(c)
	c.HTML(http.StatusOK, "user/login.html", gin.H{
		"islogin":  islogin,
		"sessions": sessions,
	})
}

func LogoutGET(c *gin.Context) {
	code := utils.SUCCESS
	userservice.LogoutSession(c)
	utils.ResponseJSONError(c, code)
}

func GetUser(c *gin.Context) {
	err := redis.Lpush("reg:username", utils.GenRandCode(6)+"t@t.com", 5)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  err,
			"data": make(map[string]interface{}),
		})
		return
	}
	res, _ := redis.Brpop("reg:username")
	if res == "" {
		time.Sleep(time.Second * 3)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1002,
		"msg":  "pop",
		"data": res,
	})
	return
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name := c.Query("name"); name != "" {
		maps["name"] = name
		data["user"], _ = model.GetUserObjectByMaps(maps)
	}
	code := utils.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.CodeToMessage(code),
		"data": data,
	})
}

func UserInfo(c *gin.Context) {
	islogin := userservice.IsLogin(c)
	sessions := userservice.GetSessions(c)
	uid := sessions.Userid
	userinfo, _ := model.GetUserObjectByID(uid)
	c.HTML(200, "user/display.html", gin.H{
		"islogin":  islogin,
		"sessions": sessions,
		"userinfo": userinfo,
		//"webname":  webname,
	})
}

// 用户详情
func DisplayUser(c *gin.Context, ctxuser *model.User) {
	id, err := ctr.GetIntParam(c, "id")
	if err != nil {
		ctr.Render404(c)
		return
	}
	// 如果要看的就是当前用户，那么就不用再去数据库中获取了
	user := ctxuser
	if id != int(ctxuser.ID) {
		user, err = model.GetUserByID(id)
	}
	if err != nil || user == nil {
		ctr.Render404(c)
		return
	}
	// 获取分页参数
	// statusesAllLength, _ := model.GetUserAllStatusCount(int(user.ID))
	// offset, limit, currentPage, pageTotalCount := ctr.GetPageQuery(c, 10, statusesAllLength)
	// if currentPage > pageTotalCount {
	// 	ctr.Redirect(c, named.G("users.show", id)+"?page=1", false)
	// 	return
	// }
	// 获取用户的微博
	//weiboList, _ := model.ListUserWeibos(int(user.ID), offset, limit)
	weiboList, _ := model.ListUserWeibos(int(user.ID))
	weibos := make([]*model.Weibo, 0)
	for _, v := range weiboList {
		weibos = append(weibos, v)
	}
	// 获取关注/粉丝
	followingsCount, _ := model.CountFollowings(id)
	followersCount, _ := model.CountFollowers(id)
	isFollowing := false
	if id != int(ctxuser.ID) {
		isFollowing = model.IsFollowing(int(ctxuser.ID), id)
	}
	ctr.Render(c, "user/show.html",
		//pagination.CreatePaginationFillToTplData(c, "page", currentPage, pageTotalCount,
		gin.H{
			"userData":        user,
			"weiboList":       weibos,
			"followingsCount": followingsCount,
			"followersCount":  followersCount,
			"isFollowing":     isFollowing,
			//"statusesLength":  statusesAllLength,
		})
}

func RefreshToken(c *gin.Context) {
	token := c.Query("token")
	newToken, time, _ := jwt.RefreshToken(token)
	data := make(map[string]interface{})
	data["token"] = newToken
	data["exp_time"] = time
	code := utils.SUCCESS
	utils.ResponseJSONOK(c, code, data)
}

func IfUsernameExist(c *gin.Context) {
	name := c.DefaultQuery("username", "")
	code := utils.SUCCESS
	data := make(map[string]interface{})
	if model.IfUsernameExist(name) {
		data["is_used"] = 1
	} else {
		data["is_used"] = 0
	}
	utils.ResponseJSONOK(c, code, data)
}

// // 编辑用户页面
// func UpdateUserGET(c *gin.Context, currentUser *userModel.User) {
// 	id, err := controllers.GetIntParam(c, "id")
// 	if err != nil {
// 		controllers.Render404(c)
// 		return
// 	}
// 	// 只能查看自己的编辑页面
// 	if ok := policies.UserPolicyUpdate(c, currentUser, id); !ok {
// 		return
// 	}
// 	controllers.Render(c, "user/edit.html", gin.H{
// 		"userData": viewmodels.NewUserViewModelSerializer(currentUser),
// 	})
// }
