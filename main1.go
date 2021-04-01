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
	"fmt"
	"github.com/spf13/viper"
)

/**
 * 程序入口
 * @Description:
 */
func main()  {
	viper.SetConfigFile("./config/.env") // 指定配置文件

	viper.SetConfigType("env") // 如果配置文件的名称中没有扩展名，则需要配置此项
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s n", err))
	}
	fmt.Println(viper.GetString("mysql.userName"))

}






