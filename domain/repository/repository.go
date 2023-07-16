package repository

import "github.com/fillu87gyc/takubo_core/domain/model"

type ITakuboRepository interface {
	// NextAccessLineNumber, Titleを返す
	GetCurrentState() (int, string, bool, model.State)
	// Responseで受け取ったStateによってごねごねする
	SetCurrentState(state model.State) error
	SetTitle(title string) error
}
