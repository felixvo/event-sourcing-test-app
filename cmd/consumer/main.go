package main

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/pkg/order"
)

const (
	OrderStream = "orders"
)

func main() {
	client, err := newRedisClient()
	if err != nil {
		panic(err)
	}
	rr, err := client.XRange(OrderStream, "-", "+").Result()
	if err != nil {
		panic(err)
	}
	for _, r := range rr {
		o := order.Order(r.Values)
		d, err := o.Data()
		fmt.Println(r.ID, " ", d, " ", err)

		_, err = client.XDel(OrderStream, r.ID).Result()
		fmt.Println("del:", err)
	}
}
