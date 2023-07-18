package usecase

import (
	"fmt"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/domain/model"
)

func (takubo *takuboUsecase) Forget(targetWord string) error {
	ln, title, _, _ := takubo.repository.GetCurrentState()
	response, err := takubo.client.FetchWithForgetWord(ln, title, targetWord)
	if err != nil {
		zap.GetLogger().Fatal("DetectのBFFへのリクエストに失敗しました:" + err.Error())
		return err
	}
	zap.GetLogger().Info("DetectのBFFへのリクエストに成功しました:" + fmt.Sprintf("%+v", response))
	takubo.Do(response)
	if response.State == model.Forget {
		// Forgetの場合はここで終了
		return nil
	}
	takubo.Talking()
	return nil
}
