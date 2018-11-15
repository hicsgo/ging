package result

import (
	"github.com/gin-gonic/gin"
)

import (
	"github.com/hicsgo/ging"
)

/* ================================================================================
 * Json结果
 * email   : golang123@outlook.com
 * author  : hicsgo
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 *  视图结果数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type (
	jsonResult struct {
		ging.ActionResult
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * json结果
 * args case: statusCode | isAbort | statusCode,isAbort
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func JsonResult(ctx *gin.Context, data interface{}, args ...interface{}) ging.IActionResult {
	result := &jsonResult{}
	result.Context = ctx
	result.ContentData = data
	result.StatusCode = 200
	result.ContentType = "json"
	argsCount := len(args)
	if argsCount > 0 {
		if argsCount == 1 {
			switch  value := args[0].(type) {
			case int:
				result.StatusCode = value
			case bool:
				result.IsAbort = value
			}
		}
	} else if argsCount == 2 {
		if value, isOk := args[0].(int); isOk {
			result.StatusCode = value
		}
		if value, isOk := args[1].(bool); isOk {
			result.IsAbort = value
		}
	}
	return result
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 渲染结果
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (r *jsonResult) Render() {
	r.Json(r.ContentData, r.StatusCode, r.IsAbort)
}
