package service

import (
	"github.com/fillu87gyc/takubo_core/domain/model"
	"github.com/fillu87gyc/takubo_core/domain/repository"
)

// interface
type ITakuboService interface {
	// 与えられた文字が起動ワードになっているか、問い合わせにいく
	IsWakeWord(word string) (bool, error)
	// 次の行動をフェッチして実行する
	FetchAndDoBehavior() (model.Behavior, error)
	// ステータスを更新する。
	// talkingが入っている場合にはtakubo.LineNumをインクリメント
	SetState(state model.State) (model.State, error)
	GetStoryCondition() (int, string, error)
}

// struct that meets interface
type takuboService struct {
	repo repository.Repository
}

func NewOldStoryService(sr repository.Repository) ITakuboService {
	return &takuboService{
		repo: sr,
	}
}

func (s *takuboService) IsWakeWord(word string) (bool, error) {
	return s.repo.IsWakeWord(word)
}

func (s *takuboService) FetchAndDoBehavior() (model.Behavior, error) {
	return s.repo.Fetch()
}

func (s *takuboService) SetState(state model.State) (model.State, error) {
	return s.repo.SetState(state)
}

func (s *takuboService) GetStoryCondition() (int, string, error) {
	return s.repo.GetStoryCondition()
}
