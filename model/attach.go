package model

import "github.com/jinzhu/gorm"

type Attach struct {
	gorm.Model
	UserID       int    `gorm:"default:0" json:"user_id"`
	WeiboID      int    `gorm:"default:0" json:"weibo_id"`
	CommentID    int    `gorm:"default:0" json:"comment_id"`
	Filename     string `gorm:"default:''" json:"filename"`  //文件名称，会过滤，并且截断，保存后的文件名不包含URL前缀 upload_url
	OrigiName    string `gorm:"default:''" json:"originame"` //上传的原文件名
	Filetype     string `gorm:"default:''" json:"filetype"`  //image/txt/zip，小图标显示
	Filesize     int    `gorm:"default:0" json:"filesize"`
	Isimage      int    `gorm:"default:0" json:"isimage"`       //是否为图片
	Width        int    `gorm:"default:0" json:"width"`         //width > 0 则为图片
	Height       int    `gorm:"default:0" json:"height"`        //
	DownloadsCnt int    `gorm:"default:0" json:"downloads_cnt"` //下载次数
}

func CreateAttach(attach *Attach) (*Attach, error) {
	err := DB.Model(&Attach{}).Create(attach).Error
	return attach, err
}

func GetAttachsByWeiboID(id int) (attachs []Attach, err error) {
	err = DB.Model(&Attach{}).Where("weibo_id = ?", id).Find(&attachs).Error
	return
}

func GetAttachsByCommentID(id int) (attachs []Attach, err error) {
	err = DB.Model(&Attach{}).Where("comment_id = ?", id).Find(&attachs).Error
	return
}

func DeleteAttachByID(id int) error {
	return DB.Model(&Attach{}).Where("id = ?", id).Unscoped().Delete(&Attach{}).Error
}

func DeleteWeibosAttachsByIDs(ids []string) (err error) {
	err = DB.Unscoped().Where("weibo_id in (?)", ids).Delete(&Attach{}).Error
	return
}

func CountAttachs() (count int, err error) {
	err = DB.Model(&Attach{}).Count(&count).Error
	return
}
