package event

import "github.com/vmihailenco/msgpack/v4"

type OrderEvent struct {
	*Base
	UserID         int64
	ItemIDs        []string
	ItemQuantities []uint
}

func (o *OrderEvent) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

func (o *OrderEvent) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}
