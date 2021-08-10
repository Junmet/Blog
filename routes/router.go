//路由的入口文件
package router

import (
	"Blog/api/v1"
	"Blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	// 路由初始化
	r := gin.Default()
	router := r.Group("api/v1")
	{
		//测试路由
		//router.GET("hello",func (c *gin.Context){
		//	c.JSON(http.StatusOK,gin.H{
		//		"msg":"Ok",
		//	})
		//})

		//用户模块路由接口
		router.GET("user:/id", v1.UserExist)
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		//分类模块路由接口

		//文章模块路由接口
	}
	_ = r.Run(utils.HttpPort)
}
