package mock

import (
	"fmt"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/domain/model"
	"github.com/fillu87gyc/takubo_core/domain/repository"
)

type takuboRepository struct {
	State                model.State
	NextAccessLineNumber int
	Title                string
	ListenFlag           bool
}

// SetTitle implements repository.ITakuboRepository.
func (t *takuboRepository) SetTitle(title string) error {
	t.Title = title
	return nil
}

// GetCurrentState implements repository.ITakuboRepository.
func (t *takuboRepository) GetCurrentState() (int, string, bool, model.State) {
	return t.NextAccessLineNumber, t.Title, t.ListenFlag, t.State
}

// SetCurrentState implements repository.ITakuboRepository.
func (t *takuboRepository) SetCurrentState(state model.State) error {
	t.State = state
	switch state {
	case model.Detect:
		t.NextAccessLineNumber = 0
		t.ListenFlag = false
		zap.GetLogger().Info("response = detect" + fmt.Sprintf("%+v", t))
	case model.Talking:
		t.NextAccessLineNumber++
		t.ListenFlag = false
		zap.GetLogger().Info("response = talking" + fmt.Sprintf("%+v", t))
	case model.Forget:
		t.ListenFlag = true
	default:
		zap.GetLogger().Fatal("意図していないstateです" + fmt.Sprintf("%+v", t))
		panic("Undefined state")
	}
	return nil
}

var _takuboRepository *takuboRepository

func NewTakuboRepository() repository.ITakuboRepository {
	if _takuboRepository == nil {
		_takuboRepository = &takuboRepository{}
	}
	return _takuboRepository
}
