package flash

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Success flash
func (flash *Flash) Success(msg string, args ...interface{}) {
	if len(args) == 0 {
		flash.Data["success"] = msg
	} else {
		flash.Data["success"] = fmt.Sprintf(msg, args...)
	}
}

// Info flash
func (flash *Flash) Info(msg string, args ...interface{}) {
	if len(args) == 0 {
		flash.Data["info"] = msg
	} else {
		flash.Data["info"] = fmt.Sprintf(msg, args...)
	}
}

// Warning flash
func (flash *Flash) Warning(msg string, args ...interface{}) {
	if len(args) == 0 {
		flash.Data["warning"] = msg
	} else {
		flash.Data["warning"] = fmt.Sprintf(msg, args...)
	}
}

// Danger falsh
func (flash *Flash) Danger(msg string, args ...interface{}) {
	if len(args) == 0 {
		flash.Data["danger"] = msg
	} else {
		flash.Data["danger"] = fmt.Sprintf(msg, args...)
	}
}

// 新建一条 success flash，并保存
func NewSuccessFlash(c *gin.Context, msg string, args ...interface{}) {
	f := CreateFlash()
	f.Success(msg, args...)
	f.Save(c)
}

// 新建一条 info flash，并保存
func NewInfoFlash(c *gin.Context, msg string, args ...interface{}) {
	f := CreateFlash()
	f.Info(msg, args...)
	f.Save(c)
}

// 新建一条 warning flash，并保存
func NewWarningFlash(c *gin.Context, msg string, args ...interface{}) {
	f := CreateFlash()
	f.Warning(msg, args...)
	f.Save(c)
}

// 新建一条 danger flash，并保存
func NewDangerFlash(c *gin.Context, msg string, args ...interface{}) {
	f := CreateFlash()
	f.Danger(msg, args...)
	f.Save(c)
}
