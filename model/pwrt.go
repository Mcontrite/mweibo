package model

import (
	"mweibo/utils"
	"time"
)

type PasswordReset struct {
	Email   string `gorm:"not null" sql:"index"`
	Token   string `gorm:"not null" sql:"index"`
	ResetAt time.Time
}

func GetPwrtByEmail(email string) (pwrt *PasswordReset, err error) {
	err = DB.Where("email=?", email).First(&pwrt).Error
	return
}

func GetPwrtByToken(token string) (pwrt *PasswordReset, err error) {
	err = DB.Where("token=?", token).First(&pwrt).Error
	return
}

func DeletePwrtByEmail(email string) (err error) {
	pwrt := &PasswordReset{}
	err = DB.Where("email=?", email).Delete(&pwrt).Error
	return
}

//func (pwrt *PasswordReset) DeletePwrtByToken(token string) (err error) {
func DeletePwrtByToken(token string) (err error) {
	pwrt := &PasswordReset{}
	err = DB.Where("token=?", token).Delete(&pwrt).Error
	return
}

func (pwrt *PasswordReset) CreatePwrt() (err error) {
	token := string(utils.CreateRandomBytes(30))
	// 如果已经存在则先删除
	oldpwrt, _ := GetPwrtByEmail(pwrt.Email)
	// err = pwrt.DeletePwrtByEmail(oldpwrt.Email)
	err = DeletePwrtByEmail(oldpwrt.Email)
	pwrt.Token = token
	err = DB.Create(&pwrt).Error
	return
}
