package output

import (
	"github.com/fillu87gyc/takubo_core/adapter/dynamixel"
	http "github.com/fillu87gyc/takubo_core/adapter/http/input"
	mock "github.com/fillu87gyc/takubo_core/adapter/mock/input"
	"github.com/fillu87gyc/takubo_core/adapter/wizavo"
	"github.com/fillu87gyc/takubo_core/domain/model"
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
	wizavo := wizavo.NewVoice()
	dynamixel := dynamixel.NewMotor()
	usecase := usecase.NewTakuboUsecase(client, repo, dynamixel, wizavo)
	handler := NewHandler(usecase)
	handler.usecase.SetState(model.Detect)

	engine.GET("/speechrecog/:recog", handler.SpeechRecog)

	return engine
}
