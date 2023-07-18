package usecase

import (
	"fmt"
	"time"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/domain/model"
)

func (takubo *takuboUsecase) Talking() error {
	ln, title, _, _ := takubo.repository.GetCurrentState()
	response, err := takubo.client.FetchSequential(ln, title)
	if err != nil {
		zap.GetLogger().Fatal("DetectのBFFへのリクエストに失敗しました:" + err.Error())
		return err
	}
	zap.GetLogger().Info("DetectのBFFへのリクエストに成功しました:" + fmt.Sprintf("%+v", response))
	// responseに沿って動作を開始
	// DoはWizavoが終わるまでブロッキングする
	if response.State == model.Forget {
		// Forgetの場合はここで終了
		time.Sleep(config.THINK_TIME)
		takubo.voice.Speak("えっとー。")
		takubo.Do(response)
		return nil
	}
	takubo.Do(response)

	takubo.Talking()
	return nil
}
