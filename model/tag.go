package model

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Content string
	Num     int
}
