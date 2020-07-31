package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github,com/lexkong/log"
)

var db *gorm.DB

func createDBUrl(uname ,pwd , host ,port ,dbname string) string{
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		uname,pwd,host,port,dbname,true,"Local",
	)
}

func ConnectDB(){
	db,err:=gorm.Open
}