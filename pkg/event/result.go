package event

type HandleEventResult struct{
	Event Event `json:"event"`
	Err   error `json:"err"`
}
