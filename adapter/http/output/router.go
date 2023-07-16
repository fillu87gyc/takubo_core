package output

import (
	http "github.com/fillu87gyc/takubo_core/adapter/http/input"
	mock "github.com/fillu87gyc/takubo_core/adapter/mock/input"
	"github.com/gin-gonic/gin"

	"github.com/fillu87gyc/takubo_core/usecase"
)

func NewHandler(usecase usecase.ITakuboUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type Handler struct {
	usecase usecase.ITakuboUsecase
}

func InitRouter(engine *gin.Engine) *gin.Engine {
	client := http.NewClient()
	repo := mock.NewTakuboRepository()
	usecase := usecase.NewTakuboUsecase(client, repo)
	handler := NewHandler(usecase)
	v1 := engine.Group("/v1")
	{
		v1.GET("/speechrecog", handler.SpeechRecog)
	}
	return engine
}
