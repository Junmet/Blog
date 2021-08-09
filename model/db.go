//db入口文件
package model

import (
	"Blog/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"

)

var db *gorm.DB
var err error

//连接配置数据库
func InitDb()	{
	//第一个参数是Db的类型，第二个参数是一个地址
	db,err = gorm.Open(utils.Db,fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		))
		if err !=nil {
			fmt.Println("连接数据库失败，请检查参数:",err)
		}
		// 设置禁用表名复数形式属性为 true，`User` 的表名将是 `user`
		//false  `User` 的表名将是 `users`
		db.SingularTable(true)

		//自动迁移模型
		db.AutoMigrate(&User{},&Article{},&Category{})

		//做连接配置参数  连接池
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		db.DB().SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		db.DB().SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		//这个连接时间要是大于gin框架的连接时间就容易抛出错误。
		db.DB().SetConnMaxLifetime(10*time.Second) //10秒

		//关闭
		//_ = db.Close()
}