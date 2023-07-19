package wizavo

import (
	"net/http"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/domain/repository"
)

func NewVoice() repository.IVoiceRepository {
	return &Voice{}
}

type Voice struct {
}

func (v *Voice) Speak(text string) error {
	// GETリクエストを作成
	network := config.NewNetwork()
	url := network.WizWebAddr() + "/speak/" + text
	resp, err := http.Get(url)
	if err != nil {
		zap.GetLogger().Error("wizavoとの接続に失敗しました:" + err.Error())
		panic(err)
	}
	defer resp.Body.Close()
	return nil
}
