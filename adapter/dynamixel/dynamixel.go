package dynamixel

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/domain/repository"
)

func NewMotor() repository.IMotorRepository {
	return &Motor{}
}

type Motor struct {
}

func (*Motor) SetPosture(poses []repository.PoseBehavior) error {
	// GETリクエストを作成
	network := config.NewNetwork()

	bytecode, _ := json.Marshal(poses)
	data := string(bytecode)
	zap.GetLogger().Info("dynamixelに送信するデータ:" + data)
	param := url.QueryEscape(data)

	url := network.DynamixelAddr() + "/drive/" + param
	resp, err := http.Get(url)
	if err != nil {
		zap.GetLogger().Error("dynamixelとの接続に失敗しました:" + err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}
