package conf

import (
	"github.com/lexkong/log"
)

func InitLog() {
	logConfig := log.PassLagerCfg{
		Writers:        "file,stdout",    // 输出位置
		LoggerLevel:    "DEBUG",          // 日志级别
		LoggerFile:     "log/mweibo.log", // 文件位置
		LogFormatText:  false,            // true 输出成 json 格式
		RollingPolicy:  "daily",          // 根据日期转存
		LogRotateDate:  1,                // 配合 daily 使用
		LogRotateSize:  1,                // 配合 size 使用
		LogBackupCount: 7,                // 指定转存文件的最大个数
	}
	log.InitWithConfig(&logConfig)
}
