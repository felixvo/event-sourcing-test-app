package main

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/pkg/order"
	"github.com/go-redis/redis/v7"
)

func main() {
	client, err := newRedisClient()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 1; i++ {

		strCMD := client.XAdd(&redis.XAddArgs{
			Stream: "orders",
			Values: order.Order{}.SetData(&order.Data{
				CustomerID:     "felix",
				ItemIDs:        []string{"cpu", "ram"},
				ItemQuantities: []int{1, 2},
			}),
		})
		s, err := strCMD.Result()
		fmt.Println(s)
		fmt.Println(err)
	}

}

func newRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-14450.c1.asia-northeast1-1.gce.cloud.redislabs.com:14450",
		Password: "37uaACndCvuQ1heADnHkishnAhMmosWq", // no password set
		DB:       0,                                  // use default DB
	})

	_, err := client.Ping().Result()
	return client, err

}
