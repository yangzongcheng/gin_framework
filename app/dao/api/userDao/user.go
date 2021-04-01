/**
 *
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package userDao

import (
	"gin_framework/app/model"
)

/**
 * 根据用户id获取单条数据
 * @Description:
 * @param uid
 * @return []model.User
 */
func GetUserOne(uid int)[]model.User {

	var user []model.User
	model.Db.First(&user,uid)
	return user
}

/**
 * 获取所有用户数据
 * @Description:
 * @return []model.User
 */
func GetAll()[]model.User  {
	var userAll []model.User
	model.Db.Find(&userAll)
	return userAll
}