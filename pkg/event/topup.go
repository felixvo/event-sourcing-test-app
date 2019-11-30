package event

import "github.com/vmihailenco/msgpack/v4"

type TopUp struct {
	*Base
	UserID int64 `json:"user_id"`
	Amount uint  `json:"amount"`
}

func (o *TopUp) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

func (o *TopUp) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}
