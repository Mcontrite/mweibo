package model

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
)

type Tag struct {
	gorm.Model
	Content string
	Num     int
}

func CreateTag(tag *Tag) error {
	return DB.FirstOrCreate(tag, "content=?", tag.Content).Error
}

func CountTags() (count int) {
	DB.Model(&Tag{}).Count(&count)
	return count
}

func ListTags() (tags []*Tag, err error) {
	rows, err := DB.Raw("select t.*, count(*) total from tags t inner join tag_weibos tw on " +
		"t.id=tw.tag_id inner join weibos w on tw.weibo_id=w.id group by tw.tag_id",
	).Rows()
	defer rows.Close()
	for rows.Next() {
		var tag Tag
		DB.ScanRows(rows, &tag)
	}
	return tags, nil
}

func ListTagsByWeiboID(weiboid string) (tags []*Tag, err error) {
	weiid, err := strconv.ParseUint(weiboid, 10, 64)
	if err != nil {
		log.Warnf("Parse error...")
	}
	rows, err := DB.Raw("select t.* from tags t inner join tag_weibos tw on t.id=tw.tag_id where tw.weiboid=?",
		uint(weiid),
	).Rows()
	defer rows.Close()
	for rows.Next() {
		var tag Tag
		DB.ScanRows(rows, &tag)
		tags = append(tags, &tag)
	}
	return
}
