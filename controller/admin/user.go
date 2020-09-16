package admin

// func DeleteUser(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	code := utils.SUCCESS
// 	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
// 	isadmin := user_service.IsAdmin(uid)
// 	if isadmin == "0" {
// 		code = utils.UNPASS
// 		utils.JsonErrResponse(c, code)
// 		return
// 	}
// 	err := user_service.DelUserByID(id)
// 	if err != nil {
// 		log.Print("api.v1.user.deluser.deluserbyid:err:", code)
// 		code = utils.ERROR_SQL_DELETE_FAIL
// 		utils.JsonErrResponse(c, code)
// 		return
// 	}
// 	utils.JsonOkResponse(c, code, nil)
// }
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
