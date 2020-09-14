package flash

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 存储上次表单 post 的数据
func SaveOldFormValue(c *gin.Context, m map[string]string) {
	flash := CreateFlashByName("oldForm")
	flash.Data = m
	flash.save(c, "oldForm")
}

// 读取上次表单 post 的数据
func ReadOldFormValue(c *gin.Context) *Flash {
	return read(c, "oldForm")
}

// 存储表单提交时数据的中间件
func SaveOldForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			req := c.Request
			if req.Form == nil {
				req.ParseForm()
			}
			oldForm := make(map[string]string)
			for k, v := range req.Form {
				oldForm[k] = v[0]
			}
			SaveOldFormValue(c, oldForm)
		}
		c.Next()
	}
}

// 存储参数验证的错误信息
func SaveValidateMessage(c *gin.Context, sli []string) {
	flash := CreateFlashByName("validateMessage")
	flash.Data = map[string]string{"errors": strings.Join(sli, "$$|$$")}
	flash.save(c, "validateMessage")
}

// 读取参数验证的错误信息
func ReadValidateMessage(c *gin.Context) []string {
	errstr := read(c, "validateMessage").Data["errors"]
	if errstr == "" {
		return []string{}
	}
	// 不做上面的判断，Split 切分空字符串会得 [""]
	return strings.Split(errstr, "$$|$$")
}
