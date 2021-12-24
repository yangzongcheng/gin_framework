/**
 *
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package user

import (
	"gin_framework/app/dao/api/userDao"
	"gin_framework/app/helper"
	"github.com/gin-gonic/gin"
	"gin_framework/app/constants"
	//"gin_framework/app/helper"
	"log"
	"strconv"
)

//返回数据
var reData struct{
	Name string `json:"name"` //定义json返回格式
	Id uint64 `json:"id"`
	Age uint8 `json:"age"`
}

type Person struct{
	Name string `json:"name"` //定义json返回格式
	Id uint64 `json:"id"`
	Age uint8 `json:"age"`
}



/**
 * 获取所有
 * @Description:
 * @param c
 */
func GetUserAll(c *gin.Context) {
	data :=userDao.GetAll()
	userData := make([]Person, 0)
	for _,j  := range data  {
		var person Person
		person.Id = j.ID
		person.Name = j.Name
		person.Age = j.Age
		userData = append(userData, person)
	}
	helper.JsonReturn(c,constants.Ok,"ok",userData)
	return
}

/**
 * 获取单个用户信息
 * @Description:
 * @param c
 */
func GetUserInfo(c *gin.Context)  {
	paramId :=c.Query("id")
	id,err:=strconv.Atoi(paramId)
	if err != nil {
		//异常处理
		log.Panicln("err:", err.Error())
	}
	info:=userDao.GetUserOne(id)[0]
	reData.Id = info.ID
	reData.Name = info.Name
	reData.Age = info.Age
	helper.JsonReturn(c,constants.Ok,"",reData)
	return
}
