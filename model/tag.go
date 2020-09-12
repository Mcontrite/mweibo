package model

import (
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	TagName   string
	WeibosCnt int // count
}

func CreateTag(tag *Tag) error {
	return DB.FirstOrCreate(tag, "tag_name=?", tag.TagName).Error
}

func ListTags() (tags []*Tag, err error) {
	rows, err := DB.Raw("select t.*, count(*) weibos_cnt from tags t inner join tag_weibos " +
		"tw on t.id=tw.tag_id inner join weibos w on tw.weibo_id=w.id group by tw.tag_id",
	).Rows()
	defer rows.Close()
	for rows.Next() {
		var tag Tag
		DB.ScanRows(rows, &tag)
		tags = append(tags, &tag)
	}
	return tags, nil
}

func ListTagNames() (str string) {
	tags, _ := ListTags()
	sli := make([]string, 0)
	for _, v := range tags {
		sli = append(sli, v.TagName)
	}
	str = strings.Join(sli, ",")
	return
}

func ListTagsByWeiboID(id string) (tags []*Tag, err error) {
	weiboid, _ := strconv.ParseUint(id, 10, 64)
	rows, err := DB.Raw("select t.* from tags t inner join tag_weibos tw on"+
		" t.id=tw.tag_id where tw.weibo_id=?", uint(weiboid),
	).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag Tag
		DB.ScanRows(rows, &tag)
		tags = append(tags, &tag)
	}
	return
}

func CountTags() (count int) {
	DB.Model(&Tag{}).Count(&count)
	return
}
