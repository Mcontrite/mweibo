package validate

import (
	"mweibo/model"
)

type PasswordEmailForm struct {
	Email string
}

func (p *PasswordEmailForm) emailExistValidate() ValidateFunc {
	return func() (msg string) {
		if _, err := model.GetUserByEmail(p.Email); err == nil {
			return ""
		}
		return "该邮箱不存在"
	}
}

func (p *PasswordEmailForm) Validate() (errs []string) {
	errs = RunValidates(
		VlidFuncsMap{
			"email": {
				ValidateRequired(p.Email),
				ValidateMaxLength(p.Email, 255),
				ValidateEmail(p.Email),
				p.emailExistValidate(),
			},
		},
		VlidMsgsMap{
			"email": {
				"邮箱不能为空",
				"邮箱长度不能大于 255 个字符",
				"邮箱格式错误",
				"该邮箱不存在",
			},
		},
	)
	return errs
}

// 验证参数并且创建验证 pwrt 的 token
func (p *PasswordEmailForm) ValidateAndGetToken() (pwrt *model.PasswordReset, errs []string) {
	errs = p.Validate()
	if len(errs) != 0 {
		return nil, errs
	}
	pwrt = &model.PasswordReset{
		Email: p.Email,
	}
	if err := pwrt.CreatePwrt(); err != nil {
		errs = append(errs, "失败: "+err.Error())
		return nil, errs
	}
	return pwrt, []string{}
}

type PassWordResetForm struct {
	Email                string
	Token                string
	Password             string
	PasswordConfirmation string
}

func (p *PassWordResetForm) tokenExistValidate() ValidateFunc {
	return func() (msg string) {
		if pwrt, err := model.GetPwrtByToken(p.Token); err == nil {
			p.Email = pwrt.Email
			return ""
		}
		return "该 token 不存在"
	}
}

func (p *PassWordResetForm) Validate() (errs []string) {
	errs = RunValidates(
		VlidFuncsMap{
			"password": {
				ValidateRequired(p.Password),
				ValidateMinLength(p.Password, 6),
				ValidateEqual(p.Password, p.PasswordConfirmation),
			},
			"token": {
				ValidateRequired(p.Token),
				p.tokenExistValidate(),
			},
		},
		VlidMsgsMap{
			"password": {
				"密码不能为空",
				"密码长度不能小于 6 个字符",
				"两次输入的密码不一致",
			},
			"token": {
				"token 不能为空",
				"该 token 不存在",
			},
		},
	)
	return
}

// 验证参数并且创建验证 pwrt 的 token
func (p *PassWordResetForm) ValidateAndUpdateUser() (user *model.User, errs []string) {
	errs = p.Validate()
	if len(errs) != 0 {
		return nil, errs
	}
	// 验证成功，删除 token
	if err := model.DeletePwrtByToken(p.Token); err != nil {
		errs = append(errs, "重置密码失败: "+err.Error())
		return nil, errs
	}
	// 更新用户密码
	user, err := model.GetUserByEmail(p.Email)
	if err != nil {
		errs = append(errs, "重置密码失败: "+err.Error())
		return nil, errs
	}
	user.Password = p.Password
	if err = user.UpdateUser(); err != nil {
		errs = append(errs, "重置密码失败: "+err.Error())
		return nil, errs
	}
	return user, []string{}
}
