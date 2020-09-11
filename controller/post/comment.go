package post

import (
	"mweibo/middleware/logging"
	"mweibo/model"
	"mweibo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCommentPOST(c *gin.Context) {
	weiboid, _ := strconv.Atoi(c.DefaultPostForm("weiboid", "1"))
	// docutype, _ := strconv.Atoi(c.DefaultPostForm("doctuype", "0"))
	content := c.DefaultPostForm("content", "")
	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	// uip := c.ClientIP()
	code := utils.SUCCESS
	// attachFileString := c.PostForm("attachfiles")
	// attachfiles := []string{}
	// filesNum := 0
	// if len(attachFileString) > 0 {
	// 	attachfiles = strings.Split(attachFileString, ",")
	// 	filesNum = len(attachfiles)
	// }
	comment := &model.Comment{
		WeiboID: uint(weiboid),
		UserID:  uint(uid),
		Content: content,
		// Isfirst:    0,
		// UserIP:     uip,
		// Doctype:    docutype,
		// Message:    message,
		// MessageFmt: message,
		// FilesNum:   filesNum,
	}
	err := model.CreateComment(comment)
	//newComment, err := model.AddComment(comment)
	if err != nil {
		logging.Info("回复入库错误", err.Error())
		code = utils.ERROR_SQL_INSERT_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
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
	// 			WeiboID:   int(weiboid),
	// 			CommentID:     int(newComment.ID),
	// 			UserID:      uid,
	// 			Filename:    fname,
	// 			Orgfilename: forginname,
	// 			Filetype:    ftype,
	// 			Filesize:    fsize,
	// 		}
	// 		model.AddAttach(attach)
	// 	}
	// }
	//comment_service.AfterAddNewComment(newComment, weiboid)
	utils.ResponseJSONOK(c, code, nil)
}

func UpdateComment(c *gin.Context) {
	comment_id, _ := strconv.Atoi(c.DefaultPostForm("comment_id", "1"))
	//doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	content := c.DefaultPostForm("content", "")
	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	//uip := c.ClientIP()
	code := utils.SUCCESS
	oldComment, err := model.GetCommentById(comment_id)
	if err != nil {
		code = utils.ERROR_UNFIND_DATA
		utils.ResponseJSONError(c, code)
		return
	}
	if oldComment.UserID != uint(uid) {
		code = utils.UNPASS
		utils.ResponseJSONError(c, code)
		return
	}
	comment := model.Comment{
		// UserIP:  uip,
		// Doctype: doctype,
		Content: content,
	}
	model.UpdateComment(comment_id, comment)
	utils.ResponseJSONOK(c, code, nil)
}
