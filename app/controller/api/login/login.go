/**
 * 登录相关逻辑
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package login

import (
  "gin_framework/app/service/api/loginService"
  "github.com/gin-gonic/gin"
)

/**
 * 登录相关
 * @Description:
 * @param c
 */
func LoginMain(c *gin.Context) {
  msg :=loginService.Login()
    c.JSON(200, gin.H{
      "message": msg,
    })
}

/**
 * 注册账号
 * @Description:
 * @param c
 */
func Reg(c *gin.Context)  {
  c.JSON(200, gin.H{
    "message": "注册",
  })
}
