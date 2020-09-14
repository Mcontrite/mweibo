package controller

import (
	"fmt"
	"html/template"
	"mweibo/conf"
	"mweibo/middleware/flash"
	gwservice "mweibo/service/gwsession"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func csrfField(c *gin.Context) (template.HTML, string, bool) {
	token := c.Keys[conf.Serverconfig.CsrfParamName]
	tokenstr, ok := token.(string)
	if !ok {
		return "", "", false
	}
	return template.HTML(fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`, conf.Serverconfig.CsrfParamName, tokenstr)), tokenstr, true
}

// 渲染 html
func Render(c *gin.Context, tplPath string, data map[string]interface{}) {
	m := make(map[string]interface{})
	m["flash"] = flash.Read(c).Data                     // flash 数据
	m["oldForm"] = flash.ReadOldFormValue(c).Data       // 上次表单的数据
	m["validateMessage"] = flash.ReadValidateMessage(c) // 上次表单的验证信息
	if conf.Serverconfig.EnableCsrf {
		if csrfHTML, csrfToken, ok := csrfField(c); ok {
			m["csrfHTML"] = csrfHTML
			m["csrfToken"] = csrfToken
		}
	}
	// 获取当前登录的用户
	if user, err := gwservice.GetUserFromContext(c); err == nil {
		//m[conf.Serverconfig.ContextUserKey] = viewmodels.NewUserViewModelSerializer(user)
		m[conf.Serverconfig.ContextUserKey] = user
	}
	// 填充传递进来的数据
	for k, v := range data {
		m[k] = v
	}
	c.HTML(http.StatusOK, tplPath, m)
}

// 渲染错误页面
func RenderError(c *gin.Context, code int, msg string) {
	errorCode := code
	if code == 419 || code == 403 {
		errorCode = 403
	}
	c.HTML(code, "error/error.html", gin.H{
		"errorMsg":  msg,
		"errorCode": errorCode,
		"errorImg":  "/static/svg/" + strconv.Itoa(code) + ".svg",
		"backUrl":   "/",
		//"backUrl":   named.G("root"),
	})
}

func Render403(c *gin.Context, msg string) {
	RenderError(c, http.StatusForbidden, msg)
}

func Render404(c *gin.Context) {
	RenderError(c, http.StatusNotFound, "很抱歉！您浏览的页面不存在。")
}

func RenderUnauthorized(c *gin.Context) {
	Render403(c, "很抱歉，您没有权限访问该页面")
}
