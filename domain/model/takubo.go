package model

type State int

var _takubo *Takubo

func GetTakuboSingleton() *Takubo {
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
