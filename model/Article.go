//文章
package model

import "github.com/jinzhu/gorm"

type Article struct {
	Category Category
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Cid int `gorm:"type:int;not null" json:"cid"`	//关联Category的id
	Desc string `gorm:"type:varchar(200)" json:"desc"`  //描述
	Content string `gorm:"type:longtext" json:"content"`
	Img string `gorm:"type:varchar(100)" json:"img"`
}