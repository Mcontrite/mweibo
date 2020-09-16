package validate

import (
	"mweibo/model"

	"github.com/gin-gonic/gin"
)

// 以后可以改为 tag 来调用验证器函数
type UserCreateForm struct {
	Username             string
	Email                string
	Password             string
	PasswordConfirmation string
}

func (u *UserCreateForm) emailUniqueValidate() ValidateFunc {
	return func() (msg string) {
		if _, err := model.GetUserByEmail(u.Email); err != nil {
			return ""
		}
		return "邮箱已经被注册过了"
	}
}

func (u *UserCreateForm) Validate() (errs []string) {
	errs = RunValidates(
		VlidFuncsMap{
			"name": {
				ValidateRequired(u.Username),
				ValidateMaxLength(u.Username, 50),
			},
			"email": {
				ValidateRequired(u.Email),
				ValidateMaxLength(u.Email, 255),
				ValidateEmail(u.Email),
				u.emailUniqueValidate(),
			},
			"password": {
				ValidateRequired(u.Password),
				ValidateMinLength(u.Password, 6),
				ValidateEqual(u.Password, u.PasswordConfirmation),
			},
		},
		VlidMsgsMap{
			"name": {
				"名称不能为空",
				"名称长度不能大于 50 个字符",
			},
			"email": {
				"邮箱不能为空",
				"邮箱长度不能大于 255 个字符",
				"邮箱格式错误",
				"邮箱已经被注册过了",
			},
			"password": {
				"密码不能为空",
				"密码长度不能小于 6 个字符",
				"两次输入的密码不一致",
			},
		},
	)
	return errs
}

// 验证参数并且创建用户
func (u *UserCreateForm) ValidateAndSave() (user *model.User, errs []string) {
	errs = u.Validate()
	if len(errs) != 0 {
		return nil, errs
	}
	user = &model.User{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
	if err := model.CreateUser(user); err != nil {
		errs = append(errs, "用户创建失败: "+err.Error())
		return nil, errs
	}
	return user, []string{}
}

type UserUpdateForm struct {
	Username             string
	Password             string
	PasswordConfirmation string
}

func (u *UserUpdateForm) Validate() (errs []string) {
	nameValidates := []ValidateFunc{
		ValidateRequired(u.Username),
		ValidateMaxLength(u.Username, 50),
	}
	nameMsgs := []string{
		"名称不能为空",
		"名称长度不能大于 50 个字符",
	}
	pwdValidates := []ValidateFunc{
		ValidateRequired(u.Password),
		ValidateMinLength(u.Password, 6),
		ValidateEqual(u.Password, u.PasswordConfirmation),
	}
	pwdMsgs := []string{
		"密码不能为空",
		"密码长度不能小于 6 个字符",
		"两次输入的密码不一致",
	}
	if u.Password == "" {
		errs = RunValidates(
			VlidFuncsMap{
				"name": nameValidates,
			},
			VlidMsgsMap{
				"name": nameMsgs,
			},
		)
	} else {
		errs = RunValidates(
			VlidFuncsMap{
				"name":     nameValidates,
				"password": pwdValidates,
			},
			VlidMsgsMap{
				"name":     nameMsgs,
				"password": pwdMsgs,
			},
		)
	}
	return errs
}

// 验证参数并且创建用户
func (u *UserUpdateForm) ValidateAndSave(user *model.User) (errs []string) {
	errs = u.Validate()
	if len(errs) != 0 {
		return errs
	}
	// 更新用户
	user.Username = u.Username
	if u.Password != "" {
		user.Password = u.Password
	}
	if err := user.UpdateUser(); err != nil {
		errs = append(errs, "用户更新失败: "+err.Error())
		return errs
	}
	return []string{}
}

type UserLoginForm struct {
	Email    string
	Password string
}

func (u *UserLoginForm) Validate() (errs []string) {
	errs = RunValidates(
		VlidFuncsMap{
			"email": {
				ValidateRequired(u.Email),
				ValidateMaxLength(u.Email, 255),
				ValidateEmail(u.Email),
			},
			"password": {
				ValidateRequired(u.Password),
			},
		},
		VlidMsgsMap{
			"email": {
				"邮箱不能为空",
				"邮箱长度不能大于 255 个字符",
				"邮箱格式错误",
			},
			"password": {
				"密码不能为空",
			},
		},
	)
	return errs
}

// 验证参数并且获取用户
func (u *UserLoginForm) ValidateAndGetUser(c *gin.Context) (user *model.User, errs []string) {
	errs = u.Validate()
	if len(errs) != 0 {
		return nil, errs
	}
	// 通过邮箱获取用户，并且判断密码是否正确
	user, err := model.GetUserByEmail(u.Email)
	if err != nil {
		errs = append(errs, "该邮箱没有注册过用户: "+err.Error())
		return nil, errs
	}
	if user.Password != u.Password {
		//flash.NewDangerFlash(c, "很抱歉，您的邮箱和密码不匹配")
		return nil, errs
	}
	return user, []string{}
}

// func ValidName(v *validation.Validation, username string) {
// 	pass, _ := regexp.MatchString("[a-zA-Z0-9]{3,16}", username)
// 	if !pass {
// 		v.SetError("username", "名称只能是3-16位字母数字组合")
// 	}
// }

// func ValidNameRequired(v *validation.Validation, username string) {
// 	v.Required(username, "username").Message("密码不能为空")
// }

// func ValidPassword(v *validation.Validation, password string) {
// 	v.Required(password, "password").Message("密码不能为空")
// }

// func AddUserValid(v *validation.Validation, username string, password string) {
// 	ValidName(v, username)
// 	ValidPassword(v, password)
// }

// func LoginValidWithName(v *validation.Validation, name string, password string) {
// 	ValidPassword(v, password)
// 	ValidNameRequired(v, name)
// }
