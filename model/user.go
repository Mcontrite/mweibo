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
	IsActive        bool      `gorm:"default:0"`
	IsAdmin         bool      `gorm:"default:1"`
	EmailVertifyAt  time.Time `gorm:""`
	RememberMeToken string    `gorm:""`
	SecretKey       string    `gorm:"default:null"`
	ExpireTime      time.Time `gorm:"default:null"`
}

func CreateUser(user *User) error {
	return DB.Save(user).Error
}

func UpdateUserAvatar(user *User, avatar string) error {
	return DB.Model(&user).Update(User{Avatar: avatar}).Error
}

func UpdateUserEmail(user *User, email string) error {
	if len(email) > 0 {
		return DB.Model(user).Update("email", email).Error
	}
	return DB.Model(user).Update("email", gorm.Expr("NULL")).Error
}

func GetUserByID(id interface{}) (*User, error) {
	user := &User{}
	err := DB.First(&user, id).Error
	return user, err
}

func GetUserByUsername(name string) (*User, error) {
	user := &User{}
	err := DB.First(&user, "username=?", name).Error
	return user, err
}

func GetUserByEmail(email string) (*User, error) {
	user := &User{}
	err := DB.First(&user, "email=?", email).Error
	return user, err
}

func ListUsers() ([]*User, error) {
	users := make([]*User, 0)
	err := DB.Find(&users, "is_admin=?", true).Error
	return users, err
}
