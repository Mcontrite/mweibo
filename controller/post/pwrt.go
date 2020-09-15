package post

import (
	ctr "mweibo/controller"
	"mweibo/middleware/flash"
	"mweibo/middleware/validate"
	"mweibo/model"

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
