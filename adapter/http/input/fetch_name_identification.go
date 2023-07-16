package input

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/config"
)

// 与えられた文字がregualr_titleに近いか判定する
func (c Client) FetchNameIdentification(recog string) (string, error) {
	networkConf := config.NewNetwork()
	url := networkConf.BackendAddr() + "detect" + "?recog=" + recog
	zap.GetLogger().Info("リクエストURL:" + url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		zap.GetLogger().Fatal("リクエストの作成に失敗しました:" + err.Error())
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zap.GetLogger().Error("リクエストの送信に失敗しました:" + err.Error())
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.GetLogger().Error("レスポンスの読み取りに失敗しました:" + err.Error())
		return "", err

	}

	zap.GetLogger().Info("レスポンスボディ:" + string(body))
	responseBody := NameIdentificationResponse{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		zap.GetLogger().Error("レスポンスボディのパースに失敗しました:" + err.Error())
		return "", err
	}
	return responseBody.RegularTitle, nil
}
