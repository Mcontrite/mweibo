package model

import (
	"database/sql"
	"html/template"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
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

func CountWeibos() (count int) {
	DB.Model(&Weibo{}).Count(&count)
	return
}

func CountWeibosByTag(tag string) (count int, err error) {
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

func GetWeiboByID(id interface{}) (*Weibo, error) {
	weibo := &Weibo{}
	err := DB.First(&weibo, id).Error
	return weibo, err
}

func GetWeiboByTagID(tag string) (count int, err error) {
	if len(tag) > 0 {
		tagid, _ := strconv.ParseUint(tag, 10, 64)
		err = DB.Raw(
			"select count(*) from weibos w inner join tag_weibos tw on w.id=tw.weibo_id "+
				"where tw.tag_id=?", tagid,
		).Row().Scan(&count)
	} else {
		err = DB.Raw("select count(*) from weibos w").Row().Scan(&count)
	}
	return
}

func ListWeibos(tag string) (weibos []*Weibo, err error) {
	if tag != "" {
		tagid, _ := strconv.ParseUint(tag, 10, 64)
		var rows *sql.Rows
		rows, err = DB.Raw(
			"select w.* from weibos w left join tag_weibos tw on w.id=tw.weibo_id "+
				"where tw.tag_id=? order by created_at desc",
			tagid,
		).Rows()
		defer rows.Close()
		for rows.Next() {
			var weibo Weibo
			DB.ScanRows(rows, &weibo)
			weibos = append(weibos, &weibo)
		}
	} else {
		err = DB.Order("created_at desc").Find(&weibos).Error
	}
	return
}

func ListMaxViewWeibos() (weibos []*Weibo, err error) {
	err = DB.Order("views_cnt desc").Limit(5).Find(&weibos).Error
	return
}

func ListMaxCommentWeibos() (weibos []*Weibo, err error) {
	var rows *sql.Rows
	rows, err = DB.Raw(
		"select w.*,c.num comments_cnt from weibos w inner join (select weibo_id,count(1) " +
			"num from comments group by weibo_id) c on w.id=c.weibo_id order by c.num " +
			"desc limit 5 ",
	).Rows()
	defer rows.Close()
	for rows.Next() {
		var weibo Weibo
		DB.ScanRows(rows, &weibo)
		weibos = append(weibos, &weibo)
	}
	return
}

func Excerpt(weibo *Weibo) template.HTML {
	policy := bluemonday.StrictPolicy()
	sanitize := policy.Sanitize(string(blackfriday.Run([]byte(weibo.Content))))
	runes := []rune(sanitize)
	if len(runes) > 300 {
		sanitize = string(runes[:300])
	}
	excerpt := template.HTML(sanitize + "...")
	return excerpt
}
