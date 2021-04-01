/**
 * 初始化链接数据库
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

//赋值到链接对象
var Db *gorm.DB


/**
 * 初始化db
 * @Description:
 */
func DbInitMain()  {

	var err error
	var userName  = viper.GetString("mysql.UserName")
	var host = viper.GetString("mysql.Host")
	var port = viper.GetString("mysql.port")
	var dbName =viper.GetString("mysql.DbName")
	var charset =viper.GetString("mysql.Charset")
	var passWord  = viper.GetString("mysql.password")

	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",userName,passWord,host,port,dbName,charset)
	//Db声明了全局变量不能在使用 := 的用法
	Db, err = gorm.Open("mysql", dbLink)
	//defer Db.Close()
	if err != nil {
		log.Panicln("mysql content err:", err.Error())
	}
	//禁用复数s
	//Db.SingularTable(true)
}