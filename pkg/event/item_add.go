package event

import "github.com/vmihailenco/msgpack/v4"

type AddItem struct {
	*Base
	ItemID string `json:"item_id"`
	Count  uint   `json:"count"`
}

func (o *AddItem) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

func (o *AddItem) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}
