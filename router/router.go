package router

import (
	"html/template"
	ctr "mweibo/controller"
	"mweibo/model"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setSession(g *gin.Engine) {
	// config := conf.GetConfiguration()
	// store := cookie.NewStore([]byte(config.SessionKey))
	skey := "MWeiBoSession"
	store := cookie.NewStore([]byte(skey))
	store.Options(sessions.Options{
		HttpOnly: true,
		MaxAge:   7 * 24 * 60 * 60,
		Path:     "/",
	})
	g.Use(sessions.Sessions("gin-session", store))
}

func setContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if id := session.Get("UserID"); id != nil {
			user, err := model.GetUserByID(id)
			if err == nil {
				c.Set(ctr.CONTEXT_USER_KEY, user)
			}
		}
		c.Next()
	}
}

func setTemplate(g *gin.Engine) {
	g.Static("/static", "static")
	funcMap := template.FuncMap{}
	g.SetFuncMap(funcMap)
	g.LoadHTMLGlob("views/**/*")
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get(ctr.CONTEXT_USER_KEY); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		c.HTML(http.StatusForbidden, "error/error.html", gin.H{
			"message": "You need to login !",
		})
		c.Abort()
	}
}

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get(ctr.CONTEXT_USER_KEY); user != nil {
			u, ok := user.(*model.User)
			if ok && u.IsAdmin {
				c.Next()
				return
			}
		}
		c.HTML(http.StatusForbidden, "error/error.html", gin.H{
			"message": "You are not administrator !",
		})
		c.Abort()
	}
}

func GetCaptcha(c *gin.Context) {
	session := sessions.Default(c)
	data := captcha.NewLen(4)
	session.Delete(ctr.SESSION_CAPTCHA)
	session.Set(ctr.SESSION_CAPTCHA, data)
	session.Save()
	captcha.WriteImage(c.Writer, data, 100, 40)
}

func InitRouter() *gin.Engine {
	g := gin.Default()
	setSession(g)
	setTemplate(g)
	g.Use(setContext())
	g.NoRoute(ctr.Handle404)
	registerApis(g)
	return g
}
