package controller

import (
	"mweibo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

const (
	SESSION_KEY      = "UserID"
	CONTEXT_USER_KEY = "User"
	SESSION_CAPTCHA  = "GinCaptcha"
)

func Home(c *gin.Context) {
	weibos, err := model.ListWeibosByTag("")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	policy := bluemonday.StrictPolicy()
	// num,_:=model.GetWeiboByTagID("")
	for _, v := range weibos {
		v.Tags, _ = model.ListTagsByWeiboID(strconv.FormatUint(uint64(v.ID), 10))
		v.Content = policy.Sanitize(string(blackfriday.Run([]byte(v.Content))))
	}
	user, _ := c.Get(CONTEXT_USER_KEY)
	tags, _ := model.ListTags()
	mvweibos, _ := model.ListMaxViewWeibos()
	//mcweibos, _ := model.ListMaxCommentWeibos()
	c.HTML(http.StatusOK, "home/home.html", gin.H{
		"user":          user,
		"weibos":        weibos,
		"tags":          tags,
		"maxViewWeibos": mvweibos,
		//"maxCommentWeibos": mcweibos,
		"path": c.Request.URL.Path,
	})
}

func Handle404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error/error.html", gin.H{
		"message": "404 not found...",
	})
	// c.HTML(http.StatusOK, "error/error.html", nil)
}
