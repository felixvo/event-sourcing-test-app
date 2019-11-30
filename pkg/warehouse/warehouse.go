package warehouse

type Item struct {
	ID     string `json:"id"`
	Price  uint   `json:"price"`
	Remain uint   `json:"remain"`
}
