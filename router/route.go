package router

import (
	"html/template"
	"mweibo/conf"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setSession(g *gin.Engine) {
	config := conf.GetConfiguration()
	store := cookie.NewStore([]byte(config.SessionKey))
	store.Options(sessions.Options{
		HttpOnly: true,
		MaxAge:   7 * 24 * 60 * 60,
		Path:     "/",
	})
	g.Use(sessions.Sessions("gin-session", store))
}

func setTemplate(g *gin.Engine) {
	funcMap := template.FuncMap{}
	g.SetFuncMap(funcMap)
	g.LoadHtmlGlob()
}

func Handel404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error/error.html", gin.H{
		"message": "404 not found...",
	})
}

func InitRouter() *gin.Engine {
	g := gin.Default()
	setSession(g)
	//setTemplate(g)
	g.NoRoute(Handle404())
}
