package handler

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/pkg/event"
	"github.com/felixvo/lmax/cmd/pkg/state"
	"github.com/felixvo/lmax/cmd/pkg/warehouse"
)

type orderHandler struct {
	state *state.State
}

func NewOrderHandler(state *state.State) Handler {
	return &orderHandler{
		state: state,
	}
}

func (h *orderHandler) Handle(e event.Event) error {
	defer func() {
		h.state.LatestEventID = e.GetID()
	}()
	orderEvent, ok := e.(*event.OrderEvent)
	if !ok {
		return fmt.Errorf("invalid event type")
	}
	if len(orderEvent.ItemIDs) <= 0 {
		return fmt.Errorf("invalid items")
	}
	u := h.state.GetUserByID(orderEvent.UserID)
	items := h.state.GetItems(orderEvent.ItemIDs)

	total, err := total(items, orderEvent.ItemQuantities)
	if err != nil {
		return err
	}
	if total > u.Balance {
		return fmt.Errorf("not enough cash")
	}
	// handle
	u.Balance = u.Balance - total
	updateItemQuantities(items, orderEvent.ItemQuantities)
	fmt.Printf("completed order %+v \n", orderEvent)
	return nil
}
func total(items []*warehouse.Item, quantities []uint) (uint, error) {
	var total uint
	for i, item := range items {
		total += item.Price
		if item.Remain < quantities[i] {
			return 0, fmt.Errorf("%s is out of stock", item.ID)
		}
	}
	return total, nil
}
func updateItemQuantities(items []*warehouse.Item, quantities []uint) {
	for i, item := range items {
		item.Remain = item.Remain - quantities[i]
	}
}
