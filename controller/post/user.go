package post

import (
	"mweibo/model"
	userservice "mweibo/service/user"
	"mweibo/utils"

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

// func AddUser(c *gin.Context) {
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	user := &model.User{}
// 	valid := &validation.Validation{}
// 	var err error
// 	code := rcode.INVALID_PARAMS
// 	user_service.AddUserValid(valid, username, password)
// 	if valid.HasErrors() {
// 		fmt.Println("valid error")
// 		validator.VErrorMsg(c, valid, code)
// 		return
// 	}
// 	if !model.ExistUserByName(username) {
// 		code = rcode.SUCCESS
// 		ip := c.ClientIP()
// 		user, err = model.AddUser(username, password, ip)
// 		if err != nil {
// 			code = rcode.ERROR
// 			fmt.Println("model error")
// 			logging.Info("注册入库错误", err.Error())
// 			app.JsonErrResponse(c, code)
// 			return
// 		}
// 	}
// 	app.JsonOkResponse(c, code, user)
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
	// 2，验证邮箱和密码
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	maps["username"] = username
	user, err := model.GetUser(maps)
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

//////////////////////////////////////////////////////////////////////////////////////////
// func GetUser(c *gin.Context) {
// 	err := gredis.Lpush("reg:username", utils.GenRandCode(6)+"t@t.com", 5)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code": 1001,
// 			"msg":  err,
// 			"data": make(map[string]interface{}),
// 		})
// 		return
// 	}
// 	res, _ := gredis.Brpop("reg:username")
// 	if res == "" {
// 		time.Sleep(time.Second * 3)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 1002,
// 		"msg":  "pop",
// 		"data": res,
// 	})
// 	return
// 	maps := make(map[string]interface{})
// 	data := make(map[string]interface{})
// 	if name := c.Query("name"); name != "" {
// 		maps["name"] = name
// 		data["user"], _ = model.GetUser(maps)
// 	}
// 	code := rcode.SUCCESS
// 	c.JSON(http.StatusOK, gin.H{
// 		"code": code,
// 		"msg":  rcode.GetMessage(code),
// 		"data": data,
// 	})
// }

// func DeleteUser(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	code := rcode.SUCCESS
// 	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
// 	isadmin := user_service.IsAdmin(uid)
// 	if isadmin == "0" {
// 		code = rcode.UNPASS
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	err := user_service.DelUserByID(id)
// 	if err != nil {
// 		log.Print("api.v1.user.deluser.deluserbyid:err:", code)
// 		code = rcode.ERROR_SQL_DELETE_FAIL
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	app.JsonOkResponse(c, code, nil)
// }

// func ResetUserPassword(c *gin.Context) {
// 	oldpassword := c.PostForm("password_old")
// 	newpassword := c.PostForm("password_new")
// 	uid := c.Param("id")
// 	code := rcode.SUCCESS
// 	// 验证原来的密码正确性
// 	maps := make(map[string]interface{})
// 	maps["id"] = uid
// 	user, err := model.GetUser(maps)
// 	if err != nil {
// 		code = rcode.ERROR_NOT_EXIST_USER
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	// 获取加密的密码
// 	hashPassword := user.Password
// 	if !utils.VerifyString(oldpassword, hashPassword) {
// 		code = rcode.ERROR
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	user.Password, _ = utils.BcryptString(newpassword)
// 	err = user_service.ResetPassword(user.Password, int(user.ID))
// 	if err != nil {
// 		code = rcode.ERROR
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	app.JsonOkResponse(c, code, nil)
// }

// func ResetUserAvatar(c *gin.Context) {
// 	userAvatar, err := c.FormFile("avatar")
// 	uid, err := strconv.Atoi(c.Param("id"))
// 	fileName := userAvatar.Filename
// 	code := rcode.SUCCESS
// 	if err != nil {
// 		code = rcode.ERROR
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	if !upload.CheckImageExt(fileName) {
// 		code = rcode.ERROR_IMAGE_BAD_EXT
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	if !upload.CheckImageSize2(userAvatar) {
// 		code = rcode.ERROR_IMAGE_TOO_LARGE
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	filePath := "upload/avatar/" + c.Param("id")
// 	filePath, err = file.CreatePathInToday(filePath)
// 	if err != nil {
// 		code = rcode.ERROR_FILE_CREATE_FAIL
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	fullFileName := filePath + "/" + fileName
// 	err = c.SaveUploadedFile(userAvatar, fullFileName)
// 	if err != nil {
// 		code = rcode.ERROR_FILE_SAVE_FAIL
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	err = user_service.ResetAvatar(fullFileName, uid)
// 	if err != nil {
// 		code = rcode.ERROR
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	session.SetSession(c, "useravatar", "/"+fullFileName)
// 	app.JsonOkResponse(c, code, fullFileName)
// }

// func ResetUserName(c *gin.Context) {
// 	userName := c.PostForm("user_name")
// 	uid, _ := strconv.Atoi(c.Param("id"))
// 	code := rcode.SUCCESS
// 	err := user_service.ResetName(userName, uid)
// 	if err != nil {
// 		code = rcode.ERROR_SQL_UPDATE_FAIL
// 		app.JsonErrResponse(c, code)
// 		return
// 	}
// 	session.SetSession(c, "username", userName)
// 	app.JsonOkResponse(c, code, nil)
// }

// func CheckNameUsed(c *gin.Context) {
// 	name := c.DefaultQuery("username", "")
// 	code := rcode.SUCCESS
// 	data := make(map[string]interface{})
// 	if model.ExistUserByName(name) {
// 		data["is_used"] = 1
// 	} else {
// 		data["is_used"] = 0
// 	}
// 	app.JsonOkResponse(c, code, data)
// }

// func RefreshToken(c *gin.Context) {
// 	token := c.Query("token")
// 	newToken, time, _ := utils.RefreshToken(token)
// 	data := make(map[string]interface{})
// 	data["token"] = newToken
// 	data["exp_time"] = time
// 	code := rcode.SUCCESS
// 	app.JsonOkResponse(c, code, data)
// }
