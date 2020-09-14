package get

import (
	ctr "mweibo/controller"
	"mweibo/model"

	"github.com/gin-gonic/gin"
)

// 用户关注列表
func ListFollowingsGET(c *gin.Context, ctxuser *model.User) {
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
	// followingsLength, _ := followerModel.FollowingsCount(id)
	// offset, limit, currentPage, pageTotalCount := ctr.GetPageQuery(c, 10, followingsLength)
	// if currentPage > pageTotalCount {
	// 	ctr.Redirect(c, named.G("users.followings")+"?page=1", false)
	// 	return
	// }
	// 获取关注者
	// followingList, _ := model.ListUserFollowings(id, offset, limit)
	followingList, _ := model.ListUserFollowings(id)
	userlist := make([]*model.User, 0)
	for _, v := range followingList {
		userlist = append(userlist, v)
	}
	ctr.Render(c, "user/show_follow.html",
		//pagination.CreatePaginationFillToTplData(c, "page", currentPage, pageTotalCount,
		gin.H{
			"title":    user.Username + " 关注的人",
			"userData": user,
			"users":    userlist,
		})
}

// 用户粉丝列表
func ListFollowersGET(c *gin.Context, ctxuser *model.User) {
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
	// followersLength, _ := followerModel.FollowersCount(id)
	// offset, limit, currentPage, pageTotalCount := ctr.GetPageQuery(c, 10, followersLength)
	// if currentPage > pageTotalCount {
	// 	ctr.Redirect(c, named.G("users.followers")+"?page=1", false)
	// 	return
	// }
	// 获取关注者
	// followers, _ := followerModel.Followers(id, offset, limit)
	followers, _ := model.ListUserFollowers(id)
	userlist := make([]*model.User, 0)
	for _, v := range followers {
		userlist = append(userlist, v)
	}
	ctr.Render(c, "user/show_follow.html",
		//pagination.CreatePaginationFillToTplData(c, "page", currentPage, pageTotalCount,
		gin.H{
			"title":    user.Username + " 的粉丝",
			"userData": user,
			"users":    userlist,
		})
}
