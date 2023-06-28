package model

import (
	"github.com/fillu87gyc/takubo_core/takubo/domain/repository"
)

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
	Repo         Repo
}
