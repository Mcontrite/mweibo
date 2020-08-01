package model

import "github.com/jinzhu/gorm"

type TagWeibo struct {
	gorm.Model
	WeiboID uint
	TagID   uint
}
