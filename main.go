package main

import (
	"fmt"
	"net/http"

	z "github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/adapter/http/output"
	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/lib"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HealthCheck() bool {
	network := config.NewNetwork()
	relatedServices := map[string]string{
		"bff":    network.BackendAddr(),
		"motor":  network.DynamixelAddr(),
		"wizavo": network.WizWebAddr(),
	}

	for name, url := range relatedServices {
		response, err := http.Get(url)
		if err != nil {
			z.GetLogger().Error(fmt.Sprintf("%sが死んでます(%s)", name, url))
			return false
		}
		defer response.Body.Close()
	}
	z.GetLogger().Info(lib.Color("health check ok!!", lib.Green))
	return true
}

func main() {
	logger := z.GetLogger()
	e := gin.New()
	e.Use(MiddleWareLogger(logger))
	e.Use(gin.Recovery())
	takubo := output.InitRouter(e)
	if !HealthCheck() {
		msg := lib.Color("health check failed!!", lib.Red)
		z.GetLogger().Error(msg)
		return
	}
	takubo.Run(":8080")
}
func MiddleWareLogger(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		l.Info("",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
		c.Next()
	}
}
