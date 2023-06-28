package model

type StatePhase int

const (
	Detect StatePhase = iota
	Talking
	Forget
)

type CurrentState struct {
	Line  uint
	Title string
}

type Takubo struct {
	Phase        StatePhase
	CurrentState CurrentState
}
