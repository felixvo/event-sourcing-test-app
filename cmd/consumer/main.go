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
	rr, err := client.XRange("orders", "-", "+").Result()
	if err != nil {
		panic(err)
	}
	for _, r := range rr {
		o := order.Order(r.Values)
		d, err := o.Data()
		fmt.Println(r.ID, " ", d, " ", err)
		_, err = client.XDel("orders", r.ID).Result()
		fmt.Println("del:", err)
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
