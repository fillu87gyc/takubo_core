package usecase

import (
	"github.com/fillu87gyc/takubo_core/domain"
)

var takubo *domain.Takubo

func GetInstance() *domain.Takubo {
	if takubo == nil {
		takubo = &domain.Takubo{}
	}
	return takubo
}
