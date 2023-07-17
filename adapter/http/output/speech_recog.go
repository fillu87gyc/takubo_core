package output

import (
	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SpeechRecog(c *gin.Context) {
	// speechrecog
	// Queryからrecogを取得
	recog := c.Param("recog")
	if recog == "" {
		zap.GetLogger().Error("recog is empty")
		c.JSON(400, gin.H{
			"message": "recog is empty",
		})
		return
	}
	// usecaseを呼び出す
	c.JSON(200, gin.H{
		"message": "OK",
	})
	c.Abort()
	go h.usecase.SpeechRecog(recog)
}
