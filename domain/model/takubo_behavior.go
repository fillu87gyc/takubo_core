package model

type Response struct {
	Text       string     `json:"text"`
	State      State      `json:"state"`
	Behavior   []Behavior `json:"behavior"`
	BestAnswer string     `json:"best_answer"`
}

type Behavior struct {
	DoTime float64 `json:"do_time"`
	Pose   string  `json:"pose"`
}
