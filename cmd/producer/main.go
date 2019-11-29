package main

import (
	"fmt"
	"github.com/felixvo/lmax/pkg/event"
	"github.com/go-redis/redis/v7"
	"math/rand"
	"time"
)

func main() {
	client, err := newRedisClient()
	if err != nil {
		panic(err)
	}

	d := time.After(time.Second * 30)

	go func() {
		for {
			Topup(client)
			AddItem(client)
			MakeOrders(client)
		}
	}()
	<-d
}
func Topup(client *redis.Client) {
	for i := 0; i < 10; i++ {
		userID := int64(i + 1)
		strCMD := client.XAdd(&redis.XAddArgs{
			Stream: "orders",
			Values: map[string]interface{}{
				"type": string(event.TopUpType),
				"data": &event.TopUp{
					Base: &event.Base{
						Type: event.TopUpType,
					},
					UserID: userID,
					Amount: 500,
				},
			},
		})
		newID, err := strCMD.Result()
		if err != nil {
			fmt.Printf("topup error:%v\n", err)
		} else {
			fmt.Printf("topup success for user:%v offset:%v\n", userID, newID)
		}
	}
}
func AddItem(client *redis.Client) {
	for i := 0; i < 10; i++ {
		itemID := []string{"cpu", "ram", "hdd", "ssd"}[rand.Intn(4)]
		count := uint(rand.Intn(50))
		strCMD := client.XAdd(&redis.XAddArgs{
			Stream: "orders",
			Values: map[string]interface{}{
				"type": string(event.AddItemType),
				"data": &event.AddItem{
					Base: &event.Base{
						Type: event.AddItemType,
					},
					ItemID: itemID,
					Count:  count,
				},
			},
		})
		newID, err := strCMD.Result()
		if err != nil {
			fmt.Printf("add item error:%v\n", err)
		} else {
			fmt.Printf("ad item success itemID:%v count:%v offset:%v\n", itemID, count, newID)
		}
	}
}

func MakeOrders(client *redis.Client) {
	for i := 0; i < 10; i++ {
		itemID := []string{"cpu", "ram", "hdd", "ssd"}[rand.Intn(4)]
		count := uint(rand.Intn(50))
		userID := int64(rand.Intn(10) + 1)
		strCMD := client.XAdd(&redis.XAddArgs{
			Stream: "orders",
			Values: map[string]interface{}{
				"type": string(event.OrderType),
				"data": &event.OrderEvent{
					Base: &event.Base{
						Type: event.OrderType,
					},
					UserID:         userID,
					ItemIDs:        []string{itemID},
					ItemQuantities: []uint{count},
				},
			},
		})
		newID, err := strCMD.Result()
		if err != nil {
			fmt.Printf("add item error:%v\n", err)
		} else {
			fmt.Printf("make order success userID:%v itemID:%v count:%v offset:%v\n", userID, itemID, count, newID)
		}
		//time.Sleep(time.Second * 2)
	}
}
func newRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		//Addr:     "redis-14450.c1.asia-northeast1-1.gce.cloud.redislabs.com:14450",
		//Password: "37uaACndCvuQ1heADnHkishnAhMmosWq", // no password set
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()
	return client, err

}
