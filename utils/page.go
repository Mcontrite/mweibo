package utils

// import (
// 	"math"
// 	"myzone/package/setting"
// 	"strconv"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/unknwon/com"
// )

// func GetPage(c *gin.Context) int {
// 	result := 0
// 	page, _ := com.StrTo(c.Query("page")).Int()
// 	if page > 0 {
// 		result = (page - 1) * setting.ServerSetting.PageSize
// 	}
// 	return result
// }

// func Pagination_tpl(url, text, active string) string {
// 	g_pagination_tpl := "<li class='page-item{active}'><a href='{url}' class='page-link'>{text}</a></li>"
// 	g_pagination_tpl = strings.Replace(g_pagination_tpl, "{url}", url, 1)
// 	g_pagination_tpl = strings.Replace(g_pagination_tpl, "{text}", text, 1)
// 	g_pagination_tpl = strings.Replace(g_pagination_tpl, "{active}", active, 1)
// 	return g_pagination_tpl
// }

// //bootstrap 翻页，命名与 bootstrap 保持一致
// func Pagination(url string, totalnum int, page int, pagesize int) (s string) {
// 	if pagesize == 0 {
// 		pagesize = 20
// 	}
// 	totalpage := math.Ceil(float64(totalnum) / float64(pagesize))
// 	if totalpage < 2 {
// 		return
// 	}
// 	page = int(math.Min(float64(totalpage), float64(page)))
// 	shownum := 5
// 	start := int(math.Max(1, float64(page-shownum)))
// 	end := int(math.Max(totalpage, float64(page+shownum)))
// 	// 不足 $shownum，补全左右两侧
// 	right := page + shownum - int(totalpage)
// 	if right > 0 {
// 		start -= right
// 		start = int(math.Max(1, float64(start)))
// 	}
// 	left := page - shownum
// 	if left < 0 {
// 		end -= left
// 		end = int(math.Min(totalpage, float64(end)))
// 	}
// 	if page != 1 {
// 		url := strings.Replace(url, "{page}", strconv.Itoa(page-1), 1)
// 		s += Pagination_tpl(url, "◀", "")
// 	}
// 	if start > 1 {
// 		text := "1 "
// 		if start > 2 {
// 			text += "..."
// 		}
// 		url := strings.Replace(url, "{page}", "1", 1)
// 		s += Pagination_tpl(url, text, "")
// 	}
// 	for i := start; i <= end; i++ {
// 		text := ""
// 		if i == page {
// 			text += " active"
// 		}
// 		url := strings.Replace(url, "{page}", strconv.Itoa(i), 1)
// 		s += Pagination_tpl(url, strconv.Itoa(i), text)
// 	}
// 	if end != int(totalpage) {
// 		text := ""
// 		if (int(totalpage) - end) > 1 {
// 			text = "..." + strconv.Itoa(int(totalpage))
// 		} else {
// 			text = strconv.Itoa(int(totalpage))
// 		}
// 		url := strings.Replace(url, "{page}", strconv.Itoa(int(totalpage)), 1)
// 		s += Pagination_tpl(url, text, "")
// 	}
// 	if page != int(totalpage) {
// 		url := strings.Replace(url, "{page}", strconv.Itoa(page+1), 1)
// 		s += Pagination_tpl(url, "▶", "")
// 	}
// 	return
// }
///////////////////////////////////////////////////////////////////////////////////////////
// import (
// 	"net/url"

// 	"github.com/gin-gonic/gin"
// )

// type paginationRenderData struct {
// 	URL                string // 分页的 root url
// 	CurrentPage        int    // 当前页码
// 	OnFirstPage        bool   // 是否在第一页
// 	HasMorePages       bool   // 是否有更多页
// 	Elements           []int  // 页码
// 	PreviousButtonText string // 前一页按钮文本
// 	PreviousPageIndex  int    // 前一页按钮的页码
// 	NextButtonText     string // 后一页按钮文本
// 	NextPageIndex      int    // 后一页按钮的页码
// }

// // baseOnCurrentPageButtonOffset: 前后有多少个按钮；返回一个区间数组，供生成区间页码按钮
// func countStartAndEndPageIndex(currentPage, totalPage, baseOnCurrentPageButtonOffset int) []int {
// 	howMuchPageButtons := baseOnCurrentPageButtonOffset*2 + 1
// 	startPage := 1
// 	endPage := 1
// 	result := make([]int, 0)
// 	if currentPage > baseOnCurrentPageButtonOffset {
// 		// 当前页码大于偏移量，则起始按钮为：当前页码 - 偏移量
// 		startPage = currentPage - baseOnCurrentPageButtonOffset
// 		if totalPage > (currentPage + baseOnCurrentPageButtonOffset) {
// 			endPage = currentPage + baseOnCurrentPageButtonOffset
// 		} else {
// 			endPage = totalPage
// 		}
// 	} else {
// 		// 当前页码小于偏移量
// 		startPage = 1
// 		if totalPage > howMuchPageButtons {
// 			endPage = howMuchPageButtons
// 		} else {
// 			endPage = totalPage
// 		}
// 	}
// 	if (currentPage + baseOnCurrentPageButtonOffset) > totalPage {
// 		startPage = startPage - (currentPage + baseOnCurrentPageButtonOffset - endPage)
// 	}
// 	if startPage <= 0 {
// 		startPage = 1
// 	}
// 	for i := startPage; i <= endPage; i++ {
// 		result = append(result, i)
// 	}
// 	return result
// }

// // 生成分页模板所需的数据
// func CreatePaginationFillToTplData(c *gin.Context, pageQueryKeyName string, currentPage, totalPage int, otherData map[string]interface{}) map[string]interface{} {
// 	queryValues := url.Values{}
// 	for k, v := range c.Request.URL.Query() {
// 		if k != pageQueryKeyName {
// 			queryValues.Add(k, v[0])
// 		}
// 	}
// 	query := queryValues.Encode()
// 	if query != "" {
// 		query = query + "&"
// 	}
// 	pageData := paginationRenderData{
// 		URL:                c.Request.URL.Path + "?" + query + pageQueryKeyName + "=",
// 		CurrentPage:        currentPage,
// 		OnFirstPage:        currentPage == 1,
// 		HasMorePages:       currentPage != totalPage,
// 		Elements:           countStartAndEndPageIndex(currentPage, totalPage, 3),
// 		PreviousButtonText: "前一页",
// 		PreviousPageIndex:  currentPage - 1,
// 		NextButtonText:     "下一页",
// 		NextPageIndex:      currentPage + 1,
// 	}
// 	otherData["pagination"] = pageData
// 	return otherData
// }
