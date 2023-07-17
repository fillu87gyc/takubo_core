package input

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/domain/model"
)

func (c Client) FetchWithForgetWord(lineNumber int, regularTitle string, targetWord string) (model.Response, error) {
	networkConf := config.NewNetwork()
	url := networkConf.BackendAddr() + "forget?" + fmt.Sprintf("line_number=%d&regular_title=%s&target_word=%s", lineNumber, regularTitle, targetWord)
	zap.GetLogger().Info("リクエストURL:" + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		zap.GetLogger().Error("リクエストの作成に失敗しました:" + err.Error())
		return model.Response{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zap.GetLogger().Error("リクエストの送信に失敗しました:" + err.Error())
		return model.Response{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.GetLogger().Error("レスポンスの読み取りに失敗しました:" + err.Error())
		return model.Response{}, err
	}

	zap.GetLogger().Info("レスポンスボディ:" + string(body))
	responseBody := model.Response{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		zap.GetLogger().Error("レスポンスボディのパースに失敗しました:" + err.Error())
		return model.Response{}, err
	}
	return responseBody, nil
}
