package handler

import (
	"github.com/felixvo/lmax/cmd/pkg/event"
	"github.com/felixvo/lmax/cmd/pkg/state"
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
