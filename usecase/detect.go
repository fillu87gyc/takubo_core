package usecase

import (
	"fmt"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/domain/model"
)

func (takubo *takuboUsecase) Detect(recog string) error {
	regTitle, err := takubo.client.FetchNameIdentification(recog)
	if err != nil {
		zap.GetLogger().Fatal("DetectのBFFへのリクエストに失敗しました:" + err.Error())
		return err
	}
	zap.GetLogger().Info("DetectのBFFへのリクエストに成功しました:" + fmt.Sprintf("%+v", regTitle))
	takubo.repository.SetTitle(regTitle)
	takubo.SetState(model.Detect) // ここだけResponseを参照しない
	if regTitle == "" {
		//名寄せが失敗していた
		return nil
	}
	//名寄せ成功処理
	zap.GetLogger().Info("名寄せ成功 title =  " + regTitle + "。 Detect終了 talkingに移動します")
	takubo.Talking()
	return nil
}
