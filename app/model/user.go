/**
 *
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package model

//用户表
type User struct {
	BaseModel
	ID uint64 `gorm:"primary_key"`
	Name string
	Age uint8
}

/**
 * 重新定义表名
 * @Description:
 * @receiver User
 * @return string
 */
func (User) TableName() string {
	return "t_user"
}

/**
 * todo
 * @Description:
 * @return []User
 */
func GetIdOne() []User {
	var users []User
	Db.Find(&users)
	return users
}
