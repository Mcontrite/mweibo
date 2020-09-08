package router

import (
	"fmt"
	"html/template"
	ctr "mweibo/controller"
	"mweibo/middleware/logger"
	"mweibo/utils"
	"net/http"
	"strings"

	"mweibo/middleware/redis"

	limit "github.com/aviddiviner/gin-limit"

	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	sessions "github.com/tommy351/gin-sessions"
)

// func InitRouter() *gin.Engine {
// 	g := gin.Default()
// 	setSession(g)
// 	setTemplate(g)
// 	g.Use(setContext())
// 	g.NoRoute(ctr.Handle404)
// 	registerApi(g)
// 	return g
// }

func InitRouter() *gin.Engine {
	g := gin.New()
	//gin.SetMode(conf.ServerConfig.RunMode)
	gin.SetMode("debug")
	setMiddleware(g)
	setSession(g)
	setTemplate(g)

	//g.Use(setContext())
	g.NoRoute(ctr.Handle404)
	registerApi(g)
	return g
}

// func setSession(e *gin.Engine) {
// 	// config := conf.GetConfiguration()
// 	// store := cookie.NewStore([]byte(config.SessionKey))
// 	skey := "MWeiBoSession"
// 	store := cookie.NewStore([]byte(skey))
// 	store.Options(sessions.Options{
// 		HttpOnly: true,
// 		MaxAge:   7 * 24 * 60 * 60,
// 		Path:     "/",
// 	})
// 	e.Use(sessions.Sessions("gin-session", store))
// }

func setMiddleware(g *gin.Engine) {
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(Cors())
	g.Use(limit.MaxAllowed(100))
	g.Use(logger.LoggerToFile())
}

func setSession(g *gin.Engine) {
	store := sessions.NewCookieStore([]byte("MWeiBoSession"))
	g.Use(sessions.Middleware("mweibo_session", store))
}

// func setContext() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		session := sessions.Default(c)
// 		if id := session.Get(ctr.SESSION_KEY); id != nil {
// 			user, err := model.GetUserByID(id)
// 			if err == nil {
// 				c.Set(ctr.CONTEXT_USER_KEY, user)
// 			}
// 			// user, _ := model.GetUserByID(id)
// 			// c.Set(ctr.CONTEXT_USER_KEY, user)
// 		}
// 		c.Next()
// 	}
// }

// funcmap
// func setTemplate(e *gin.Engine) {
// 	e.Static("/static", "static")
// 	funcMap := template.FuncMap{}
// 	e.SetFuncMap(funcMap)
// 	e.LoadHTMLGlob("views/**/*")
// }

func setTemplate(g *gin.Engine) {
	funcMap := template.FuncMap{}
	g.SetFuncMap(funcMap)
	g.StaticFS("/static", http.Dir("./static"))
	g.StaticFS("/upload", http.Dir("./upload"))
	//g.LoadHTMLGlob("views/*/**/***")
	g.LoadHTMLGlob("views/**/*")
}

// func Auth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if user, _ := c.Get(ctr.CONTEXT_USER_KEY); user != nil {
// 			if _, ok := user.(*model.User); ok {
// 				c.Next()
// 				return
// 			}
// 		}
// 		c.HTML(http.StatusForbidden, "error/error.html", gin.H{
// 			"message": "You need to login !",
// 		})
// 		c.Abort()
// 	}
// }

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.GetSession(c, "islogin") != "1" {
			c.Redirect(301, "/login")
			return
		}
		c.Next()
	}
}

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.GetSession(c, "isadmin") != "1" {
			c.Redirect(301, "/")
			return
		}
		c.Next()
	}
}

// func Admin() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if user, _ := c.Get(ctr.CONTEXT_USER_KEY); user != nil {
// 			u, ok := user.(*model.User)
// 			if ok && u.IsAdmin {
// 				c.Next()
// 				return
// 			}
// 		}
// 		c.HTML(http.StatusForbidden, "error/error.html", gin.H{
// 			"message": "You are not administrator !",
// 		})
// 		c.Abort()
// 	}
// }

// func GetCaptcha(c *gin.Context) {
// 	session := sessions.Default(c)
// 	data := captcha.NewLen(4)
// 	session.Delete(ctr.SESSION_CAPTCHA)
// 	session.Set(ctr.SESSION_CAPTCHA, data)
// 	session.Save()
// 	captcha.WriteImage(c.Writer, data, 100, 40)
// }

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
