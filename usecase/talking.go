package usecase

import (
	"fmt"
	"time"

	errort "github.com/fillu87gyc/lambda-go/errorT"
	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/domain/model"
)

func (takubo *takuboUsecase) Talking() error {
	ln, title, _, _ := takubo.repository.GetCurrentState()
	response, err := takubo.client.FetchSequential(ln, title)
	if err != nil {
		if err == errort.ErrOutOfRange {
			// ErrOutOfRangeはちょっとだけ許す
			zap.GetLogger().Warn("DetectのBFFへのリクエストに失敗しました。多分終わりまでいったのかな？:" + err.Error())
			takubo.SetRecognitionState(true)
			takubo.repository.SetCurrentState(model.Detect) //初期状態に返してあげる
			if e := takubo.Do(model.Response{
				Text:       "",
				State:      model.Detect,
				Behavior:   []model.Behavior{{Pose: "track", DoTime: 1.0}},
				BestAnswer: "",
			}); e != nil {
				zap.GetLogger().Error("初期状態に戻すのに失敗しました。")
			}
			return nil
		}
		panic("BFFとの疎通に失敗" + err.Error())
	}
	zap.GetLogger().Info("DetectのBFFへのリクエストに成功しました:" + fmt.Sprintf("%+v", response))
	// responseに沿って動作を開始
	// DoはWizavoが終わるまでブロッキングする
	if response.State != model.Forget {
		// forget状態に移動していないのでそのまま次へ
		if err := takubo.Do(response); err != nil {
			zap.GetLogger().Error("Doの実行に失敗しました。")
		}
		takubo.Talking()
		return nil
	}
	// Forgetの場合はforgetルーチンスタート
	time.Sleep(config.THINK_TIME)
	takubo.Speak("えっとー。")

	if err := takubo.Do(response); err != nil {
		zap.GetLogger().Error("Doの実行に失敗しました。")
	}
	takubo.SetRecognitionState(true)
	takubo.forgetCond.bestAnswer = response.BestAnswer
	takubo.forgetCond.question = response.Text
	takubo.forgetCond.closeChannel = make(chan struct{}, 100)
	takubo.forgetCond.spokenChannel = make(chan struct{}, 100) //十分大きい値
	go takubo.BeginTalkingHint(
		takubo.forgetCond.closeChannel,
		takubo.forgetCond.spokenChannel,
	)
	return nil
}
