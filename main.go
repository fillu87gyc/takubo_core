package main

import (
	"github.com/fillu87gyc/takubo_core/takubo/domain/model"
	"github.com/fillu87gyc/takubo_core/takubo/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	takubo := usecase.GetTakuboSingleton()
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/speechrecog", speechrecog)
}

func speechrecog(c *gin.Context) {
	takubo := usecase.GetTakuboSingleton()

    if takubo.Phase == model.Detect{
        takubo.
    }

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
