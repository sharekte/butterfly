package eng

import (
	"butterfly/pkg/e"
	"github.com/gin-gonic/gin"
)

type G struct {
	g gin.Context
}

type Gin = *gin.Context

type GEng = *gin.Engine

func (g *G) Send(httpCode, code int, data interface{}) {
	g.g.JSON(httpCode, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
