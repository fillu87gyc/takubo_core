package input

import (
	"github.com/fillu87gyc/takubo_core/domain/repository"
)

func NewClient() repository.IBackendRepository {
	return &Client{}
}

type Client struct {
}
