package usecase

import (
	"fmt"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/domain/model"
)

func (takubo *takuboUsecase) SpeechRecog(recog string) error {
	ln, t, speechRecogEnable, state := takubo.repository.GetCurrentState()

	if !speechRecogEnable {
		zap.GetLogger().Info(recog + "LinstenFlagを使ってないので受け入れるよー" + fmt.Sprintf("ln = %d, title = %s", ln, t))
		// return nil
	} else {
		zap.GetLogger().Info(recog + "を音声認識しました")
	}
	switch state {
	case model.Detect:
		return takubo.Detect(recog)
	case model.Talking:
		zap.GetLogger().Error("意図しない状態遷移です。 stateがtalkingなのにSpeechRecogが呼ばれました")
		// panic("意図しない状態遷移です。 stateがtalkingなのにSpeechRecogが呼ばれました")
	case model.Forget:
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		takubo.forgetCond.spokenChannel <- struct{}{}
		return takubo.Forget(recog)
	default:
		zap.GetLogger().Error("意図していないstateです: " + fmt.Sprintf("%+v", state))
		//TODO
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		takubo.forgetCond.spokenChannel <- struct{}{}
		return takubo.Forget(recog)
	}
	return nil
}
