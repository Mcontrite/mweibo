package post

import (
	ctr "mweibo/controller"
	"mweibo/middleware/flash"
	"mweibo/middleware/validate"
	"mweibo/model"
	"mweibo/utils"

	"github.com/gin-gonic/gin"
)

// 显示重置密码的邮箱发送页面
func SendResetLinkEmailGET(c *gin.Context) {
	ctr.Render(c, "password/email.html", gin.H{})
}

// 更新密码页面
func ResetPasswordGET(c *gin.Context) {
	token := c.Param("token")
	p, err := model.GetPwrtByToken(token)
	if err != nil {
		ctr.Render404(c)
		return
	}
	ctr.Render(c, "password/reset.html", gin.H{
		"token": token,
		"email": p.Email,
	})
}

// 更新密码
func ResetPasswordPOST(c *gin.Context) {
	passwordForm := &validate.PassWordResetForm{
		Token:                c.PostForm("token"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}
	user, errs := passwordForm.ValidateAndUpdateUser()
	if len(errs) != 0 || user == nil {
		flash.SaveValidateMessage(c, errs)
		//controllers.RedirectRouter(c, "password.reset", "token", c.PostForm("token"))
		return
	}
	flash.NewSuccessFlash(c, "重置密码成功")
	//controllers.RedirectRouter(c, "root")
}

func ResetUserPassword(c *gin.Context) {
	oldpassword := c.PostForm("password_old")
	newpassword := c.PostForm("password_new")
	uid := c.Param("id")
	code := utils.SUCCESS
	// 验证原来的密码正确性
	maps := make(map[string]interface{})
	maps["id"] = uid
	user, err := model.GetUserObjectByMaps(maps)
	if err != nil {
		code = utils.ERROR_NOT_EXIST_USER
		utils.ResponseJSONError(c, code)
		return
	}
	// 获取加密的密码
	hashPassword := user.Password
	if !utils.VerifyString(oldpassword, hashPassword) {
		code = utils.ERROR
		utils.ResponseJSONError(c, code)
		return
	}
	user.Password, _ = utils.BcryptString(newpassword)
	// var wmap = make(map[string]interface{})
	// err := model.UpdateUser(wmap, map[string]interface{}{"password": password})
	err = model.ResetPasswordByID(user.Password, int(user.ID))
	if err != nil {
		code = utils.ERROR
		utils.ResponseJSONError(c, code)
		return
	}
	utils.ResponseJSONOK(c, code, nil)
}
