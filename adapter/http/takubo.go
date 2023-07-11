package http

import (
	"github.com/fillu87gyc/takubo_core/domain/model"
)

struct Takubo struct {
}

// FetchAndDoBehavior implements ITakubo.
func (Takubo) FetchAndDoBehavior() (model.Behavior, error) {
	panic("unimplemented")
}

// GetStoryCondition implements ITakubo.
func (Takubo) GetStoryCondition() (int, string, error) {
	panic("unimplemented")
}

// IsWakeWord implements ITakubo.
func (Takubo) IsWakeWord(word string) (bool, error) {
	panic("unimplemented")
}

// SetState implements ITakubo.
func (Takubo) SetState(state model.State) (model.State, error) {
	panic("unimplemented")
}
