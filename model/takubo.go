package model

import repo "github.com/fillu87gyc/takubo_core/domain/repository"

type State int

var _takubo *Takubo

func GetTakuboSingleton(repo repo.Repository) *ITakubo {
	if _takubo == nil {
		_takubo = &Takubo{}
	}
	return _takubo
}

const (
	Detect State = iota
	Talking
	Forget
)

type Takubo struct {
	State      State
	LineNumber int
	Title      string
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

type ITakubo interface {
	// 与えられた文字が起動ワードになっているか、問い合わせにいく
	IsWakeWord(word string) (bool, error)
	// 次の行動をフェッチして実行する
	FetchAndDoBehavior() (Behavior, error)
	// ステータスを更新する。
	// talkingが入っている場合にはtakubo.LineNumをインクリメント
	SetState(state State) (State, error)
	GetStoryCondition() (int, string, error)
}
