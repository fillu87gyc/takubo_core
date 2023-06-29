package main

import (
	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/model"
	"github.com/gin-gonic/gin"
)

func main() {
	takubo := model.GetTakuboSingleton()
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/speechrecog", speechrecog)
}

func speechrecog(c *gin.Context) {
	takubo := model.GetTakuboSingleton()
	apiVersion := config.ApiVersion
	if takubo.State == model.Detect {
		takubo.RequestDetect(apiVersion + "/detect/")
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
