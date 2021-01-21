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

func UserCreate(user *User) error {
	return DB.Save(user).Error
}

func GetUserObjectByID(id int) (user User, err error) {
	err = DB.Model(&User{}).Where("id=?", id).First(&user).Error
	return
}

func GetUserObjectByMap(m interface{}) (user User, err error) {
	err = DB.Model(&User{}).Where(m).First(&user).Error
	return
}
