package route

import (
	"html/template"
	"mweibo2/conf"
	ctrget "mweibo2/controller/get"
	"net/http"

	"github.com/gin-gonic/gin"
	sessions "github.com/tommy351/gin-sessions"
)

func InitRoute() *gin.Engine {
	e := gin.New()
	gin.SetMode(conf.ServerConf.ServerRunmode)
	setMiddleware(e)
	setSession(e)
	setTemplate(e)
	e.NoRoute(ctrget.ErrorPage)
	initAPI(e)
	return e
}

func setMiddleware(e *gin.Engine) {
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
}

func setSession(e *gin.Engine) {
	store := sessions.NewCookieStore([]byte("Session1"))
	e.Use(sessions.Middleware("Session2", store))
}

func setTemplate(e *gin.Engine) {
	funcmap := template.FuncMap{}
	e.SetFuncMap(funcmap)
	e.StaticFS("/static", http.Dir("./static"))
	e.LoadHTMLGlob("view/**/*")
	// e.LoadHTMLGlob("view/*")
}
