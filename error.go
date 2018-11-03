package ging

/* ================================================================================
 * 自定义错误数据域结构
 * email   : golang123@outlook.com
 * author  : hicsgo
 * ================================================================================ */
type CustomError struct {
	Code int32
	Msg  string
}

func (err CustomError) Error() string {
	return err.Msg
}
