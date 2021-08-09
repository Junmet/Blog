package model

import "github.com/jinzhu/gorm"

//User结构体
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username"`  //主动性能
	Password string `gorm:"type:varchar(20);not null " json:"password"`
	Role int `gorm:"type:int " json:"role"`
}