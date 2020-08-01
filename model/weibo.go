package model

import (
	"github.com/jinzhu/gorm"
)

type Weibo struct {
	gorm.Model
	UserID      uint
	Content     string
	ViewsCnt    int
	CommentsCnt int
	Tags        []*Tag
	Comments    []*Comment
}
