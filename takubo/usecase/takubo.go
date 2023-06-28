package usecase

import (
	"github.com/fillu87gyc/takubo_core/takubo/domain/model"
)

// type takuboUsecase struct {
// 	svc service.TakuboService
// }

var _takubo *model.Takubo

// func NewTakuboUsecase(repo Repo) {
// 	_takubo = &model.Takubo{
// 		Repo: repo,
// 	}
// }

func GetTakuboSingleton() *model.Takubo {
	if _takubo == nil {
		_takubo = &model.Takubo{}
	}
	return _takubo
}

type ITakuboUsecase interface {
}
