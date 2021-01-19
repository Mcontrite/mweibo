package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home/home.html", nil)
}

func ErrorPage(c *gin.Context) {
	c.HTML(http.StatusNotFound, "home/error.html", nil)
}
