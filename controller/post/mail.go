package post

import (
	// ctr "mweibo/controller"
	"mweibo/middleware/flash"
	"mweibo/middleware/mail"
	"mweibo/middleware/validate"
	"mweibo/model"
	userservice "mweibo/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 邮箱验证、用户激活
func ConfirmEmail(c *gin.Context) {
	token := c.Param("token")
	user, err := model.GetUserByActiveToken(token)
	if user == nil || err != nil {
		//ctr.Render404(c)
		return
	}
	// 更新用户
	user.IsActive = true
	user.ActiveToken = ""
	if err = user.UpdateUser(); err != nil {
		// flash.NewSuccessFlash(c, "用户激活失败: "+err.Error())
		// ctr.RedirectRouter(c, "root")
		return
	}
	var sok chan int = make(chan int, 1)
	go userservice.LoginSession(c, *user, sok)
	//flash.NewSuccessFlash(c, "恭喜你，激活成功！")
	//ctr.RedirectRouter(c, "users.show", user.ID)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func sendConfirmEmail(user *model.User) error {
	subject := "感谢注册 Weibo 应用！请确认你的邮箱。"
	tpl := "mail/confirm.html"
	confirmURL := "signup/confirm/" + user.ActiveToken
	return mail.SendMail([]string{user.Email}, subject, tpl, gin.H{"confirmURL": confirmURL})
}

func sendResetEmail(pwrt *model.PasswordReset) error {
	subject := "重置密码！请确认你的邮箱。"
	tpl := "mail/reset_password.html"
	//passwordResetURL := named.G("password.reset", "token", pwrt.Token)
	passwordResetURL := "password/reset/" + pwrt.Token
	return mail.SendMail([]string{pwrt.Email}, subject, tpl, gin.H{"passwordResetURL": passwordResetURL})
}

// 发送重设密码邮件
func SendResetLinkEmail(c *gin.Context) {
	email := c.PostForm("email")
	passwordForm := &validate.PasswordEmailForm{
		Email: email,
	}
	pwrt, errors := passwordForm.ValidateAndGetToken()
	if len(errors) != 0 || pwrt == nil {
		flash.SaveValidateMessage(c, errors)
		//ctr.RedirectRouter(c, "password.request")
		return
	}
	if err := sendResetEmail(pwrt); err != nil {
		flash.NewDangerFlash(c, "重置密码邮件发送失败: "+err.Error())
		// 删除 token
		model.DeletePwrtByEmail(pwrt.Email)
	} else {
		flash.NewSuccessFlash(c, "重置密码已发送到你的邮箱上，请注意查收。")
	}
	//ctr.RedirectRouter(c, "password.request")
}
