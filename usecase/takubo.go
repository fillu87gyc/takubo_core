package usecase

import (
	"github.com/fillu87gyc/takubo_core/domain/model"
	"github.com/fillu87gyc/takubo_core/domain/repository"
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
