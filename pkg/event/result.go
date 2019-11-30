package event

type HandleEventResult struct {
	Event Event  `json:"event"`
	Err   string `json:"err"`
}
