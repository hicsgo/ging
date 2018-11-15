package ging

import (
	"fmt"
)

/* ================================================================================
 * json结果数据结果模型
 * email   : golang123@outlook.com
 * author  : hicsgo
 * ================================================================================ */
type (
	JsonData struct {
		Code int32
		Msg  string
		Data interface{}
	}

	PagingData struct {
		JsonData
		Paging *Paging
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置错误状态信息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (result *JsonData) SetError(err error) {
	if customErr, ok := err.(*CustomError); ok {
		result.Code = customErr.Code
		result.Msg = customErr.Msg
	} else {
		msg := fmt.Sprintf("%s", err.Error())
		result.Code = 111
		result.Msg = msg
	}
}
