package utils

import (
	"strings"
	"time"
)

func GetCurrentTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc)
}

// 格式化时间
func DateFormat(date time.Time, layout string) string {
	return date.Format(layout)
}

// 截取字符串
func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) > n {
		return string(runes[:n])
	}
	return s
}

// 截取字符串
func Substring(source string, start, end int) string {
	rs := []rune(source)
	length := len(rs)
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	return string(rs[start:end])
}

// 判断数字是否是偶数
func IsEven(number int) bool {
	return number%2 == 0
}

// 判断数字是否是奇数
func IsOdd(number int) bool {
	return !IsEven(number)
}

// 求和
func Add(a1, a2 int) int {
	return a1 + a2
}

// 相减
func Minus(a1, a2 int) int {
	return a1 - a2
}

// 简单的解析模板方法
func ParseEasyTemplate(tplString string, data map[string]string) string {
	replaceArr := []string{}
	for k, v := range data {
		replaceArr = append(replaceArr, k, v)
	}
	r := strings.NewReplacer(replaceArr...)
	return r.Replace(tplString)
}
