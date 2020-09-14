package get

import (
	ctr "mweibo/controller"
	"mweibo/model"
	userservice "mweibo/service/user"
	"mweibo/utils"
	"net/http"

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
