package post

import (
	"mweibo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func writeJSON(c *gin.Context, h gin.H) {
	if _, ok := h["succeed"]; !ok {
		h["succeed"] = false
	}
	c.JSON(http.StatusOK, h)
}

func CreateTag(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	tagname := c.PostForm("tagname")
	tag := &model.Tag{
		TagName: tagname,
	}
	err := model.CreateTag(tag)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["secceed"] = true
	res["data"] = tag
}

//*
func ListTags(c *gin.Context) {
	var (
		tagname string
		//num     int
		policy *bluemonday.Policy
		weibos []*model.Weibo
		err    error
	)
	tagname = c.Param("tagname")
	weibos, err = model.ListWeibosByTag(tagname)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//num, _ = model.CountWeibosByTag(tagname)
	policy = bluemonday.StrictPolicy()
	for _, v := range weibos {
		v.Tags, _ = model.ListTagsByWeiboID(strconv.FormatUint(uint64(v.ID), 10))
		v.Content = policy.Sanitize(string(blackfriday.Run([]byte(v.Content))))
	}
	tags, _ := model.ListTags()
	mvweibos, _ := model.ListMaxViewWeibos()
	//mcweibos, _ := model.ListMaxCommentWeibos()
	c.HTML(http.StatusOK, "home/home.html", gin.H{
		"weibos":        weibos,
		"tags":          tags,
		"maxViewWeibos": mvweibos,
		//"maxCommentWeibos": mcweibos,
	})
}
