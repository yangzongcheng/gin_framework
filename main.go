/**
 * gin框架实战
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package main

import (
	 "gin_framework/app/model"
	_"gin_framework/config" //加载配置文件
	"gin_framework/router"
	"github.com/gin-gonic/gin"
)



/**
 * 程序入口
 * @Description:
 */
func main()  {
	model.DbInitMain();//初始化db
	routerObj := gin.Default()
	router.RouterHandler(routerObj)//初始化Router路由
	router.ApiHandler(routerObj)//初始化Api路由
	defer model.Db.Close()
	routerObj.Run("127.0.0.1:8080") // 监听并在 0.0.0.0:8080 上启动服务
}




