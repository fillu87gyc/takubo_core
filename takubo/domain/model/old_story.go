package model

type Body struct {
	Text     string
	IsForget bool
	Answers  []string
}

type Story struct {
	Title       string
	BodyList    []Body
	ScreenTitle string
	WakeWords   []string
}
