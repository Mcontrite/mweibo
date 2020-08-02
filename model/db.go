package model

import (
	"fmt"
	"time"

	"github.com/lexkong/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// type BaseModel struct{
// 	gorm.Model
// }

var DB *gorm.DB
var err error

func createDBUrl(user, pwd, host, port, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		user, pwd, host, port, dbname,
	)
}

func InitDB() *gorm.DB {
	// dbtype := conf.GetConfiguration().DBType
	// user := conf.GetConfiguration().DBUser
	// pwd := conf.GetConfiguration().DBPassword
	// host := conf.GetConfiguration().DBHost
	// port := conf.GetConfiguration().DBPort
	// dbname := conf.GetConfiguration().DBName
	// url := createDBUrl(user, pwd, host, port, dbname)
	// db, err := gorm.Open(dbtype, url)
	// db, err := gorm.Open(conf.GetConfiguration().DBType, conf.GetConfiguration().DSN)
	str := "root:123456@tcp(127.0.0.1:3306)/mweibo?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", str)
	if err != nil {
		log.Fatal("Connect database  failed: ", err)
		panic("Connect database  failed...")
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	db.LogMode(true)
	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8;").AutoMigrate()
	DB = db
	return db
}
