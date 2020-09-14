package controller

import (
	"math"
	"mweibo/conf"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func redirect(c *gin.Context, route string) {
	// 注意这个地方不能用 301 永久重定向
	c.Redirect(http.StatusFound, route)
}

// 路由重定向 use path
func Redirect(c *gin.Context, route string, withRoot bool) {
	path := route
	if withRoot {
		path = conf.Serverconfig.ServerURL + route
	}
	redirect(c, path)
}

// // 路由重定向 use router name
// func RedirectRouter(c *gin.Context, routerName string, args ...interface{}) {
// 	redirect(c, named.G(routerName, args...))
// }

// 从 path params 中获取 int 参数：http://a.com/xx/1 => 获取到 int 1
func GetIntParam(c *gin.Context, key string) (int, error) {
	i, err := strconv.Atoi(c.Param(key))
	if err != nil {
		return 0, err
	}
	return i, nil
}

// GetPageQuery 从 query 中获取有关分页的参数：// xx.com?page=1&pageline=10
func GetPageQuery(c *gin.Context, defaultPageLine, totalCount int) (offset, limit, currentPage, pageTotalCount int) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	currentPage = page
	pageline, err := strconv.Atoi(c.Query("pageline"))
	if err != nil {
		pageline = defaultPageLine
	}
	page = page - 1
	if page == 0 {
		offset = 0
	} else {
		offset = page * pageline
	}
	limit = pageline
	pageTotalCount = int(math.Ceil(float64(totalCount) / float64(pageline)))
	if pageTotalCount <= 0 {
		pageTotalCount = 1
	}
	return
}

// func backTo(c *gin.Context, currentUser *userModel.User) {
// 	back := c.DefaultPostForm("back", "")
// 	if back != "" {
// 		controllers.Redirect(c, back, true)
// 		return
// 	}
// 	controllers.RedirectRouter(c, "users.show", currentUser.ID)
// }
