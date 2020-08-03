package model

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
)

type Comment struct {
	gorm.Model
	UserID  uint
	WeiboID uint
	Content string
	IsRead  bool `gorm:"default:0"`
}

func CreateComment(comment *Comment) error {
	return DB.Create(&comment).Error
}

func DeleteComment(comment *Comment) error {
	return DB.Delete(comment, "user_id=? and id=?", comment.UserID, comment.ID).Error
}

func CountComments() (count int) {
	DB.Model(&Comment{}).Count(&count)
	return
}

func UpdateCommentIsRead(comment *Comment) error {
	return DB.Model(&comment).Updates(map[string]interface{}{
		"is_read": true,
	}).Error
}

func SetAllCommentRead() error {
	return DB.Model(&Comment{}).Where("is_read=?", false).Updates(map[string]interface{}{
		"is_read": true,
	}).Error
}

func ListUnReadComments() (comments []*Comment, err error) {
	err = DB.Where("is_read=?", false).Order("created_at desc").Find(&comments).Error
	return comments, err
}

func ListCommentsByWeiboID(weiboid string) (comments []*Comment, err error) {
	weiid, err := strconv.ParseUint(weiboid, 10, 64)
	if err != nil {
		log.Warnf("Parse int error...")
	}
	rows, err := DB.Raw(
		"select c.*,u.avatar from comments c inner join users u on c.user_id=u.id"+
			" where c.post_id=? order by created_at desc ",
		uint(weiid),
	).Rows()
	if err != nil {
		log.Warnf("DB Row error...")
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comment
		DB.ScanRows(rows, &comment)
		comments = append(comments, &comment)
	}
	return comments, err
}
