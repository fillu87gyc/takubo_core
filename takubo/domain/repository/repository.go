package repository

import "github.com/fillu87gyc/takubo_core/takubo/domain/model"

type OldStoryRepository interface {
	RequestStory(lineNumber int, regTitle string) (model.Body, error)
	RequestRegularTitle(title string) (string, error)
	CheckCorrectWord(lineNumber int, word string, title string) (bool, error)
}
