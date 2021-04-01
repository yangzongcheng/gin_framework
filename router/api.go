/**
 *
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package router

import (
	"gin_framework/app/controller/api/login"
	"gin_framework/app/controller/api/user"
	"github.com/gin-gonic/gin"
)

/**
 * 初始化方法
 * @Description:
 * @param router gin Default 方法 指针
 */
func ApiHandler(router *gin.Engine)  {
	//路由分组
	api :=router.Group("/api")
	{
		api.GET("/api", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api ing",
			})
		})
		// 版本v1
		v1 := api.Group("/v1")
		{
			//登录接口
			v1.GET("/login",login.LoginMain)
			//注册接口
			v1.GET("/reg",login.Reg)
			//获取用户信息
			v1.GET("/user",user.GetUserInfo)
			//获取所有用户
			v1.GET("/userAll",user.GetUserAll)
		}

		// 版本v2
		v2 := api.Group("/v2")
		{
			v2.GET("/login", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "v2 login",
				})
			})

		}
	}


}
