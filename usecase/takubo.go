package usecase

import (
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
	motor      repository.IMotorRepository
	voice      repository.IVoiceRepository
	forgetCond forgetCond
}

const (
	//定型文
	FIXEDTEXT = iota
	//疑問文を繰り返す
	INTERROGATIVE_SENTENCE
	//小出しする
	SMALLAMOUNTS
)

type forgetCond struct {
	closeChannel  chan struct{}
	spokenChannel chan struct{}
	question      string
	bestAnswer    string
}

func NewTakuboUsecase(
	client repository.IBackendRepository,
	repo repository.ITakuboRepository,
	motor repository.IMotorRepository,
	voice repository.IVoiceRepository,
) ITakuboUsecase {
	return &takuboUsecase{
		client:     client,
		repository: repo,
		motor:      motor,
		voice:      voice,
	}
}

func (takubo *takuboUsecase) SetState(state model.State) error {
	return takubo.repository.SetCurrentState(state)
}

func (takubo *takuboUsecase) Speak(text string) error {
	takubo.repository.SetCurrentState(model.Speaking)
	zap.GetLogger().Debug(lib.Color(" Wizavo発話[["+text+"]]終了まで待機", lib.Yellow))
	err := takubo.voice.Speak(text)
	takubo.repository.SetCurrentState(model.SpeakEnd)
	return err
}
func (takubo *takuboUsecase) BeginTalkingHint(cancel chan struct{}, spoken chan struct{}) {
	behaviorPattern := 0
	progressCount := 0
	period := 10 * time.Second
	ticker := time.NewTicker(period)
	defer ticker.Stop()
	defer close(spoken)
	for {
		select {
		case <-cancel:
			zap.GetLogger().Info("BeginTalkingHint: cancel")
			return

		case <-ticker.C:
			select {
			case <-spoken:
				zap.GetLogger().Info("BeginTalkingHint: spoken")
				for i := 0; i < len(spoken); i++ {
					<-spoken
				}
			default:
				if (behaviorPattern % 3) == SMALLAMOUNTS {
					progressCount++
				}
				takubo.forgetHint(behaviorPattern, progressCount)
				behaviorPattern++
			}
		}
	}
}
