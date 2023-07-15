package repository

type DetectParam struct {
	Title string `json:"title"`
}

type TalkingParam struct {
	LineNumber   int    `json:"line_number"`
	RegularTitle string `json:"regular_title"`
}

type ForgetParam struct {
	LineNumber   int    `json:"line_number" binding:"required"`
	RegularTitle string `json:"regular_title" binding:"required"`
	TargetWord   string `json:"target_word" binding:"required"`
}
