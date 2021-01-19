package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username        string
	Password        string
	Email           string
	Avatar          string
	IsActive        bool
	IsAdmin         bool
	ActiveToken     string
	RememberMeToken string
}
