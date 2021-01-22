package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	UserID  uint
	WeiboID uint
	Content string
	IsRead  bool
	User    User
	Weibo   Weibo
}

func CreateComment(comment *Comment) error {
	return DB.Create(comment).Error
}

func GetCommentByID(id int) (comment Comment, err error) {
	err = DB.Model(&Comment{}).Where("id=?", id).First(&comment).Error
	return
}

func UpdateCommentByID(comment Comment, id int) (newcom Comment, err error) {
	err = DB.Model(&Comment{}).Where("id=?", id).Updates(comment).Error
	newcom, err = GetCommentByID(id)
	return
}

func ListCommentsByWeiboID(id, limit int) (comments []Comment, err error) {
	err = DB.Preload("User").Model(&Comment{}).Where("weibo_id=?", id).Limit(limit).Find(&comments).Error
	return
}
