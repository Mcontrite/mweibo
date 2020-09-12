package model

import (
	"fmt"
	"strconv"
)

type Follower struct {
	ID         uint `gorm:"not null;primary_key;auto_increment"`
	UserID     uint `gorm:"not null" sql:"index"` // 多对多：被关注者
	FollowerID uint `gorm:"not null" sql:"index"` // 多对多：粉丝
}

func ListUserFollowings(userid int) (followers []*User, err error) {
	sqlstr := "inner join followers on users.id=followers.user_id "
	err = DB.Model(&User{}).Joins(sqlstr).Where("followers.follower_id=?", userid).Order("id").Find(&followers).Error
	return
}

func ListUserFollowers(userid int) (followers []*User, err error) {
	sqlstr := "inner join followers on users.id=followers.follower_id "
	err = DB.Model(&User{}).Joins(sqlstr).Where("followers.user_id=?", userid).Order("id").Find(&followers).Error
	return
}

func ListUserFollowingIDs(userid int) (ids []uint) {
	followings, _ := ListUserFollowings(userid)
	for _, v := range followings {
		ids = append(ids, v.ID)
	}
	return
}

func IsFollowing(selfid, userid int) bool {
	flwingsid := ListUserFollowingIDs(selfid)
	for _, v := range flwingsid {
		if v == uint(userid) {
			return true
		}
	}
	return false
}

func FollowSomeUsers(selfid uint, userids ...uint) error {
	sqlstr := "insert into followers (follower_id,user_id) values "
	for k, v := range userids {
		sqlstr += fmt.Sprintf("(%d,%d)", selfid, v)
		if k < len(userids)-1 {
			sqlstr += ","
		}
	}
	return DB.Exec(sqlstr).Error
}

func UnfollowSomeUsers(selfid uint, userids ...uint) error {
	sqlstr := fmt.Sprintf("delete from followers where follower_id=%d and user_id in (", selfid)
	for k, v := range userids {
		sqlstr += strconv.Itoa(int(v))
		if k < len(userids)-1 {
			sqlstr += ","
		}
	}
	sqlstr += ")"
	return DB.Exec(sqlstr).Error
}

func CountFollowings(userid int) (count int, err error) {
	sqlstr := "inner join followers on users.id=followers.user_id "
	err = DB.Model(&User{}).Joins(sqlstr).Where("followers.follower_id=?", userid).Count(&count).Error
	return
}

func CountFollowers(userid int) (count int, err error) {
	sqlstr := "inner join followers on users.id=followers.follower_id "
	err = DB.Model(&User{}).Joins(sqlstr).Where("followers.user_id=?", userid).Count(&count).Error
	return
}
