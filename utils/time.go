package utils

import (
	"math"
	"strconv"
	"time"
)

func TimeFormat(template string) (formatTime string) {
	switch template {
	case "Ymd":
		formatTime = time.Now().Format("20060102")
	case "Y/m/d":
		formatTime = time.Now().Format("2006/01/02")
	case "Y-m-d":
		formatTime = time.Now().Format("2006-01-02")
	case "H:i:s":
		formatTime = time.Now().Format("15:04:05")
	case "Ymd H:i:s":
		formatTime = time.Now().Format("20060102 15:04:05")
	case "Y-m-d  H:i:s":
		formatTime = time.Now().Format("2006-01-02  15:04:05")
	default:
		formatTime = time.Now().Format("2006/01/02  15:04:05")
	}
	return
}

// 类似于 1小时前 这样的展示方式
func SinceForHuman(t time.Time) string {
	duration := time.Since(t)
	hour := duration.Hours()
	minutes := duration.Minutes()
	seconds := duration.Seconds()
	unit := "秒"
	s := 0
	if hour > (365 * 24) {
		s = int(math.Floor(hour / (365 * 24)))
		unit = "年"
	} else if hour > (30 * 24) {
		s = int(math.Floor(hour / (30 * 24)))
		unit = "月"
	} else if hour > 24 {
		s = int(math.Floor(hour / 24))
		unit = "天"
	} else if hour > 1 {
		s = int(math.Floor(hour))
		unit = "小时"
	} else if minutes > 1 {
		s = int(math.Floor(minutes))
		unit = "分钟"
	} else if seconds > 0 {
		return "刚刚"
	}
	return strconv.Itoa(s) + unit + "前"
}
