package main

import (
	z "github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/adapter/http/output"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger := z.GetLogger()
	e := gin.New()
	e.Use(MiddleWareLogger(logger))
	e.Use(gin.Recovery())

	takubo := output.InitRouter(e)
	takubo.Run(":8080")
}
func MiddleWareLogger(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		l.Info("",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}
}
