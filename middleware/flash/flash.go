package flash // 模仿 beego flash 的实现

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type Flash struct {
	FlashKey string
	Data     map[string]string
}

func CreateFlash() *Flash {
	return &Flash{
		FlashKey: "flash",
		Data:     make(map[string]string),
	}
}

func CreateFlashByName(flashkey string) *Flash {
	return &Flash{
		FlashKey: flashkey,
		Data:     make(map[string]string),
	}
}

// 将 flash 数据保存到 gin context keys 和 cookie 中
func (flash *Flash) save(c *gin.Context, flashkey string) {
	c.Keys[flashkey] = flash.Data
	var flashValue string
	for k, v := range flash.Data {
		flashValue += "\x00" + k + "\x23" + "|" + "\x23" + v + "\x00"
	}
	c.SetCookie(flashkey, flashValue, 0, "/", "", false, true)
}

func (flash *Flash) Save(c *gin.Context) {
	flash.save(c, "flash")
}

// 从 request 中的 cookie 里解析出 flash 数据
func read(c *gin.Context, flashkey string) *Flash {
	flash := CreateFlashByName(flashkey)
	if cookie, err := c.Request.Cookie(flashkey); err == nil {
		v, _ := url.QueryUnescape(cookie.Value)
		vals := strings.Split(v, "\x00")
		for _, v := range vals {
			if len(v) > 0 {
				kv := strings.Split(v, "\x23"+"|"+"\x23")
				if len(kv) == 2 {
					flash.Data[kv[0]] = kv[1]
				}
			}
		}
		// 读取 flash 时先从 session 中 delete 再 save
		c.SetCookie(flashkey, "", -1, "/", "", false, true)
	}
	c.Keys[flashkey] = flash.Data
	return flash
}

// 从 request 中的 cookie 里解析出 flash 数据
func Read(c *gin.Context) *Flash {
	return read(c, "flash")
}

func (flash *Flash) Set(key string, msg string, args ...interface{}) {
	if len(args) == 0 {
		flash.Data[key] = msg
	} else {
		flash.Data[key] = fmt.Sprintf(msg, args...)
	}
}
