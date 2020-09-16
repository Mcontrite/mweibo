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

func CreateCommentPOST(c *gin.Context) {
	weiboid, _ := strconv.Atoi(c.DefaultPostForm("weiboid", "1"))
	// docutype, _ := strconv.Atoi(c.DefaultPostForm("doctuype", "0"))
	content := c.DefaultPostForm("content", "")
	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	// uip := c.ClientIP()
	code := utils.SUCCESS
	attachFileString := c.PostForm("attachfiles")
	attachfiles := []string{}
	attachsCnt := 0
	if len(attachFileString) > 0 {
		attachfiles = strings.Split(attachFileString, ",")
		attachsCnt = len(attachfiles)
	}
	comment := &model.Comment{
		WeiboID:    uint(weiboid),
		UserID:     uint(uid),
		Content:    content,
		AttachsCnt: attachsCnt,
		// Isfirst:    0,
		// UserIP:     uip,
		// Doctype:    docutype,
		// Message:    message,
		// MessageFmt: message,

	}
	newComment, err := model.NewComment(comment)
	if err != nil {
		logging.Info("回复入库错误", err.Error())
		code = utils.ERROR_SQL_INSERT_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
	if len(attachFileString) > 0 {
		for _, attachfile := range attachfiles {
			fileSli := strings.Split(attachfile, "|")
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
				WeiboID:   int(weiboid),
				CommentID: int(newComment.ID),
				UserID:    uid,
				Filename:  fname,
				OrigiName: origname,
				Filetype:  ftype,
				Filesize:  fsize,
			}
			model.CreateAttach(attach)
		}
	}
	//comment_service.AfterAddNewComment(newComment, weiboid)
	utils.ResponseJSONOK(c, code, nil)
}

func UpdateCommentPOST(c *gin.Context) {
	comment_id, _ := strconv.Atoi(c.DefaultPostForm("comment_id", "1"))
	//doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	content := c.DefaultPostForm("content", "")
	uid, _ := strconv.Atoi(utils.GetSession(c, "userid"))
	//uip := c.ClientIP()
	code := utils.SUCCESS
	oldComment, err := model.GetCommentByID(comment_id)
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
	model.UpdateCommentByID(comment_id, comment)
	utils.ResponseJSONOK(c, code, nil)
}

// func CommentDelete(c *gin.Context) {
// 	var (
// 		err error
// 		res = gin.H{}
// 		cid uint64
// 	)
// 	defer writeJSON(c, res)
// 	session := sessions.Default(c)
// 	sessionUserId := session.Get(SESSION_KEY)
// 	userId := sessionUserId.(uint)
// 	commentId := c.Param("id")
// 	cid, err = strconv.ParseUint(commentId, 10, 64)
// 	if err != nil {
// 		res["message"] = err.Error()
// 		return
// 	}
// 	comment := &model.Comment{
// 		UserID: uint(userId),
// 	}
// 	comment.ID = uint(cid)
// 	err = service.DeleteCommentId(comment)
// 	if err != nil {
// 		res["message"] = err.Error()
// 		return
// 	}
// 	res["succeed"] = true
// }

func ReadComment(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	commentid := c.Param("id")
	comid, _ := strconv.ParseUint(commentid, 10, 64)
	comment := &model.Comment{}
	comment.ID = uint(comid)
	err := model.SetCommentRead(comment)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func ReadAllComments(c *gin.Context) {
	res := gin.H{}
	defer writeJSON(c, res)
	err := model.SetAllCommentsRead()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}
