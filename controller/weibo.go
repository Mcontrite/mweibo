package controller

import (
	"mweibo/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateWeiboGet(c *gin.Context) {
	c.HTML(http.StatusOK, "weibo/create.html", nil)
}

func CreateWeiboPost(c *gin.Context) {
	content := c.PostForm("content")
	tags := c.PostForm("tags")
	weibo := &model.Weibo{
		Content: content,
	}
	err := model.CreateWeibo(weibo)
	if err != nil {
		c.HTML(http.StatusOK, "weibo/create.html", gin.H{
			"weibo":   weibo,
			"message": err.Error(),
		})
		return
	}
	if tags != "" {
		sli := strings.Split(tags, ",")
		for _, v := range sli {
			tagid, _ := strconv.ParseUint(v, 10, 64)
			tw := &model.TagWeibo{
				WeiboID: weibo.ID,
				TagID:   uint(tagid),
			}
			err = model.CreateTagWeibo(tw)
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
