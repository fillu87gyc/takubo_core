package usecase

import (
	"fmt"
	"net/http"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/domain/model"
)

func (takubo *takuboUsecase) SpeechRecog(recog string) error {
	// ln, t, speechRecogEnable, state := takubo.repository.GetCurrentState()
	_, _, _, state := takubo.repository.GetCurrentState()
	if state == model.Talking {
		return nil
	}
	// if !speechRecogEnable {
	// 	zap.GetLogger().Info(recog + "LinstenFlagを使ってないので受け入れるよー" + fmt.Sprintf("ln = %d, title = %s", ln, t))
	// 	// return nil
	// } else {
	// 	zap.GetLogger().Info(lib.Color((recog + "を音声認識しました"), lib.Red))
	// }
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
	}
	return nil
}

func (takubo *takuboUsecase) SetRecognitionState(isStart bool) error {
	var s string
	if s = "stop"; isStart {
		s = "start"
	}
	network := config.NewNetwork()
	url := network.RecogAddr() + "/" + s
	resp, err := http.Get(url)
	if err != nil {
		zap.GetLogger().Error("speech recogとの接続に失敗しました:" + err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}
