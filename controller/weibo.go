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

func UpdateWeiboGet(c *gin.Context) {
	weiboid := c.Param("id")
	weibo, err := model.GetWeiboByID(weiboid)
	if err != nil {
		return
	}
	weibo.Tags, _ = model.ListTagsByWeiboID(weiboid)
	c.HTML(http.StatusOK, "weibo/update.html", gin.H{
		"weibo": weibo,
	})
}

func UpdateWeibo(c *gin.Context) {
	weiboid := c.Param("id")
	content := c.PostForm("content")
	tags := c.PostForm("tags")
	weiid, _ := strconv.ParseUint(weiboid, 10, 64)
	weibo := &model.Weibo{
		Content: content,
	}
	err := model.UpdateWeibo(weibo, weiboid)
	if err != nil {
		c.HTML(http.StatusOK, "weibo/update.html", gin.H{
			"weibo":   weibo,
			"message": err.Error(),
		})
		return
	}
	model.DeleteTagWeiboByWeiboID(weiid)
	if len(tags) > 0 {
		sli := strings.Split(tags, ",")
		for _, v := range sli {
			tagid, _ := strconv.ParseUint(v, 10, 64)
			tw := &model.TagWeibo{
				WeiboID: uint(weiid),
				TagID:   uint(tagid),
			}
			err = model.CreateTagWeibo(tw)
			if err != nil {
				return
			}
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func DeleteWeibo(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	weiboid := c.Param("id")
	weiid, _ := strconv.ParseUint(weiboid, 10, 64)
	weibo := &model.Weibo{}
	weibo.ID = uint(weiid)
	err := model.DeleteWeibo(weibo)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	err = model.DeleteTagWeiboByWeiboID(weiid)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["secceed"] = true
}

func DisplayWeibo(c *gin.Context) {
	res := gin.H{}
	weiboid := c.Param("id")
	weibo, err := model.GetWeiboByID(weiboid)
	if err != nil {
		Handle404(c)
		return
	}
	weibo.ViewsCnt++
	err = model.UpdateWeiboViewsCnt(weibo)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	weibo.Tags, _ = model.ListTagsByWeiboID(weiboid)
	weibo.Comments, _ = model.ListCommentsByWeiboID(weiboid)
	user, _ := c.Get(CONTEXT_USER_KEY)
	c.HTML(http.StatusOK, "weibo/display.html", gin.H{
		"user":  user,
		"weibo": weibo,
	})
}

func ListWeibos(c *gin.Context) {
	user, _ := c.Get(CONTEXT_USER_KEY)
	weibos, _ := model.ListWeibos("")
	c.HTML(http.StatusOK, "admin/weibos.html", gin.H{
		"user":   user,
		"weibos": weibos,
	})
}
