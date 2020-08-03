package model

import "github.com/jinzhu/gorm"

type TagWeibo struct {
	gorm.Model
	WeiboID uint
	TagID   uint
}

// func ListTagWeiboByWeiboID(weiboid string)(tags []*Tag,err error){
// 	weiid,_:=strconv.ParseUint(weiboid,10,64)
// 	rows,err:=DB.Raw(
// 		"select t.* from tags t inner join tag_weibos tw on t.id=tw.tag_id where tw.weibo_id=?",
// 		uint(weiid),
// 	).Rows()
// 	defer rows.Close()
// 	for rows.Next(){
// 		var tagweibo TagWeibo

// 	}
// }

func CreateTagWeibo(tagweibo *TagWeibo) error {
	return DB.Create(tagweibo).Error
}
