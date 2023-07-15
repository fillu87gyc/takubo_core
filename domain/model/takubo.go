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
