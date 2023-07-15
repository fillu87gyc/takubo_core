package model

type Response struct {
	Text     string     `json:"text"`
	State    string     `json:"state"`
	Behavior []Behavior `json:"behavior"`
}

type Behavior struct {
	DoTime int64  `json:"do_time"`
	Pose   string `json:"pose"`
}