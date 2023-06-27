package main

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

func main() {
	takubo := Takubo{
		Phase: Detect,
		CurrentState: CurrentState{
			Line:  0,
			Title: "",
		},
	}
}
