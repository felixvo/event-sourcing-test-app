package event

import "github.com/vmihailenco/msgpack/v4"

type OrderEvent struct {
	*Base
	UserID         int64    `json:"user_id"`
	ItemIDs        []string `json:"item_i_ds"`
	ItemQuantities []uint   `json:"item_quantities"`
}

func (o *OrderEvent) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

func (o *OrderEvent) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}
