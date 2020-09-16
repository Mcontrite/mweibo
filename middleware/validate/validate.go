package validate

import (
	"mweibo/utils"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type (
	ValidateFunc = func() (msg string)       // 验证器函数
	VlidFuncsMap = map[string][]ValidateFunc // 验证器数组 map
	VlidMsgsMap  = map[string][]string       // 错误信息数组
)

func RunValidates(funcsm VlidFuncsMap, msgsm VlidMsgsMap) (errors []string) {
	for mk, mv := range funcsm {
		customMsgArr := msgsm[mk] // 自定义错误信息数组
		for fk, fv := range mv {
			msg := fv()
			if msg != "" {
				if fk < len(customMsgArr) && customMsgArr[fk] != "" {
					msg = customMsgArr[fk] // 采用自定义的错误信息输出
				} else {
					sli := strings.Split(mk, "|") // 采用默认的错误信息输出
					data := make(map[string]string)
					for sk, sv := range sli {
						data["$key"+strconv.Itoa(sk+1)+"$"] = sv
					}
					msg = utils.ParseEasyTemplate(msg, data)
				}
				errors = append(errors, msg)
				break // 进行下一个字段的验证
			}
		}
	}
	return errors
}

// value 必须存在
func ValidateRequired(value string) ValidateFunc {
	return func() (msg string) {
		if value == "" {
			return "$key1$ 必须存在"
		}
		return ""
	}
}

func ValidateMinLength(value string, min int) ValidateFunc {
	return func() (msg string) {
		if len(value) < min {
			return "$key1$ 的长度必须大于 " + strconv.Itoa(min)
		}
		return ""
	}
}

func ValidateMaxLength(value string, max int) ValidateFunc {
	return func() (msg string) {
		if len(value) > max {
			return "$key1$ 的长度必须小于 " + strconv.Itoa(max)
		}
		return ""
	}
}

func ValidateEqual(v1 string, v2 string) ValidateFunc {
	return func() (msg string) {
		if v1 != v2 {
			return "$key1$ 必须等于 $key2$"
		}
		return ""
	}
}

func ValidateEmail(value string) ValidateFunc {
	return func() (msg string) {
		pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
		reg := regexp.MustCompile(pattern)
		if !reg.MatchString(value) {
			return "$key1$ 邮箱格式错误"
		}
		return ""
	}
}

func VlidErrorMsg(c *gin.Context, valid *validation.Validation, code int) {
	m := make(map[string]interface{})
	for _, v := range valid.Errors {
		m[v.Key] = v.Message
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.CodeToMessage(code),
		"data": m,
	})
}

func ValidCaptchaKey(v *validation.Validation, cap_key string) {
	v.Required(cap_key, "cap_key").Message("唯一key不能为空")
}

func ValidCaptcha(v *validation.Validation, captcha string) {
	v.Required(captcha, "captcha").Message("验证码不能为空")
}

func UserCaptchaValid(v *validation.Validation, cap_key string, captcha string) {
	ValidCaptchaKey(v, cap_key)
	ValidCaptcha(v, captcha)
}
