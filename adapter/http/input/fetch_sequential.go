package input

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	errort "github.com/fillu87gyc/lambda-go/errorT"
	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/domain/model"
)

type NameIdentificationResponse struct {
	RegularTitle string `json:"regular_title"`
}

func (c Client) FetchSequential(lineNumber int, regularTitle string) (model.Response, error) {
	networkConf := config.NewNetwork()
	url := networkConf.BackendAddr() + "talking?" + fmt.Sprintf("line_number=%d&regular_title=%s", lineNumber, regularTitle)
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
	//statusコードが200かチェックしてエラーを出す
	if resp.StatusCode != http.StatusOK {
		zap.GetLogger().Error("レスポンスのステータスコードが200ではありませんでした:" + string(body))
		return model.Response{}, errort.ErrOutOfRange
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
