package model

import (
	"fmt"
	"mweibo2/conf"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dbtype := conf.DBConf.DBType
	dbuser := conf.DBConf.DBUser
	dbpass := conf.DBConf.DBPass
	dbhost := conf.DBConf.DBHost
	dbport := conf.DBConf.DBPort
	dbname := conf.DBConf.DBName

	dburl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbuser, dbpass, dbhost, dbport, dbname,
	)
	db, err := gorm.Open(dbtype, dburl)
	if err != nil {
		fmt.Println("Connect DB err: ", err)
		return nil
	}
	db.DB().SetMaxOpenConns(20)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	db.LogMode(true)
	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8;").AutoMigrate()
	DB = db
	return db
}
