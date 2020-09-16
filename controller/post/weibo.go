package post

import (
	"mweibo/middleware/file"
	"mweibo/middleware/logging"
	"mweibo/model"
	"mweibo/utils"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateWeiboPOST(c *gin.Context) {
	// doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	//title := c.DefaultPostForm("title", "")
	// message := c.DefaultPostForm("message", "")

	content := c.PostForm("content")
	tags := c.PostForm("tags")
	code := utils.SUCCESS

	attachFileString := c.PostForm("attachfiles")
	attachfiles := []string{}
	attachsCnt := 0
	if len(attachFileString) > 0 {
		attachfiles = strings.Split(attachFileString, ",")
		attachsCnt = len(attachfiles)
	}

	// 微博内容、长度验证
	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	// uip := c.ClientIP()
	weibo := &model.Weibo{
		UserID: uint(uid),
		// UserIP:   uip,
		Content:    content,
		AttachsCnt: attachsCnt,
		// LastDate: time.Now(),
	}
	newWeibo, err := model.NewWeibo(weibo)
	if err != nil {
		logging.Info("weibo入库错误", err.Error())
		code = utils.ERROR_SQL_INSERT_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
	//article_service.AfterAddNewWeibo(newWeibo)
	if len(attachFileString) > 0 {
		for _, v := range attachfiles {
			fileSli := strings.Split(v, "|")
			fname := fileSli[0]
			origname := fileSli[1]
			ftype := file.GetType(fname)
			ofile, err := os.Open(fname)
			defer ofile.Close()
			if err != nil {
				continue
			}
			fsize, _ := file.GetSize(ofile)
			attach := &model.Attach{
				WeiboID: int(newWeibo.ID),
				//ReplyID:     int(new.ID),
				UserID:    uid,
				Filename:  fname,
				OrigiName: origname,
				Filetype:  ftype,
				Filesize:  fsize,
			}
			_, err = model.CreateAttach(attach)
			if err != nil {
				logging.Info("attach入库错误", err.Error())
				code = utils.ERROR_SQL_INSERT_FAIL
				utils.ResponseJSONError(c, code)
				return
			}
		}
	}
	if len(tags) > 0 {
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
	utils.ResponseJSONOK(c, code, nil)
	//c.Redirect(http.StatusMovedPermanently, "/")
}

// func DeleteWeibos(){}
func DeleteWeibo(c *gin.Context) {
	weiboid := c.Param("id")
	weiid, _ := strconv.ParseUint(weiboid, 10, 64)
	code := utils.SUCCESS
	weibo := &model.Weibo{}
	weibo.ID = uint(weiid)
	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	oldWeibo, _ := model.GetWeiboObjectByID(int(weiid))
	if oldWeibo.UserID != uint(uid) {
		code = utils.UNPASS
		utils.ResponseJSONError(c, code)
		return
	}
	// // 权限判断
	// if ok := policies.StatusPolicyDestroy(c, currentUser, status); !ok {
	// 	return
	// }
	err := model.DeleteWeibo(weibo)
	if err != nil {
		code = utils.ERROR
		utils.ResponseJSONError(c, code)
		return
	}
	err = model.DeleteTagWeiboByWeiboID(weiid)
	if err != nil {
		code = utils.ERROR
		utils.ResponseJSONError(c, code)
		return
	}
	utils.ResponseJSONOK(c, code, nil)
}

func UpdateWeiboPOST(c *gin.Context) {
	weibo_id, _ := strconv.Atoi(c.Param("id"))
	content := c.PostForm("content")
	tags := c.PostForm("tags")
	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	// uip := c.ClientIP()
	code := utils.SUCCESS
	oldWeibo, err := model.GetWeiboObjectByID(weibo_id)
	if err != nil {
		code = utils.ERROR_UNFIND_DATA
		utils.ResponseJSONError(c, code)
		return
	}
	if oldWeibo.UserID != uint(uid) {
		code = utils.UNPASS
		utils.ResponseJSONError(c, code)
		return
	}
	weibo := &model.Weibo{
		// UserIP: uip,
		Content: content,
	}
	model.UpdateWeibo(weibo, c.Param("id"))
	model.DeleteTagWeiboByWeiboID(weibo_id)
	if len(tags) > 0 {
		sli := strings.Split(tags, ",")
		for _, v := range sli {
			tagid, _ := strconv.ParseUint(v, 10, 64)
			tw := &model.TagWeibo{
				WeiboID: uint(weibo_id),
				TagID:   uint(tagid),
			}
			model.CreateTagWeibo(tw)
		}
	}
	utils.ResponseJSONOK(c, code, nil)
}

// // 添加附件
// // 直接添加到表中，因为以及各有了文章  所以可以直接添加
// func AddweiboAttach(c *gin.Context) {
// 	// 获取文件内容
// 	// 获取weiboid commentid uid
// 	// 修改weibo表的files字段 + 1
// 	// 在attach表中添加一天新的记录
// }

// // 删除的附件  知己额删除  提供好attach的id  就能删除
// func DelweiboAttach(c *gin.Context) {
// 	// 删除数据内容  删除文件内容
// 	// 获取weiboid
// 	// 修改weibo表的files字段 - 1
// 	// 在attach表中直接删除记录
// }
