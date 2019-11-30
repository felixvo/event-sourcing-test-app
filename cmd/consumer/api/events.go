package api

import (
	"github.com/felixvo/lmax/pkg/event"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ListEvents(result chan *event.HandleEventResult) gin.HandlerFunc {
	events := []*event.HandleEventResult{}
	go func() {
		for {
			newEvent := <-result
			events = append(events, newEvent)
			if len(events) > 10 {
				newEvents := make([]*event.HandleEventResult, 10)
				copy(newEvents, events[1:])
				events = newEvents
			}
			time.Sleep(time.Microsecond * 100)
		}
	}()
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, events)
	}
}
