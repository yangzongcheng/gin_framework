/**
 *
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package constants

//状态码
const(
	Ok=200
	Error=201
	ParamError=202
)

/**	@var 状态码对应提示*/
var ErrorMsg = map[int]string{
	Ok:"操作成功",
	Error:"操作失败",
	ParamError:"参数错误",
}

/**
 * 通过状态码获取提示信息
 * @Description:
 * @param code
 * @return string
 */
func GetCodeMsg(code int) string {
	return ErrorMsg[code];
}
