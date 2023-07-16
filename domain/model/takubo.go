package model

type State int

const (
	Detect State = iota
	Talking
	Forget
)

type Takubo struct {
}
