package v1

import "github.com/gin-gonic/gin"

//c *gin.Context是上下文，要有这个参数，路由router的 v1.AddUser 才不会报错。

//查询用户是否存在
func UserExist(c *gin.Context) {

}

//添加用户
func AddUser(c *gin.Context) {

}

//查询单个用户

//查询用户列表
func GetUsers(c *gin.Context) {

}

//编辑用户
func EditUser(c *gin.Context) {

}

//删除用户
func DeleteUser(c *gin.Context) {

}
