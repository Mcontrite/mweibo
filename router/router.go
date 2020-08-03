package router

import (
	"html/template"
	"mweibo/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	SESSION_KEY      = "UserID"
	CONTEXT_USER_KEY = "User"
	SESSION_CAPTCHA  = "GinCaptcha"
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
				c.Set(CONTEXT_USER_KEY, user)
			}
		}
		c.Next()
	}
}

func setTemplate(g *gin.Engine) {
	g.Static("/static", "stactic")
	funcMap := template.FuncMap{}
	g.SetFuncMap(funcMap)
	g.LoadHTMLGlob("views/**/*")
}

func Handle404(c *gin.Context) {
	// c.HTML(http.StatusNotFound, "error/error.html", gin.H{
	// 	"message": "404 not found...",
	// })
	c.HTML(http.StatusOK, "error/error.html", nil)
}

func InitRouter() *gin.Engine {
	g := gin.Default()
	setSession(g)
	setTemplate(g)
	g.Use(setContext())
	g.NoRoute(Handle404)
	registerApis(g)
	return g
}