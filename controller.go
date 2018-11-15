package ging

import (
	"log"
	"strings"
)

import (
	"github.com/gin-gonic/gin"
)

type IController interface {
	Action(action func(ctx *gin.Context) )
}
