package handler

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/pkg/event"
	"github.com/felixvo/lmax/cmd/pkg/state"
	"github.com/felixvo/lmax/cmd/pkg/warehouse"
	"math/rand"
)

type itemAddHandler struct {
	state *state.State
}

func NewAddItemHandler(state *state.State) Handler {
	return &itemAddHandler{
		state: state,
	}
}

func (h *itemAddHandler) Handle(e event.Event) error {
	addItem, ok := e.(*event.AddItem)
	defer func() {
		h.state.LatestEventID = addItem.GetID()
	}()
	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	i, exist := h.state.Items[addItem.ItemID]
	if !exist {
		i = &warehouse.Item{
			ID:     addItem.ItemID,
			Price:  uint(rand.Intn(100)),
			Remain: uint(rand.Intn(200)),
		}
		h.state.Items[addItem.ItemID] = i
	}
	i.Remain += addItem.Count

	fmt.Printf("completed add item %+v \n", addItem)
	return nil
}
