package handler

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/consumer/state"
	"github.com/felixvo/lmax/pkg/event"
)

type logHandler struct {
	state *state.State
}

func NewLogHandler(
	state *state.State,
) Handler {
	return &logHandler{state: state}
}

func (h *logHandler) Handle(e event.Event) error {
	defer func() {
		h.state.LatestEventID = e.GetID()
	}()
	fmt.Printf("new event:%+v\n", e)
	return nil
}
