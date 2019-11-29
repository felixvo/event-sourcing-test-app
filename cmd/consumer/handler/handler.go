package handler

import (
	"github.com/felixvo/lmax/pkg/event"
)

type Handler interface {
	Handle(e event.Event) error
}
