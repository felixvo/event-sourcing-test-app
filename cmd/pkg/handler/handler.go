package handler

import "github.com/felixvo/lmax/cmd/pkg/event"

type Handler interface {
	Handle(e event.Event) error
}
