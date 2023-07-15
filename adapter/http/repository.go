package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fillu87gyc/takubo_core/config"
	"github.com/fillu87gyc/takubo_core/domain/model"
	"github.com/fillu87gyc/takubo_core/domain/repository"
)

type Repository struct {
	networkConf config.Network
}

func NewRepository() repository.IRepository {
	return &Repository{
		networkConf: config.NewNetwork(),
	}
}

func (r *Repository) IsWakeWord(word model.DetectParams) (model.Response, error) {
	url := r.networkConf.BackendAddr() + "detect" + "?recog=" + word.Recog
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("リクエストの作成に失敗しました:", err)
		return model.Response{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("リクエストの送信に失敗しました:", err)
		return model.Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("レスポンスの読み取りに失敗しました:", err)
		return model.Response{}, err
	}

	fmt.Println("レスポンスボディ:", string(body))
	responseBody := model.Response{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		fmt.Println("レスポンスボディのパースに失敗しました:", err)
		return model.Response{}, err
	}
	return responseBody, nil
}

func (r *Repository) FetchNextBehavior(takubo model.TalkingParams) (model.Response, error) {
	// 実際の処理を実装する
	return model.Response{}, nil
}

func (r *Repository) IsCorrectWord(cond model.ForgetParams) (model.Response, error) {
	// 実際の処理を実装する
	return model.Response{}, nil
}
