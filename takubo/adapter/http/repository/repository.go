package repository

import (
	"github.com/fillu87gyc/takubo_core/takubo/domain/model"
)

type checkCorrectWord struct {
	LineNumber   int    `json:"line_number" binding:"required"`
	RegularTitle string `json:"regular_title" binding:"required"`
	TargetWord   string `json:"target_word" binding:"required"`
}

type checkCorrectWordResponseErr struct {
	Error  string `json:"error"`
	Detail string `json:"detail"`
	Text   string `json:"text"`
}

type checkCorrectWordResponse struct {
	IsCorrect bool `json:"is_correct"`
}

type requestRegularTitle struct {
	Title string `json:"title" required:"true"`
}

type requestStory struct {
	LineNumber   int    `json:"line_number" binding:"required"`
	RegularTitle string `json:"regular_title" binding:"required"`
}
