package model

type DetectParams struct {
	Recog string `json:"recog"`
}

type TalkingParams struct {
	Title      string `json:"title"`
	LineNumber string `json:"line_number"`
}

type ForgetParams struct {
	Word       string `json:"word"`
	Title      string `json:"title"`
	LineNumber string `json:"line_number"`
}
