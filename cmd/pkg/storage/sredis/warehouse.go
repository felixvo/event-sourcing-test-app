package sredis

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/pkg/warehouse"
	"github.com/go-redis/redis/v7"
	"strconv"
)

const (
	warehouseHashKey   = "warehouses"
	warehouseRemainKey = "remain"
)

type warehouseRedis struct {
	client *redis.Client
}

func NewWareHouse(client *redis.Client) warehouse.Repository {
	return &warehouseRedis{
		client: client,
	}
}
func (w *warehouseRedis) GetItems(itemIDs []string) ([]*warehouse.Item, error) {
	keys := warehouseKeys(itemIDs)
	pipeline := w.client.Pipeline()
	cmds := make([]*redis.StringStringMapCmd, len(keys))
	for i, k := range keys {
		cmds[i] = pipeline.HGetAll(k)
	}
	_, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	// parse data after exe pipeline
	rs := make([]*warehouse.Item, len(keys))
	for i, c := range cmds {
		data, err := c.Result()
		if err != nil {
			return nil, err
		}
		item, err := parseItem(itemIDs[i], data)
		if err != nil {
			return nil, err
		}
		rs[i] = item
	}
	return rs, nil
}

func (w *warehouseRedis) UpdateRemains(items []*warehouse.Item) error {
	pipeline := w.client.TxPipeline()
	for _, item := range items {
		pipeline.HSet(warehouseKey(item.ID), warehouseRemainKey, item.Remain)
	}
	_, err := pipeline.Exec()
	return err
}
func warehouseKey(itemID string) string {
	return fmt.Sprintf("%s.%s", warehouseHashKey, itemID)
}
func warehouseKeys(itemIDs []string) []string {
	rs := make([]string, len(itemIDs))
	for i, itemID := range itemIDs {
		rs[i] = warehouseKey(itemID)
	}
	return rs
}

func parseItem(itemID string, v map[string]string) (*warehouse.Item, error) {
	item := warehouse.Item{
		ID:     itemID,
		Remain: 0,
	}
	bStr, exist := v[warehouseRemainKey]
	if !exist {
		return &item, nil
	}
	b, err := strconv.ParseInt(bStr, 10, 64)
	if err != nil {
		return nil, err
	}
	item.Remain = b
	return &item, nil
}
