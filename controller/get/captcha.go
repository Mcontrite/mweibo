package get

import (
	"mweibo/middleware/validate"
	"mweibo/utils"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetCapacha(c *gin.Context) {
	height, _ := strconv.Atoi(c.DefaultQuery("height", "60"))
	width, _ := strconv.Atoi(c.DefaultQuery("width", "200"))
	code := utils.SUCCESS
	data := make(map[string]interface{})
	cap_key, captcha_base64 := utils.CodeCaptchaCreate(height, width)
	data["cap_key"] = cap_key
	data["captcha_base64"] = captcha_base64
	utils.ResponseJSONOK(c, code, data)
}

func VerfiyCaptcha(c *gin.Context) {
	cap_key := c.DefaultPostForm("cap_key", "")
	captcha := c.DefaultPostForm("captcha", "")
	code := utils.SUCCESS
	data := make(map[string]interface{})
	valid := &validation.Validation{}
	validate.UserCaptchaValid(valid, cap_key, captcha)
	if valid.HasErrors() {
		code = utils.INVALID_PARAMS
		validate.VlidErrorMsg(c, valid, code)
		return
	}
	pass := utils.VerfiyCaptcha(cap_key, captcha)
	if !pass {
		code = utils.UNPASS
	}
	utils.ResponseJSONOK(c, code, data)
}
