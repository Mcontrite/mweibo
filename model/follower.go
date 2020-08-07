package model

import (
	"fmt"
	"strconv"
)

type Follower struct {
	ID         uint
	UserID     uint `gorm:"not null" sql:"index"`
	FollowerID uint `gorm:"not null" sql:"index"`
}

func FollowUser(userid uint, followersid ...uint) error {
	sqlstr := "insert into followers (user_id,follower_id) values "
	for k, v := range followersid {
		sqlstr += fmt.Sprintf("(%d,%d)", userid, v)
		if k < len(followersid)-1 {
			sqlstr += ","
		}
	}
	return DB.Exec(sqlstr).Error
}

func UnfollowUser(userid uint, followersid ...uint) error {
	sqlstr := fmt.Sprintf("delete from followers where follower_id=%d and user_id in (", userid)
	for k, v := range followersid {
		sqlstr += strconv.Itoa(int(v))
		if k < len(followersid)-1 {
			sqlstr += ","
		}
	}
	sqlstr += ")"
	return DB.Exec(sqlstr).Error
}

func GetFollowersCount(userid int) (count int, err error) {
	sqlstr := "inner join followers on users.id=followers.follower_id"
	err = DB.Model(&User{}).Joins(sqlstr).Where("followers.user_id=?", userid).Count(&count).Error
	return
}

func ListUserFollowers(userid int) (followers []*Follower, err error) {
	sqlstr := "inner join followers on users.id=followers.follower_id"
	err = DB.Model(&User{}).Joins(sqlstr).Where("followers.user_id=?", userid).Order("id").Find(&followers).Error
	return
}

func GetFollowingsCount(userid int) (count int, err error) {
	sqlstr := "inner join followers on users.id=followers.user_id"
	err = DB.Model(&User{}).Joins(sqlstr).Where("followers.follower_id=?", userid).Count(&count).Error
	return
}

func ListUserFollowings(userid int) (followers []*Follower, err error) {
	sqlstr := "inner join followers on users.id=followers.user_id"
	err = DB.Model(&User{}).Joins(sqlstr).Where("followers.follower_id=?", userid).Order("id").Find(&followers).Error
	return
}

func ListUserFollowingsID(userid int) (ids []uint) {
	followings, _ := ListUserFollowings(userid)
	for _, v := range followings {
		ids = append(ids, v.ID)
	}
	return
}

func IsFollowing(ctxuserid, userid int) bool {
	flwingsid := ListUserFollowingsID(ctxuserid)
	for _, v := range flwingsid {
		if uint(userid) == v {
			return true
		}
	}
	return false
}
