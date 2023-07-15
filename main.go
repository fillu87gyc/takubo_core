package main

import (
	"github.com/fillu87gyc/takubo_core/adapter/http"
	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/domain/model"
	"github.com/gin-gonic/gin"
)

func main() {
	takubo := model.GetTakuboSingleton()
	r := gin.Default()
	// 受け取りサイドの処理
	v1 := r.Group("/v1")
	v1.GET("/speechrecog", speechrecog)
	repo := http.NewRepository()

}

const STATUS_OK = 200

func speechrecog(c *gin.Context) {
	takubo := model.GetTakuboSingleton()
	apiVersion := config.ApiVersion
	if takubo.State == model.Detect {
		takubo.RequestDetect(apiVersion + "/detect/")
	}

	c.JSON(STATUS_OK, gin.H{
		"message": "pong",
	})
}
