package repository

import "github.com/fillu87gyc/takubo_core/domain/model"

type IBackendRepository interface {
	//名寄せする
	FetchNameIdentification(recog string) (string, error)
	// 次の行動をシーケンスにフェッチ
	FetchSequential(lineNumber int, regularTitle string) (model.Response, error)
	// 与えられた文字が起動ワードになっているかという情報を含めて、フェッチ
	FetchWithForgetWord(lineNumber int, regularTitle string, targetWord string) (model.Response, error)
}
