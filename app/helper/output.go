/**
 *
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package helper

import (
	"gin_framework/app/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * 返回json
 * @Description:
 * @param persons
 * @return gin.H
 */
func JsonReturn(c *gin.Context,code int,msg string,data interface{}) {
	var msgVal string;
	if msg == "" {
		msgVal = constants.GetCodeMsg(code)
	}else {
		msgVal = msg
	}
	if data == nil {
		data = make(map[string]interface{})
	}
	var reData = map[string]interface{}{
		"code":code,
		"msg":msgVal,
		"data":data,
	}
	c.JSON(http.StatusOK,reData)
	return
}