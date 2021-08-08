//路由的入口文件
package router
import (
	"Blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)
func InitRouter(){
	gin.SetMode(utils.AppMode)
	// 路由初始化
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("hello",func (c *gin.Context){
			c.JSON(http.StatusOK,gin.H{
				"msg":"Ok",
			})
		})
	}
	_ = r.Run(utils.HttpPort)
}
