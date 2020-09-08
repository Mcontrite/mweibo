package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseHTML(c *gin.Context, httpcode int, filepath string, data gin.H) {
	c.HTML(httpcode, filepath, data)
	return
}

func ResponseJSON(c *gin.Context, httpcode, errcode int, data interface{}) {
	c.JSON(httpcode, gin.H{
		"code":    errcode,
		"message": CodeToMessage(errcode),
		"data":    data,
	})
	return
}

func ResponseJSONOK(c *gin.Context, errcode int, data interface{}) {
	ResponseJSON(c, http.StatusOK, errcode, data)
}

func ResponseJSONError(c *gin.Context, errcode int) {
	ResponseJSON(c, http.StatusOK, errcode, nil)
}
