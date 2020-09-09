package post

import (
	"mweibo/middleware/logging"
	"mweibo/model"
	"mweibo/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateWeiboPOST(c *gin.Context) {
	// doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	//title := c.DefaultPostForm("title", "")
	// message := c.DefaultPostForm("message", "")
	// attachFileString := c.PostForm("attachfiles")
	// attachfiles := []string{}
	// filesNum := 0
	content := c.PostForm("content")
	tags := c.PostForm("tags")
	code := utils.SUCCESS
	// if len(attachFileString) > 0 {
	// 	attachfiles = strings.Split(attachFileString, ",")
	// 	filesNum = len(attachfiles)
	// }
	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	// uip := c.ClientIP()
	weibo := &model.Weibo{
		UserID: uint(uid),
		// UserIP:   uip,
		Content: content,
		// FilesNum: filesNum,
		// LastDate: time.Now(),
	}
	err := model.CreateWeibo(weibo)
	if err != nil {
		logging.Info("weibo入库错误", err.Error())
		code = utils.ERROR_SQL_INSERT_FAIL
		utils.ResponseJSONError(c, code)
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
	// comment := &model.Comment{
	// 	WeiboID:  int(newWeibo.ID),
	// 	UserID:     uid,
	// 	Isfirst:    1,
	// 	UserIP:     uip,
	// 	Doctype:    doctype,
	// 	Message:    message,
	// 	MessageFmt: message,
	// }
	// newComment, err := model.AddComment(comment)
	// if err != nil {
	// 	logging.Info("comment入库错误", err.Error())
	// 	code = utils.ERROR
	// 	code = utils.ERROR_SQL_INSERT_FAIL
	// 	utils.ResponseJSONError(c, code)
	// 	return
	// }
	// model.UpdateWeibo(int(newWeibo.ID), model.Weibo{FirstCommentID: int(newComment.ID), LastDate: time.Now()})
	// weibo_service.AfterAddNewWeibo(newWeibo)
	// if len(attachFileString) > 0 {
	// 	for _, attachfile := range attachfiles {
	// 		file := strings.Split(attachfile, "|")
	// 		fname := file[0]
	// 		forginname := file[1]
	// 		ftype := file_package.GetType(fname)
	// 		ofile, err := os.Open(fname)
	// 		defer ofile.Close()
	// 		if err != nil {
	// 			continue
	// 		}
	// 		fsize, _ := file_package.GetSize(ofile)
	// 		attach := &model.Attach{
	// 			WeiboID:   int(newWeibo.ID),
	// 			CommentID:     int(newComment.ID),
	// 			UserID:      uid,
	// 			Filename:    fname,
	// 			Orgfilename: forginname,
	// 			Filetype:    ftype,
	// 			Filesize:    fsize,
	// 		}
	// 		_, err = model.AddAttach(attach)
	// 		if err != nil {
	// 			logging.Info("attach入库错误", err.Error())
	// 			code = utils.ERROR_SQL_INSERT_FAIL
	// 			utils.ResponseJSONError(c, code)
	// 			return
	// 		}
	// 	}
	// }
	utils.ResponseJSONOK(c, code, nil)
	//c.Redirect(http.StatusMovedPermanently, "/")
}

// type Tids struct {
// 	Tidarr []string `json:"tidarr"`
// }
// func DeleteWeibos(c *gin.Context) {
// 	ids := c.PostForm("tidarr")
// 	code := utils.SUCCESS
// 	idsSlice := strings.Split(ids, ",")
// 	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
// 	// isadmin := user_service.IsAdmin(uid)
// 	// if isadmin == "0" {
// 	// 	code = utils.UNPASS
// 	// 	utils.ResponseJSONError(c, code)
// 	// 	return
// 	// }
// 	err := weibo_service.DelWeibos(idsSlice)
// 	if err != nil {
// 		code = utils.ERROR
// 		utils.ResponseJSONError(c, code)
// 		return
// 	}
// 	utils.ResponseJSONOK(c, code, ids)
// }
func DeleteWeibo(c *gin.Context) {
	weiboid := c.Param("id")
	weiid, _ := strconv.ParseUint(weiboid, 10, 64)
	code := utils.SUCCESS
	weibo := &model.Weibo{}
	weibo.ID = uint(weiid)
	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	oldWeibo, _ := model.GetWeiboByID(int(weiid))
	if oldWeibo.UserID != uint(uid) {
		code = utils.UNPASS
		utils.ResponseJSONError(c, code)
		return
	}
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
	// comment_id, _ := strconv.Atoi(c.DefaultPostForm("comment_id", "1"))
	// doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	// title := c.DefaultPostForm("title", "")
	// message := c.DefaultPostForm("message", "")
	content := c.PostForm("content")
	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	// uip := c.ClientIP()
	code := utils.SUCCESS
	oldWeibo, err := model.GetWeiboByID(weibo_id)
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
	// oldComment, err := model.GetWeiboFirstCommentByTid(weibo_id)
	// if err != nil {
	// 	code = utils.ERROR_UNFIND_DATA
	// 	utils.ResponseJSONError(c, code)
	// 	return
	// }
	// if int(oldComment.ID) != comment_id {
	// 	code = utils.UNPASS
	// 	utils.ResponseJSONError(c, code)
	// 	return
	// }
	weibo := &model.Weibo{
		// UserIP: uip,
		Content: content,
	}
	model.UpdateWeibo(weibo, c.Param("id"))
	// comment := model.Comment{
	// 	UserIP:  uip,
	// 	Doctype: doctype,
	// 	Message: message,
	// }
	// model.UpdateComment(comment_id, comment)
	utils.ResponseJSONOK(c, code, nil)
}

// // // 添加附件
// // // 直接添加到表中，因为以及各有了文章  所以可以直接添加
// // func AddweiboAttach(c *gin.Context) {
// // 	// 获取文件内容
// // 	// 获取weiboid commentid uid
// // 	// 修改weibo表的files字段 + 1
// // 	// 在attach表中添加一天新的记录
// // }

// // // 删除的附件  知己额删除  提供好attach的id  就能删除
// // func DelweiboAttach(c *gin.Context) {
// // 	// 删除数据内容  删除文件内容
// // 	// 获取weiboid
// // 	// 修改weibo表的files字段 - 1
// // 	// 在attach表中直接删除记录
// // }
