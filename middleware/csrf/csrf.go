package csrf

import (
	"mweibo/conf"

	"mweibo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 从 cookie 中获取 csrf token，没有则设置
func getCsrfTokenFromCookie(c *gin.Context) (token string) {
	csrfparam := conf.Serverconfig.CsrfParamName
	if s, err := c.Request.Cookie(csrfparam); err == nil {
		token = s.Value
	}
	if token == "" {
		token = string(utils.CreateRandomBytes(32))
		c.SetCookie(csrfparam, token, 0, "/", "", false, false)
	}
	c.Keys[csrfparam] = token
	return
}

// 从 params 或 headers 中获取 csrf token
func getCsrfTokenFromParamsOrHeader(c *gin.Context) (token string) {
	req := c.Request
	if req.Form == nil {
		req.ParseForm()
	}
	token = req.FormValue(conf.Serverconfig.CsrfParamName)
	if token == "" {
		token = req.Header.Get(conf.Serverconfig.CsrfHeaderName)
	}
	return
}

func Csrf() gin.HandlerFunc {
	return func(c *gin.Context) {
		if conf.Serverconfig.EnableCsrf {
			csrfToken := getCsrfTokenFromCookie(c)
			if c.Request.Method == http.MethodPost {
				paramCsrfToken := getCsrfTokenFromParamsOrHeader(c)
				if paramCsrfToken == "" || paramCsrfToken != csrfToken {
					//controllers.Render403(c, "您的 Session 已过期，刷新后再试一次。")
					c.Abort()
					return
				}
			}
		}
		c.Next()
	}
}
