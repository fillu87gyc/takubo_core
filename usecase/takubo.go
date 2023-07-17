package usecase

import (
	"fmt"
	"time"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/domain/model"
	"github.com/fillu87gyc/takubo_core/domain/repository"
	"github.com/fillu87gyc/takubo_core/lib"
)

type ITakuboUsecase interface {
	// 与えられた文字が起動ワードになっているか、問い合わせて行動に移す
	Detect(recog string) error
	Talking() error
	Forget(targetWord string) error
	// stateがtalkingなら takubo.NextAccessLineNumberをインクリメント
	SetState(state model.State) error
	Do(model.Response) error
	SpeechRecog(recog string) error
}

type takuboUsecase struct {
	client     repository.IBackendRepository
	repository repository.ITakuboRepository
}

// SpeechRecog implements ITakuboUsecase.
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

func NewTakuboUsecase(client repository.IBackendRepository, repo repository.ITakuboRepository) ITakuboUsecase {
	return &takuboUsecase{
		client:     client,
		repository: repo,
	}
}

func (takubo *takuboUsecase) Do(r model.Response) error {
	zap.GetLogger().Info(fmt.Sprintf(lib.Color("[[wizavo]]: %s || ", lib.Yellow)+
		lib.Color("[[dynamixel]]: %+v", lib.Cyan)+
		lib.Color("[[state]]: %s", lib.Green),
		r.Text, r.Behavior, r.State))

	takubo.repository.SetCurrentState(model.Speaking)
	zap.GetLogger().Info("Wizavo発話中")
	for _, b := range []int{0, 1, 2} {
		time.Sleep(1.0 * time.Second)
		zap.GetLogger().Debug(lib.Color(" Wizavo発話[["+r.Text+"]]終了まで", lib.Yellow) + fmt.Sprintf("%d秒", 3-b))
	}

	// time.Sleep(config.WaitTimeDuringTurn * time.Second)
	for _, b := range []int{0, 1, 2, 3} {
		time.Sleep(1.0 * time.Second)
		zap.GetLogger().Debug(lib.Color("===========次の行動までのマチ===========||"+fmt.Sprintf("%d秒", 4-b), lib.Cyan))
	}
	takubo.repository.SetCurrentState(r.State)
	return nil
}

// Detect implements ITakuboUsecase.
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

// Forget implements ITakuboUsecase.
func (takubo *takuboUsecase) Forget(targetWord string) error {
	ln, title, _, _ := takubo.repository.GetCurrentState()
	response, err := takubo.client.FetchWithForgetWord(ln, title, targetWord)
	if err != nil {
		zap.GetLogger().Fatal("DetectのBFFへのリクエストに失敗しました:" + err.Error())
		return err
	}
	zap.GetLogger().Info("DetectのBFFへのリクエストに成功しました:" + fmt.Sprintf("%+v", response))
	// stateごとの処理
	takubo.repository.SetCurrentState(response.State)
	// responseに沿って動作を開始
	// DoはWizavoが終わるまでブロッキングする
	takubo.Do(response)
	if response.State == model.Forget {
		// Forgetの場合はここで終了
		return nil
	}
	takubo.repository.SetCurrentState(model.SpeakEnd)
	takubo.Talking()
	return nil
}

func (takubo *takuboUsecase) SetState(state model.State) error {
	return takubo.repository.SetCurrentState(state)
}

// Talking implements ITakuboUsecase.
func (takubo *takuboUsecase) Talking() error {
	ln, title, _, _ := takubo.repository.GetCurrentState()
	response, err := takubo.client.FetchSequential(ln, title)
	if err != nil {
		zap.GetLogger().Fatal("DetectのBFFへのリクエストに失敗しました:" + err.Error())
		return err
	}
	zap.GetLogger().Info("DetectのBFFへのリクエストに成功しました:" + fmt.Sprintf("%+v", response))
	// stateごとの処理
	takubo.repository.SetCurrentState(response.State)
	// responseに沿って動作を開始
	// DoはWizavoが終わるまでブロッキングする
	takubo.Do(response)
	if response.State == model.Forget {
		// Forgetの場合はここで終了
		return nil
	}
	takubo.Talking()
	return nil
}
