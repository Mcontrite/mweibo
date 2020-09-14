package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username        string `gorm:"not null" json:"username"`
	Password        string `gorm:"not null" json:"password"`
	Email           string `gorm:"" json:"email"`
	Avatar          string `gorm:"not null" json:"avatar"`
	IsActive        bool   `gorm:"default:0" json:"isactive"`
	IsAdmin         bool   `gorm:"default:1" json:"isadmin"`
	ActiveToken     string `gorm:"" json:"activetoken"`
	RememberMeToken string `gorm:""`
	// EmailVertifyAt  time.Time `gorm:""`
}

// func (user *User) CreateUser() error {
// 	return DB.Save(user).Error
// }
func CreateUser(user *User) error {
	return DB.Save(user).Error
	// return DB.Create(user).Error
}

func (user *User) UpdateUser() (err error) {
	return DB.Save(&user).Error
}

func UpdateUserByMap(maps interface{}, items map[string]interface{}) (err error) {
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

func IfUsernameExist(username string) bool {
	var user User
	DB.Model(&User{}).Select("id").Where("username=?", username).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func GetUserByMaps(maps interface{}) (user User, err error) {
	err = DB.Model(&User{}).Where(maps).First(&user).Error
	return
}

// func GetUserByID(id interface{}) (user *User, err error) {
// 	err = DB.First(&user, id).Error
// 	return
// }
func GetUserByID(id int) (user *User, err error) {
	err = DB.First(&user, id).Error
	return
}

func GetUserObjectByID(id int) (user User, err error) {
	err = DB.Model(&User{}).Where("id=?", id).First(&user).Error
	return
}

// func GetUserByUsername(name string) (*User, error) {
// 	var user User
// 	err := DB.First(&user, "username=?", name).Error
// 	return &user, err
// }
func GetUserByUsername(name string) (user *User, err error) {
	err = DB.First(&user, "username=?", name).Error
	return
}

func GetUserByEmail(email string) (user *User, err error) {
	// err=DB.Where("email=?",email).First(&user).Error
	err = DB.First(&user, "email=?", email).Error
	return
}

func GetUserByActiveToken(token string) (user *User, err error) {
	err = DB.Where("activae_token=?", token).First(&user).Error
	return
}

func GetUserByRememberMeToken(token string) (user *User, err error) {
	err = DB.Where("remember_me_token=?", token).First(&user).Error
	return
}

func GetUserByWeiboID(id int) (user *User, err error) {
	weibo, _ := GetWeiboByID(id)
	user, _ = GetUserByID(int(weibo.UserID))
	return
}

func GetUserObjectByWeiboID(id int) (user User, err error) {
	weibo, _ := GetWeiboObjectByID(id)
	user, _ = GetUserObjectByID(int(weibo.UserID))
	return
}

func ListUsers() (users []*User, err error) {
	//err = DB.Find(&users, "is_admin=?", true).Error
	err = DB.Order("id").Find(&users).Error
	return
}

func GetUsers(limit int, order string, maps interface{}) (user []User, err error) {
	err = DB.Model(&User{}).Order(order).Limit(limit).Find(&user).Error
	return
}

// func (user *User) GetUserAvatar() string {
// 	if user.Avatar != "" {
// 		return user.Avatar
// 	}
// 	return ""
// }

func CountUsers() (count int) {
	DB.Model(&User{}).Count(&count)
	return
}
