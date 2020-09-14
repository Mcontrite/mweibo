package admin

// // 删除用户
// func DeleteUser(c *gin.Context, currentUser *userModel.User) {
// 	page := c.DefaultQuery("page", "1")
// 	id, err := controllers.GetIntParam(c, "id")
// 	if err != nil {
// 		controllers.Render404(c)
// 		return
// 	}
// 	// 是否有删除权限
// 	if ok := policies.UserPolicyDestroy(c, currentUser, id); !ok {
// 		return
// 	}
// 	// 删除用户
// 	if err = userModel.Delete(id); err != nil {
// 		flash.NewDangerFlash(c, "删除失败: "+err.Error())
// 	} else {
// 		flash.NewSuccessFlash(c, "成功删除用户！")
// 	}
// 	controllers.Redirect(c, named.G("users.index")+"?page="+page, false)
// }
