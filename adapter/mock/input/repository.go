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

var tmpListenFlag = false
var tmpState = model.Detect

// SetCurrentState implements repository.ITakuboRepository.
func (t *takuboRepository) SetCurrentState(state model.State) error {
	switch state {
	case model.Detect:
		t.State = state
		t.NextAccessLineNumber = 0
		t.ListenFlag = true
		zap.GetLogger().Info("set state response = detect//  " + fmt.Sprintf("%+v", t))
	case model.Talking:
		t.State = state
		t.NextAccessLineNumber++
		t.ListenFlag = false
		zap.GetLogger().Info("set state response = talking//  " + fmt.Sprintf("%+v", t))
	case model.Forget:
		t.State = state
		t.ListenFlag = true
		zap.GetLogger().Info("set state response = forget//  " + fmt.Sprintf("%+v", t))
	case model.Speaking:
		tmpListenFlag = t.ListenFlag
		tmpState = t.State
		t.ListenFlag = false
		t.State = state
		zap.GetLogger().Info(" ----speaking----  ")
		return nil
	case model.SpeakEnd:
		if t.State != model.Speaking {
			zap.GetLogger().Warn("意図していないstateです。speakじゃない場所からspeakEndにきました" + fmt.Sprintf("%+v", t))
			// panic("Undefined state")
		}
		t.ListenFlag = tmpListenFlag
		t.State = tmpState
		zap.GetLogger().Info(" ----speak_end----  ")
		return nil
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
