package model

import (
	"time"
)

type PwdReset struct {
	Email     string `gorm:"not null" sql:"index"`
	Token     string `gorm:"not null" sql:"index"`
	CreatedAt time.Time
}
