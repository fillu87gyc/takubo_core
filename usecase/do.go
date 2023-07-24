package usecase

import (
	"fmt"
	"time"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/domain/model"
	"github.com/fillu87gyc/takubo_core/domain/repository"
	"github.com/fillu87gyc/takubo_core/lib"
)

func (takubo *takuboUsecase) Do(r model.Response) error {
	zap.GetLogger().Info(fmt.Sprintf(lib.Color("[[wizavo]]: %s || ", lib.Yellow)+
		lib.Color("[[dynamixel]]: %+v ||", lib.Cyan)+
		lib.Color("[[state]]: %s", lib.Green),
		r.Text, r.Behavior, r.State))

	repoModel := make([]repository.PoseBehavior, len(r.Behavior))
	for i, adapterModel := range r.Behavior {
		p, err := repository.ParsePresetPose(adapterModel.Pose)
		if err != nil {
			zap.GetLogger().Error("未定義のポーズがBFFから送信されたため終了します:" + adapterModel.Pose + ":" + err.Error())
		}
		repoModel[i] = repository.PoseBehavior{
			DoTime:  adapterModel.DoTime,
			Pose:    p,
			NodFlag: r.State == model.Talking,
		}
	}
	zap.GetLogger().Debug(lib.Color(fmt.Sprintf("セットポーズ %+v", repoModel), lib.Green))
	takubo.motor.SetPosture(repoModel)
	takubo.Speak(r.Text)
	takubo.repository.SetCurrentState(r.State)
	time.Sleep(config.WaitTimeDuringTurn)
	return nil
}
