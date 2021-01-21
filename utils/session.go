package utils

import (
	"github.com/gin-gonic/gin"
	sessions "github.com/tommy351/gin-sessions"
)

func SetSession(c *gin.Context, key, value string) {
	session := sessions.Get(c)
	session.Set(key, value)
	session.Save()
}

func GetSession(c *gin.Context, key string) string {
	session := sessions.Get(c)
	if value, ok := session.Get(key).(string); ok {
		return value
	}
	return ""
}

func DelSession(c *gin.Context, key string) {
	session := sessions.Get(c)
	session.Delete(key)
	session.Save()
}
