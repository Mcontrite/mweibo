package post

import (
	"mweibo/model"
	"mweibo/pkgs"
	"mweibo/service"
	"mweibo/utils"

	"github.com/gin-gonic/gin"
)

// func RegisterPOST(c *gin.Context) {
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	user := &model.User{}
// 	var err error
// 	code := pkgs.INVALID_PARAMS
// 	// valid := &validation.Validation{}
// 	// user_service.AddUserValid(valid, username, password)
// 	// if valid.HasErrors() {
// 	// 	fmt.Println("valid error")
// 	// 	validator.VErrorMsg(c, valid, code)
// 	// 	return
// 	// }
// 	if !model.IfUserExist(username) {
// 		code = pkgs.SUCCESS
// 		//ip := c.ClientIP()
// 		user, err = model.AddUser(username, password, ip)
// 		if err != nil {
// 			code = pkgs.ERROR
// 			//logging.Info("注册入库错误", err.Error())
// 			pkgs.ResponseJSONError(c, code)
// 			return
// 		}
// 	}
// 	pkgs.ResponseJSONOK(c, code, user)
// }

func LoginPOST(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	code := pkgs.INVALID_PARAMS
	// valid := &validation.Validation{}
	// user_service.LoginValidWithName(valid, username, password)
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
		code = pkgs.ERROR_NOT_EXIST_USER
		pkgs.ResponseJSONOK(c, code, data)
		return
	}
	// 获取加密的密码
	// hashPassword := user.Password
	// if !utils.VerifyString(password, hashPassword) {
	// 	code = pkgs.ERROR_NOT_EXIST_USER
	// 	pkgs.ResponseJSONOK(c, code, data)
	// 	return
	// }
	if user.Password != utils.MD5(username+password) {
		code = pkgs.ERROR_NOT_EXIST_USER
		pkgs.ResponseJSONOK(c, code, data)
		return
	}
	// 3，验证通过 生成token和session
	code = pkgs.SUCCESS
	// 生成session  使nginx报502错误
	var sok chan int = make(chan int, 1)
	go service.LoginSession(c, user, sok)
	<-sok
	pkgs.ResponseJSONError(c, code)
}
