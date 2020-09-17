/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/7/22 11:06 上午
*/

package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ApiSuccess = 0
	ApiErrMsg  = 500 // 显示错误信息
)

type ApiRet struct {
	Code int32                  `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}

type Controller interface {
	JsonSuccess(c *gin.Context, data map[string]interface{})
	JsonError(c *gin.Context, code int32, msg string)
}

type BaseController struct {

}

func (ctrl *BaseController) JsonHandler(c *gin.Context, data map[string]interface{}, code int32, msg string) {
	//c.Header("Access-Control-Allow-Origin", "*")


	if code!=ApiSuccess {
		ctrl.JsonError(c, code, msg)
	}else {
		ctrl.JsonSuccess(c, data)
	}

	//ret := &ApiRet{
	//	Code: ApiSuccess,
	//	Msg:  "",
	//	Data: data,
	//}
	//
	////c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	//
	//c.JSON(http.StatusOK, ret)
}

func (ctrl *BaseController) JsonSuccess(c *gin.Context, data map[string]interface{}) {
	ret := &ApiRet{
		Code: ApiSuccess,
		Msg:  "",
		Data: data,
	}

	//c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	c.JSON(http.StatusOK, ret)
}

func (ctrl *BaseController) JsonError(c *gin.Context, code int32, msg string) {
	ret := &ApiRet{
		Code: code,
		Msg:  msg,
	}

	//c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	//c.BindHeader()

	c.JSON(http.StatusOK, ret)
}