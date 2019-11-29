package handler

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/consumer/state"
	"github.com/felixvo/lmax/pkg/event"
	"github.com/felixvo/lmax/pkg/user"
)

type topupHandler struct {
	state *state.State
}

func NewTopupHandler(state *state.State) Handler {
	return &topupHandler{
		state: state,
	}
}

func (h *topupHandler) Handle(e event.Event) error {
	topup, ok := e.(*event.TopUp)
	defer func() {
		h.state.LatestEventID = topup.GetID()
	}()
	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	u, exist := h.state.Users[topup.UserID]
	if !exist { // should have an event to create user before use
		u = &user.User{
			UseID:   topup.UserID,
			Balance: 0,
		}
		h.state.Users[topup.UserID] = u
	}

	u.Balance += topup.Amount

	fmt.Printf("completed topup %+v \n", topup)
	return nil
}
