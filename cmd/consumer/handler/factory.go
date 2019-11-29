package handler

import (
	"github.com/felixvo/lmax/cmd/consumer/state"
	"github.com/felixvo/lmax/pkg/event"
)

func HandlerFactory(st *state.State) func(t event.Type) Handler {

	return func(t event.Type) Handler {
		switch t {
		case event.OrderType:
			return NewOrderHandler(st)
		case event.TopUpType:
			return NewTopupHandler(st)
		case event.AddItemType:
			return NewAddItemHandler(st)
		}
		return NewLogHandler(st)
	}
}
