package model

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
)

type Weibo struct {
	gorm.Model
	UserID      uint   `gorm:"not null" sql:"index"`
	Content     string `gorm:"type:text;not null"`
	ViewsCnt    int    `gorm:"default:0"`
	CommentsCnt int    `gorm:"default:0"`
	Tags        []*Tag
	Comments    []*Comment
}

func CreateWeibo(weibo *Weibo) error {
	return DB.Create(weibo).Error
}

func UpdateWeibo(weibo *Weibo, id string) error {
	return DB.Model(weibo).Where("id=?", id).Updates(map[string]interface{}{
		"content": weibo.Content,
	}).Error
}

func UpdateWeiboViewsCnt(weibo *Weibo) error {
	return DB.Model(&weibo).Updates(map[string]interface{}{
		"views_cnt": weibo.ViewsCnt,
	}).Error
}

func DeleteWeibo(weibo *Weibo) error {
	return DB.Delete(weibo).Error
}

func GetWeiboByID(id interface{}) (*Weibo, error) {
	weibo := &Weibo{}
	err := DB.First(&weibo, id).Error
	return weibo, err
}

func CountWeiboByTag(tag string) (count int, err error) {
	if len(tag) > 0 {
		tagid, err := strconv.ParseUint(tag, 10, 64)
		if err != nil {
			log.Warnf("Parse tagid error...")
		}
		err = DB.Raw(
			"select count(*) from weibos w inner join tag_weibos tw on w.id=tw.weibo_id where tw.tag_id=?",
			tagid,
		).Row().Scan(&count)
	} else {
		err = DB.Raw("select count(*) from weibos").Row().Scan(&count)
	}
	return
}
