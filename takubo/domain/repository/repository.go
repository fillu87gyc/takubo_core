package repository

import (
	back "github.com/takubo_behavior_gin_server/old_story/usecase/model"
)

type OldStoryRepository interface {
	RequestStory(lineNumber int, regTitle string) (string, error)
}
