package post

import (
	"mweibo/middleware/file"
	"mweibo/middleware/image"
	"mweibo/model"
	userservice "mweibo/service/user"
	"mweibo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func RegisterPOST(c *gin.Context) {
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	user := &model.User{}
// 	var err error
// 	code := utils.INVALID_PARAMS
// 	// valid := &validation.Validation{}
// 	// user_userservice.AddUserValid(valid, username, password)
// 	// if valid.HasErrors() {
// 	// 	fmt.Println("valid error")
// 	// 	validator.VErrorMsg(c, valid, code)
// 	// 	return
// 	// }
// 	if !model.IfUserExist(username) {
// 		code = utils.SUCCESS
// 		//ip := c.ClientIP()
// 		user, err = model.AddUser(username, password, ip)
// 		if err != nil {
// 			code = utils.ERROR
// 			//logging.Info("注册入库错误", err.Error())
// 			utils.ResponseJSONError(c, code)
// 			return
// 		}
// 	}
// 	utils.ResponseJSONOK(c, code, user)
// }

func LoginPOST(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	code := utils.INVALID_PARAMS
	// valid := &validation.Validation{}
	// user_userservice.LoginValidWithName(valid, username, password)
	// if valid.HasErrors() {
	// 	validator.VErrorMsg(c, valid, code)
	// 	return
	// }
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	maps["username"] = username
	user, err := model.GetUserObjectByMaps(maps)
	if err != nil {
		code = utils.ERROR_NOT_EXIST_USER
		utils.ResponseJSONOK(c, code, data)
		return
	}
	// 获取加密的密码
	// hashPassword := user.Password
	// if !utils.VerifyString(password, hashPassword) {
	// 	code = utils.ERROR_NOT_EXIST_USER
	// 	utils.ResponseJSONOK(c, code, data)
	// 	return
	// }

	// 2，验证邮箱和密码
	// user, errors := userLoginForm.ValidateAndGetUser(c)
	// if !user.IsActivated() {
	// 	flash.NewWarningFlash(c, "你的账号未激活，请检查邮箱中的注册邮件进行激活。")
	// 	controllers.RedirectRouter(c, "root")
	// 	return
	// }
	if user.Password != utils.MD5(username+password) {
		code = utils.ERROR_NOT_EXIST_USER
		utils.ResponseJSONOK(c, code, data)
		return
	}
	// 3，验证通过 生成token和session
	code = utils.SUCCESS
	// 生成session  使nginx报502错误
	var sok chan int = make(chan int, 1)
	go userservice.LoginSession(c, user, sok)
	<-sok
	utils.ResponseJSONError(c, code)
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

func UpdateUserName(c *gin.Context) {
	userName := c.PostForm("user_name")
	uid, _ := strconv.Atoi(c.Param("id"))
	code := utils.SUCCESS
	err := model.UpdateUserNameByID(userName, uid)
	if err != nil {
		code = utils.ERROR_SQL_UPDATE_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
	utils.SetSession(c, "username", userName)
	utils.ResponseJSONOK(c, code, nil)
}

func UpdateUserAvatar(c *gin.Context) {
	userAvatar, err := c.FormFile("avatar")
	uid, err := strconv.Atoi(c.Param("id"))
	fileName := userAvatar.Filename
	code := utils.SUCCESS
	if err != nil {
		code = utils.ERROR
		utils.ResponseJSONError(c, code)
		return
	}
	if !image.CheckImageExt(fileName) {
		code = utils.ERROR_IMAGE_BAD_EXT
		utils.ResponseJSONError(c, code)
		return
	}
	if !image.CheckImageSize2(userAvatar) {
		code = utils.ERROR_IMAGE_TOO_LARGE
		utils.ResponseJSONError(c, code)
		return
	}
	filePath := "upload/avatar/" + c.Param("id")
	filePath, err = file.CreatePathInToday(filePath)
	if err != nil {
		code = utils.ERROR_FILE_CREATE_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
	fullFileName := filePath + "/" + fileName
	err = c.SaveUploadedFile(userAvatar, fullFileName)
	if err != nil {
		code = utils.ERROR_FILE_SAVE_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
	err = model.UpdateUserAvatarByID(fullFileName, uid)
	if err != nil {
		code = utils.ERROR
		utils.ResponseJSONError(c, code)
		return
	}
	utils.SetSession(c, "useravatar", "/"+fullFileName)
	utils.ResponseJSONOK(c, code, fullFileName)
}

////////////////////////////////////GinWeibo///////////////////////////////////////////////

// // 编辑用户
// func UpdateUserPOST(c *gin.Context, currentUser *userModel.User) {
// 	id, err := controllers.GetIntParam(c, "id")
// 	if err != nil {
// 		controllers.Render404(c)
// 		return
// 	}
// 	// 只能更新自己
// 	if ok := policies.UserPolicyUpdate(c, currentUser, id); !ok {
// 		return
// 	}
// 	// 验证参数和更新用户
// 	userUpdateForm := &userRequest.UserUpdateForm{
// 		Name:                 c.PostForm("name"),
// 		Password:             c.PostForm("password"),
// 		PasswordConfirmation: c.PostForm("password_confirmation"),
// 	}
// 	errors := userUpdateForm.ValidateAndSave(currentUser)
// 	if len(errors) != 0 {
// 		flash.SaveValidateMessage(c, errors)
// 		controllers.RedirectRouter(c, "users.edit", currentUser.ID)
// 		return
// 	}
// 	flash.NewSuccessFlash(c, "个人资料更新成功！")
// 	controllers.RedirectRouter(c, "users.show", currentUser.ID)
// }
