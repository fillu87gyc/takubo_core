package usecase

import (
	"fmt"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/domain/model"
)

func (takubo *takuboUsecase) SpeechRecog(recog string) error {
	ln, t, speechRecogEnable, state := takubo.repository.GetCurrentState()
	if !speechRecogEnable {
		zap.GetLogger().Info(recog + "を音声認識しましたが、SpeechRecogが有効ではありません" + fmt.Sprintf("ln = %d, title = %s", ln, t))
		return nil
	} else {
		zap.GetLogger().Info(recog + "を音声認識しました")
	}
	switch state {
	case model.Detect:
		return takubo.Detect(recog)
	case model.Talking:
		panic("意図しない状態遷移です。 stateがtalkingなのにSpeechRecogが呼ばれました")
	case model.Forget:
		return takubo.Forget(recog)
	default:
		zap.GetLogger().Fatal("意図していないstateです" + fmt.Sprintf("%+v", takubo))
		panic("Undefined state")
	}
}
