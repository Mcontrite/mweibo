package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username        string    `gorm:"not null"`
	Password        string    `gorm:"not null"`
	Email           string    `gorm:"unique;not null"`
	Avatar          string    `gorm:"not null"`
	ActiveToken     string    `gorm:""`
	IsActive        uint      `gorm:"type:tinyint(1);default:0"`
	IsAdmin         uint      `gorm:"type:tinyint(1);default:1"`
	EmailVertifyAt  time.Time `gorm:""`
	RememberMeToken string    `gorm:""`
	SecretKey       string    `gorm:"default:null"`
	ExpireTime      time.Time `gorm:"default:null"`
}
