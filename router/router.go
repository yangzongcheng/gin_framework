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
	"github.com/gin-gonic/gin"
)

/**
 * 初始化方法
 * @Description:
 * @param router
 */
func RouterHandler(router *gin.Engine)  {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong ing",
		})
	})
	router.GET("/test",testHandle)
}

/**
 * test 方法实现
 * @Description:
 * @param c
 */
func testHandle(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "test",
	})
}