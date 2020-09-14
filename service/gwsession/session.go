package gwsession

import (
	"errors"
	"mweibo/conf"
	"mweibo/model"
	"mweibo/utils"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	rememberFormKey    = "remember"
	rememberCookieName = "remember_me"
	rememberMaxAge     = 88888888 // 过期时间
)

// 记住我
func setRememberMeTokenInCookie(c *gin.Context, user *model.User) {
	rememberMe := c.PostForm(rememberFormKey) == "on"
	if !rememberMe {
		return
	}
	// 更新用户的 RememberMeToken
	token := string(utils.CreateRandomBytes(10))
	user.RememberMeToken = token
	if err := user.UpdateUser(); err != nil {
		return
	}
	c.SetCookie(rememberCookieName, user.RememberMeToken, rememberMaxAge, "/", "", false, true)
}

func getRememberMeTokenFromCookie(c *gin.Context) string {
	if cookie, err := c.Request.Cookie(rememberCookieName); err == nil {
		if v, err := url.QueryUnescape(cookie.Value); err == nil {
			return v
		}
	}
	return ""
}

func deleteRememberMeToken(c *gin.Context) {
	c.SetCookie(rememberCookieName, "", -1, "/", "", false, true)
}

func LoginSession(c *gin.Context, user *model.User) {
	utils.SetSession(c, conf.Serverconfig.SessionKey, strconv.Itoa(int(user.ID)))
	setRememberMeTokenInCookie(c, user)
}

func LogoutSession(c *gin.Context) {
	utils.DeleteSession(c, conf.Serverconfig.SessionKey)
	deleteRememberMeToken(c)
}

func getCurrentUserFromSession(c *gin.Context) (*model.User, error) {
	rememberMeToken := getRememberMeTokenFromCookie(c)
	if rememberMeToken != "" {
		if user, err := model.GetUserByRememberMeToken(rememberMeToken); err == nil {
			LoginSession(c, user)
			return user, nil
		}
		deleteRememberMeToken(c)
	}
	idstr := utils.GetSession(c, conf.Serverconfig.SessionKey)
	if idstr == "" {
		return nil, errors.New("没有获取到 session")
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return nil, err
	}
	user, err := model.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 保存用户数据到 context 中
func SaveCurrentUserToContext(c *gin.Context) {
	user, err := getCurrentUserFromSession(c)
	if err != nil {
		return
	}
	c.Keys[conf.Serverconfig.ContextUserKey] = user
}

// 从 session 中获取 user
func GetUserFromSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		SaveCurrentUserToContext(c)
		c.Next()
	}
}

// 从 context 中获取用户
func GetUserFromContext(c *gin.Context) (*model.User, error) {
	err := errors.New("没有获取到用户数据")
	ctxUserKey := c.Keys[conf.Serverconfig.ContextUserKey]
	if ctxUserKey == nil {
		return nil, err
	}
	user, ok := ctxUserKey.(*model.User)
	if !ok {
		return nil, err
	}
	return user, nil
}

// 从 context 或者数据库中获取用户
func GetUserFromContextOrDataBase(c *gin.Context, id int) (*model.User, error) {
	// 当前用户存在并且就是想要获取的那个用户
	ctxuser, err := GetUserFromContext(c)
	if ctxuser != nil && err == nil {
		if int(ctxuser.ID) == id {
			return ctxuser, nil
		}
	}
	// 获取的是其他指定 id 的用户
	otherUser, err := model.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return otherUser, nil
}
