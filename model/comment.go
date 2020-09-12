package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"default:0" json:"user_id"`
	WeiboID uint   `gorm:"default:0" json:"weibo_id"`
	Content string `gorm:"default:''" json:"content"`
	IsRead  bool   `gorm:"default:0" json:"isread"`
	// ImageURL
	User   User
	Weibo  Weibo
	Attach []Attach
}

func CreateComment(comment *Comment) error {
	return DB.Create(&comment).Error
}

func DeleteComment(comment *Comment) error {
	return DB.Delete(comment, "user_id=? and id=?", comment.UserID, comment.ID).Error
}

func SetCommentRead(comment *Comment) error {
	return DB.Model(&comment).Updates(map[string]interface{}{
		"is_read": true,
	}).Error
}

func SetAllCommentsRead() error {
	return DB.Model(&Comment{}).Where("is_read=?", false).Updates(map[string]interface{}{
		"is_read": true,
	}).Error
}

// func ListCommentsByWeiboID(weiboid string) (comments []*Comment, err error) {
// 	weiid, err := strconv.ParseUint(weiboid, 10, 64)
// 	if err != nil {
// 		log.Warnf("Parse int error...")
// 	}
// 	rows, err := DB.Raw(
// 		"select c.*,u.avatar from comments c inner join users u on c.user_id=u.id"+
// 			" where c.post_id=? order by created_at desc ",
// 		uint(weiid),
// 	).Rows()
// 	if err != nil {
// 		log.Warnf("DB Row error...")
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var comment Comment
// 		DB.ScanRows(rows, &comment)
// 		comments = append(comments, &comment)
// 	}
// 	return
// }
func ListCommentsByWeiboID(id int, limit int) (comment []Comment, err error) {
	err = DB.Preload("User").Model(&Comment{}).Where("weibo_id=?", id).Limit(limit).Find(&comment).Error
	return
}

func ListUnReadComments() (comments []*Comment, err error) {
	err = DB.Where("is_read=?", false).Order("created_at desc").Find(&comments).Error
	return comments, err
}

func GetCommentByID(id int) (comment Comment, err error) {
	err = DB.Model(&Comment{}).Where("id = ?", id).First(&comment).Error
	return
}

func UpdateComment(id int, comment Comment) (newcommentt Comment, err error) {
	err = DB.Model(&Comment{}).Where("id = ?", id).Updates(comment).Error
	newcommentt, err = GetCommentByID(id)
	return
}

func CountComments() (count int) {
	DB.Model(&Comment{}).Count(&count)
	return
}
