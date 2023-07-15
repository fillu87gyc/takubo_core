package repository

import (
	"github.com/fillu87gyc/takubo_core/domain/model"
)

type IRepository interface {
	IsWakeWord(cond DetectParam) (model.Response, error)
	FetchNextBehavior(cond TalkingParam) (model.Response, error)
	IsCorrectWord(cond ForgetParam) (model.Response, error)
}
