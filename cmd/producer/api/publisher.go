package api

import (
	"fmt"
	"github.com/felixvo/lmax/pkg/event"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"net/http"
)

type eventData struct {
	Type           string   `json:"type"`
	UserID         int64    `json:"user_id"`
	Amount         uint     `json:"amount"`
	ItemID         string   `json:"item_id"`
	Count          uint     `json:"count"`
	ItemIDs        []string `json:"item_ids"`
	ItemQuantities []uint   `json:"item_quantities"`
}

func NewPublisherHandler(client *redis.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		e := eventData{}
		if err := c.BindJSON(&e); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		newID, err := publishEvent(client, &e)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id": newID,
		})
	}
}
func publishEvent(client *redis.Client, eData *eventData) (string, error) {
	eType := event.Type(eData.Type)
	var e event.Event
	switch eType {
	case event.TopUpType:
		e = &event.TopUp{
			Base: &event.Base{
				Type: event.TopUpType,
			},
			UserID: eData.UserID,
			Amount: eData.Amount,
		}
		break
	case event.AddItemType:
		e = &event.AddItem{
			Base: &event.Base{
				Type: event.AddItemType,
			},
			ItemID: eData.ItemID,
			Count:  eData.Count,
		}
		break
	case event.OrderType:
		e = &event.OrderEvent{
			Base: &event.Base{
				Type: event.OrderType,
			},
			UserID:         eData.UserID,
			ItemIDs:        eData.ItemIDs,
			ItemQuantities: eData.ItemQuantities,
		}
		break
	default:
		return "", fmt.Errorf("event type not exist")
	}

	strCMD := client.XAdd(&redis.XAddArgs{
		Stream: "orders",
		Values: map[string]interface{}{
			"type": eData.Type,
			"data": e,
		},
	})
	newEventID, err := strCMD.Result()
	return newEventID, err
}
