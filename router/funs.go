package router

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
	"mweibo/middleware/redis"
	"mweibo/model"
	gwservice "mweibo/service/gwsession"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthUser(handler func(*gin.Context, *model.User)) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 用户未登录则跳转到登录页
		ctxuser, err := gwservice.GetUserFromContext(c)
		if ctxuser == nil || err != nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		handler(c, ctxuser)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}

func XSS() gin.HandlerFunc {
	return func(c *gin.Context) {
		xssToken := c.DefaultPostForm("xss_token", "")
		if len(xssToken) == 0 {
			c.JSON(200, gin.H{
				"code":    401,
				"message": "请提交xsstoken",
			})
			c.Abort()
			return
		}
		_, err := redis.Get(xssToken)
		if err == nil {
			c.JSON(200, gin.H{
				"code":    403,
				"message": "已经提交过了，不要重复提交",
			})
			c.Abort()
			return
		}
		redis.Set(xssToken, xssToken, 100)
		c.Next()
	}
}

// import (
// 	"ginweibo/controllers"
// 	"ginweibo/middleware/auth"
// 	"ginweibo/middleware/flash"
// 	userModel "ginweibo/models/user"

// 	"github.com/gin-gonic/gin"
// )

// type (
// 	AuthHandlerFunc = func(*gin.Context, *userModel.User)
// )

// func Guest(handler gin.HandlerFunc) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// 用户已经登录了则跳转到 root page
// 		ctxuser, err := auth.GetUserFromContext(c)
// 		if ctxuser != nil || err == nil {
// 			flash.NewInfoFlash(c, "您已登录，无需再次操作。")
// 			controllers.RedirectRouter(c, "root")
// 			return
// 		}
// 		handler(c)
// 	}
// }

///////////////////////////////////////FuncMap/////////////////////////////////////////

func selfPlus(num int) int {
	return num + 1
}

//避免模板某些字段自动转义实体
func unescaped(x string) interface{} {
	return template.HTML(x)
}

func StrTime(atime int64) string {
	var byTime = []int64{365 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}
	var unit = []string{"年前", "天前", "小时前", "分钟前", "秒钟前"}
	now := time.Now().Unix()
	ct := now - atime
	if ct < 0 {
		return "刚刚"
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = MergeString(tempStr, unit[i])
		}
		break
	}
	return res
}

func MergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
	}
	return buffer.String()
}

func numPlusPlus(num int) int {
	num++
	return num
}

func Long2IPString(si string) string {
	i, _ := strconv.Atoi(si)
	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)
	return ip.String()
}

// 截取字符串
func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) > n {
		return string(runes[:n])
	}
	return s
}
