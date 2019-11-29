package event

import "fmt"

func New(t Type) (Event, error) {
	b := &Base{
		Type: t,
	}
	switch t {
	case OrderType:
		return &OrderEvent{
			Base: b,
		}, nil
	case TopUpType:
		return &TopUp{
			Base: b,
		}, nil
	case AddItemType:
		return &AddItem{
			Base: b,
		}, nil
	}

	return nil, fmt.Errorf("type %v not supported", t)
}
