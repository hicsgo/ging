package ging

import (
	"github.com/gin-gonic/gin"
	"github.com/hicsgo/ging/render"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 结果接口
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type IActionResult interface {
	Render()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 视图结果数据结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type (
	ActionResult struct {
		Context     *gin.Context
		ContentData interface{}
		ContentType string
		StatusCode  int
		IsAbort     bool
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 渲染Json字符串数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (result *ActionResult) Json(arg ...interface{}) {
	if result.Context.IsAborted() {
		return
	}
	render.Json(result.Context, arg...)
}
