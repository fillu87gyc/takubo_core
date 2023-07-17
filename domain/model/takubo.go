package model

type State string

const (
	Detect   State = "detect"
	Talking  State = "talking"
	Forget   State = "forget"
	Speaking State = "speaking"
	SpeakEnd State = "speak_end"
)

type Takubo struct {
}
