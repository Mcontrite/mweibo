package model

import "github.com/jinzhu/gorm"

type Weibo struct {
	gorm.Model
	UserID      uint
	Content     string
	ViewsCnt    int
	CommentsCnt int
	User        User
	Comments    []*Comment
}

func CreateWeibo(weibo *Weibo) error {
	return DB.Create(weibo).Error
}

func GetWeiboObjectByID(id int) (weibo Weibo, err error) {
	err = DB.Model(&Weibo{}).Where("id=?", id).First(&weibo).Error
	return
}

func UpdateWeibo(weibo *Weibo, weiboid int) error {
	return DB.Model(&weibo).Where("id=?", weiboid).Updates(map[string]interface{}{
		"content": weibo.Content,
	}).Error
}

func UpdateWeiboViewsCnt(weibo *Weibo) error {
	return DB.Model(weibo).Updates(map[string]interface{}{
		"view_cnt": weibo.ViewsCnt,
	}).Error
}

func DeleteWeibo(weibo *Weibo) error {
	return DB.Delete(weibo).Error
}
