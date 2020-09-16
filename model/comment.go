package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserID     uint   `gorm:"default:0" json:"user_id"`
	WeiboID    uint   `gorm:"default:0" json:"weibo_id"`
	Content    string `gorm:"default:''" json:"content"`
	IsRead     bool   `gorm:"default:0" json:"isread"`
	AttachsCnt int    `gorm:"default:0" json:"attachs_cnt"`
	// ImageURL
	User   User
	Weibo  Weibo
	Attach []Attach
}

func CreateComment(comment *Comment) error {
	return DB.Create(&comment).Error
}

func NewComment(comment *Comment) (*Comment, error) {
	err := DB.Create(comment).Error
	return comment, err
}

func UpdateCommentByID(id int, comment Comment) (newcommentt Comment, err error) {
	err = DB.Model(&Comment{}).Where("id = ?", id).Updates(comment).Error
	newcommentt, err = GetCommentByID(id)
	return
}

func UpdateCommentAttachsCnt(id int, num int) error {
	return DB.Model(&Comment{}).Where("id=?", id).Update("attachs_cnt", num).Error
}

func DeleteComment(comment *Comment) error {
	return DB.Delete(comment, "user_id=? and id=?", comment.UserID, comment.ID).Error
}

func DeleteCommentsByWeiboIDs(ids []string) (err error) {
	err = DB.Unscoped().Where("weibo_id in (?)", ids).Delete(&Comment{}).Error
	return
}

func GetCommentByID(id int) (comment Comment, err error) {
	err = DB.Model(&Comment{}).Where("id=?", id).First(&comment).Error
	return
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
	err = DB.Preload("User").Preload("Attach").Model(&Comment{}).Where("weibo_id=?", id).Limit(limit).Find(&comment).Error
	return
}

func ListUnReadComments() (comments []*Comment, err error) {
	err = DB.Where("is_read=?", false).Order("created_at desc").Find(&comments).Error
	return comments, err
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

func CountComments() (count int) {
	DB.Model(&Comment{}).Count(&count)
	return
}

// func AfterAddNewComment(comment *model.Comment, articleID int) {
// 	articleInfo, _ := model.GetArticleById(articleID)
// 	updateArticle := model.Article{
// 		LastDate:    time.Now(),
// 		CommentsCnt:   articleInfo.CommentsCnt + 1,
// 		LastUserID:  comment.UserID,
// 		LastCommentID: int(comment.ID),
// 	}
// 	model.UpdateArticle(articleID, updateArticle)
// 	model.AddMyComment(comment.UserID, articleID, int(comment.ID))
// 	oldUserInfo, _ := model.GetUserByID(comment.UserID)
// 	model.UpdateUserCommentsCnt(comment.UserID, oldUserInfo.CommentsCnt+1)
// }
