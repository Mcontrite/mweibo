package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username        string `gorm:"column:username;type:varchar(255);not null"`
	Password        string `""`
	Email           string
	Avatar          string
	ActiveToken     string
	IsActive        uint
	IsAdmin         uint
	EmailVertifyAt  time.Time
	RememberMeToken string
	SecretKey       string `gorm:"default:null"`
}
