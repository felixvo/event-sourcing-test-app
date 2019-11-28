package redis_storage

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/pkg/warehouse"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWareHouse(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	wHouse := NewWareHouse(client)
	err := wHouse.UpdateRemains([]*warehouse.Item{
		{
			ID:     "cpu",
			Remain: 10,
		},
		{
			ID:     "ram",
			Remain: 100,
		},
	})
	assert.NoError(t, err)
	items, err := wHouse.GetItems([]string{"cpu", "ram"})
	assert.NoError(t, err)
	for _, it := range items {
		fmt.Println(it)
	}
}
