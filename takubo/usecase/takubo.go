package usecase

var takubo *domain.Takubo

func GetInstance() *domain.Takubo {
	if takubo == nil {
		takubo = &domain.Takubo{}
	}
	return takubo
}
