package repository

import "github.com/fillu87gyc/takubo_core/model"

type Repository interface {
	IsWakeWord(word string) (bool, error)
	Fetch() (model.Behavior, error)
}
