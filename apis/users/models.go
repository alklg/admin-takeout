package users

import "time"

type User struct {
	Id         int       `gorm:"id" json:"id"`
	OpenId     string    `gorm:"open_id" json:"openId"`
	Phone      int64     `gorm:"phone" json:"phone"`
	Name       string    `gorm:"name" json:"name"`
	Username   string    `gorm:"user_name" json:"username"`
	Sex        int       `gorm:"sex" json:"sex"` // 0 is male and 1 is female
	School     int       `gorm:"school" json:"school"`
	Sid        string    `gorm:"sid" json:"sid"`
	Avatar     string    `gorm:"avatar" json:"avatar"`
	CreateTime time.Time `gorm:"create_time" json:"createTime"` // xxxx.xx.xx xx:xx:xx
	Balance    int64     `gorm:"balance" json:"balance"`
	Password   string    `gorm:"password" json:"password"`
	Status     int       `gorm:"status" json:"status"`
	Email      string    `gorm:"email" json:"email"`
}

type ByUsername struct {
	Username string `json:"username"`
	Code     int    `json:"code"`
}

type ByName struct {
	Name string `json:"username"`
	Code int    `json:"code"`
}

type AddUserStruct struct {
	Name       string `json:"name" gorm:"name"`
	School     string `json:"school"`
	Phone      int64  `json:"phone" gorm:"phone"`
	SchoolCode int    `gorm:"school"`
	Sid        string `json:"sid" gorm:"sid"`
}

type SchoolCode struct {
	Code   int    `json:"code" gorm:"code"`
	School string `json:"school" gorm:"school"`
}
