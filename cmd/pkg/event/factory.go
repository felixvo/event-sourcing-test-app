package event

import "fmt"

func New(t Type) (Event, error) {
	switch t {
	case OrderType:
		return &OrderEvent{
			Base: &Base{},
		}, nil
	case TopUpType:
		return &TopUp{
			Base: &Base{},
		}, nil
	case AddItemType:
		return &AddItem{
			Base: &Base{},
		}, nil
	}

	return nil, fmt.Errorf("type %v not supported", t)
}
