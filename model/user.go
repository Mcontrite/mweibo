package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Email       string `gorm:""`
	Avatar      string `gorm:"not null"`
	ActiveToken string `gorm:""`
	IsActive    bool   `gorm:"default:0"`
	IsAdmin     bool   `gorm:"default:1"`
	// EmailVertifyAt  time.Time `gorm:""`
	RememberMeToken string    `gorm:""`
	SecretKey       string    `gorm:"default:null"`
	ExpireTime      time.Time `gorm:"default:null"`
}

func CreateUser(user *User) error {
	return DB.Save(user).Error
	// return DB.Create(user).Error
}

// func (user *User) CreateUser() error {
// 	return DB.Save(user).Error
// }
func UpdateUser(maps interface{}, items map[string]interface{}) (err error) {
	err = DB.Model(&User{}).Where(maps).Updates(items).Error
	return
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

func (user *User) DeleteUser(id int) error {
	user.ID = uint(id)
	// Unscoped 永久删除
	return DB.Unscoped().Delete(&user).Error
}

func DelUser(maps interface{}) (err error) {
	err = DB.Unscoped().Where(maps).Delete(&User{}).Error
	return
}

// func GetUserByID(id int) (user *User, err error) {
// 	err = DB.First(&user, id).Error
// 	return
// }

// func GetUserByID(id interface{}) (user *User, err error) {
// 	err = DB.First(&user, id).Error
// 	return
// }
func IfUserExist(username string) bool {
	var user User
	DB.Model(&User{}).Select("id").Where("username=?", username).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func GetUser(maps interface{}) (user User, err error) {
	err = DB.Model(&User{}).Where(maps).First(&user).Error
	return
}

func GetUserByID(id int) (user User, err error) {
	err = DB.Model(&User{}).Where("id=?", id).First(&user).Error
	return
}

// func GetUserByUsername(name string) (user *User, err error) {
// 	err = DB.First(&user, "username=?", name).Error
// 	return
// }

func GetUserByUsername(name string) (*User, error) {
	var user User
	err := DB.First(&user, "username=?", name).Error
	return &user, err
}

func GetUserByEmail(email string) (user *User, err error) {
	err = DB.First(&user, "email=?", email).Error
	return
}

// func GetUserByWeiboID(weiboid int) (user *User, err error) {
// 	weibo, _ := GetWeiboByID(weiboid)
// 	user, err = GetUserByID(int(weibo.UserID))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, err
// }

func CountUsers() (count int) {
	DB.Model(&User{}).Count(&count)
	return
}

func ListUsers() (users []*User, err error) {
	err = DB.Find(&users, "is_admin=?", true).Error
	return
}

func GetUsers(limit int, order string, maps interface{}) (user []User, err error) {
	err = DB.Model(&User{}).Order(order).Limit(limit).Find(&user).Error
	return
}

func (user *User) GetUserAvatar() string {
	if user.Avatar != "" {
		return user.Avatar
	}
	return ""
}
