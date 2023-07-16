package output

import "github.com/gin-gonic/gin"

func (h *Handler) SpeechRecog(c *gin.Context) {
	// speechrecog
	// Queryからrecogを取得
	recog := c.Query("recog")
	// usecaseを呼び出す
	h.usecase.SpeechRecog(recog)
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
