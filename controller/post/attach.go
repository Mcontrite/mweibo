package post

import (
	filemid "mweibo/middleware/file"
	"mweibo/model"
	"mweibo/utils"

	"mweibo/middleware/image"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pixgif = "data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg=="

func CkeditorUpload(c *gin.Context) {
	file, _ := c.FormFile("upload")
	userid := utils.GetSession(c, "userid")
	fileName := file.Filename
	if !image.CheckImageSize2(file) {
		c.JSON(200, gin.H{
			"fileName": fileName,
			"uploaded": 1,
			"url":      pixgif,
		})
		return
	}
	newFilename := filemid.MakeFileName(userid, fileName)
	filepath := "upload/weibo/" + userid
	filepath, err := filemid.CreatePathInToday(filepath)
	if err != nil {
		c.JSON(200, gin.H{
			"fileName": fileName,
			"uploaded": 1,
			"url":      pixgif,
		})
		return
	}
	fullName := filepath + "/" + newFilename
	c.SaveUploadedFile(file, fullName)
	c.JSON(200, gin.H{
		"fileName": fileName,
		"uploaded": 1,
		"url":      "/" + fullName,
	})
}

func UploadFile(c *gin.Context) {
	action := c.Query("action")
	uid := c.Query("uid")
	code := utils.SUCCESS
	file, _ := c.FormFile("upload")
	fileName := file.Filename
	newFilename := filemid.MakeFileName(uid, fileName)
	if !image.CheckImageSize2(file) {
		code = utils.ERROR_IMAGE_TOO_LARGE
		utils.ResponseJSONError(c, code)
		return
	}
	filepath := "upload/" + action + "/" + uid + "/"
	err := filemid.CreatePath(filepath)
	if err != nil {
		code = utils.ERROR_FILE_CREATE_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
	fullName := filepath + newFilename
	err = c.SaveUploadedFile(file, fullName)
	if err != nil {
		code = utils.ERROR_FILE_SAVE_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
	c.JSON(200, gin.H{"filename": "测试图", "filetype": 1, "url": "/" + fullName, "attatchid": 99})
}

func UploadAttach(c *gin.Context) {
	userid := utils.GetSession(c, "userid")
	code := utils.SUCCESS
	file, _ := c.FormFile("upload")
	fileName := file.Filename
	fileType := filemid.GetType(fileName)
	newFilename := filemid.MakeFileName(userid, fileName)
	if !image.CheckImageSize2(file) {
		code = utils.ERROR_IMAGE_TOO_LARGE
		utils.ResponseJSONError(c, code)
		return
	}
	filepath := "upload/attach/" + userid
	filepath, err := filemid.CreatePathInToday(filepath)
	if err != nil {
		code = utils.ERROR_FILE_CREATE_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
	fullName := filepath + "/" + newFilename
	err = c.SaveUploadedFile(file, fullName)
	if err != nil {
		code = utils.ERROR_FILE_SAVE_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
	data := map[string]interface{}{"orgfilename": fileName, "filetype": fileType, "url": fullName}
	utils.ResponseJSONOK(c, code, data)
}

func UploadCreateAttach(c *gin.Context) {
	userid := utils.GetSession(c, "userid")
	weiboId, _ := strconv.Atoi(c.DefaultPostForm("weibo_id", "0"))
	posweiboId := weiboId
	commentId, _ := strconv.Atoi(c.PostForm("comment_id"))
	code := utils.SUCCESS
	file, _ := c.FormFile("upload")
	fileName := file.Filename
	fileType := filemid.GetType(fileName)
	fileSize := file.Size
	newFilename := filemid.MakeFileName(userid, fileName)
	if !image.CheckImageSize2(file) {
		code = utils.ERROR_IMAGE_TOO_LARGE
		utils.ResponseJSONError(c, code)
		return
	}
	filepath := "upload/attach/" + userid
	filepath, err := filemid.CreatePathInToday(filepath)
	if err != nil {
		code = utils.ERROR_FILE_CREATE_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
	fullName := filepath + "/" + newFilename
	err = c.SaveUploadedFile(file, fullName)
	if err != nil {
		code = utils.ERROR_FILE_SAVE_FAIL
		utils.ResponseJSONError(c, code)
		return
	}
	commentInfo, _ := model.GetCommentByID(commentId)
	if weiboId == 0 {
		weiboId = int(commentInfo.WeiboID)
	}
	useridInt, _ := strconv.Atoi(userid)
	model.CreateAttach(&model.Attach{
		WeiboID:   weiboId,
		CommentID: commentId,
		UserID:    useridInt,
		Filesize:  int(fileSize),
		Filename:  fullName,
		OrigiName: fileName,
		Filetype:  fileType,
	})
	if posweiboId != 0 {
		weiboInfo, _ := model.GetWeiboByID(weiboId)
		model.UpdateWeiboAttachsCnt(weiboId, weiboInfo.AttachsCnt+1)
	}
	model.UpdateCommentAttachsCnt(commentId, commentInfo.AttachsCnt+1)
	data := map[string]interface{}{"orgfilename": fileName, "filetype": fileType, "url": fullName}
	utils.ResponseJSONOK(c, code, data)
}

func DeleteAttach(c *gin.Context) {
	userid := utils.GetSession(c, "userid")
	_ = userid
	attachId, _ := strconv.Atoi(c.PostForm("attach_id"))
	weiboId, _ := strconv.Atoi(c.DefaultPostForm("weibo_id", "0"))
	commentId, _ := strconv.Atoi(c.DefaultPostForm("comment_id", "0"))
	code := utils.SUCCESS
	if weiboId != 0 {
		weiboInfo, _ := model.GetWeiboByID(weiboId)
		if weiboInfo.AttachsCnt != 0 {
			model.UpdateWeiboAttachsCnt(weiboId, weiboInfo.AttachsCnt-1)
		}
		//commentId = weiboInfo.FirstCommentID
	}
	if commentId != 0 {
		commentInfo, _ := model.GetCommentByID(commentId)
		if commentInfo.AttachsCnt != 0 {
			model.UpdateCommentAttachsCnt(commentId, commentInfo.AttachsCnt-1)
		}
	}
	model.DeleteAttachByID(attachId)
	utils.ResponseJSONOK(c, code, nil)
}
