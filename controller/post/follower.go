package post

import (
	ctr "mweibo/controller"
	"mweibo/middleware/flash"
	"mweibo/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 关注用户
func FollowUser(c *gin.Context, ctxuser *model.User) {
	id, err := ctr.GetIntParam(c, "id")
	if err != nil {
		ctr.Render404(c)
		return
	}
	// if ok := policies.UserPolicyFollow(c, ctxuser, id); !ok {
	// 	return
	// }
	isFollowing := false
	if id != int(ctxuser.ID) {
		isFollowing = model.IsFollowing(int(ctxuser.ID), id)
	}
	if !isFollowing {
		if err := model.FollowSomeUsers(ctxuser.ID, uint(id)); err != nil {
			flash.NewDangerFlash(c, "关注失败: "+err.Error())
		}
	}
	ctr.Redirect(c, "users/show"+strconv.Itoa(id)+"?page=1", false)
}

// 取消关注用户
func UnFollowUser(c *gin.Context, ctxuser *model.User) {
	id, err := ctr.GetIntParam(c, "id")
	if err != nil {
		ctr.Render404(c)
		return
	}
	// if ok := policies.UserPolicyFollow(c, ctxuser, id); !ok {
	// 	return
	// }
	isFollowing := false
	if id != int(ctxuser.ID) {
		isFollowing = model.IsFollowing(int(ctxuser.ID), id)
	}
	if isFollowing {
		if err := model.UnFollowSomeUsers(ctxuser.ID, uint(id)); err != nil {
			flash.NewDangerFlash(c, "取消关注失败: "+err.Error())
		}
	}
	ctr.Redirect(c, "users/show"+strconv.Itoa(id)+"?page=1", false)
}
