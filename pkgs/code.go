package pkgs

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	UNPASS         = 403

	ERROR_EXIST_USER        = 10001
	ERROR_NOT_EXIST_USER    = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_BIND_DATA   = 30001
	ERROR_UNFIND_DATA = 30004

	ERROR_SQL_INSERT_FAIL = 40001
	ERROR_SQL_DELETE_FAIL = 40002
	ERROR_SQL_UPDATE_FAIL = 40003

	ERROR_IMAGE_BAD_EXT    = 50001
	ERROR_IMAGE_TOO_LARGE  = 50002
	ERROR_FILE_CREATE_FAIL = 50003
	ERROR_FILE_SAVE_FAIL   = 50004
)

var CodeMessage = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "发生错误",
	INVALID_PARAMS: "请求参数错误",
	UNPASS:         "未通过验证",

	ERROR_EXIST_USER:        "用户已存在",
	ERROR_NOT_EXIST_USER:    "用户名或密码错误",
	ERROR_NOT_EXIST_ARTICLE: "文章不存在",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_BIND_DATA:   "数据绑定错误",
	ERROR_UNFIND_DATA: "数据未找到",

	ERROR_SQL_INSERT_FAIL: "数据插入失败",
	ERROR_SQL_DELETE_FAIL: "数据删除失败",
	ERROR_SQL_UPDATE_FAIL: "数据修改失败",

	ERROR_IMAGE_BAD_EXT:    "图片格式不正确",
	ERROR_IMAGE_TOO_LARGE:  "图片太大了",
	ERROR_FILE_CREATE_FAIL: "文件夹创建失败",
	ERROR_FILE_SAVE_FAIL:   "文件保存失败",
}

func CodeToMessage(code int) string {
	msg, ok := CodeMessage[code]
	if ok {
		return msg
	}
	return CodeMessage[ERROR]
}
