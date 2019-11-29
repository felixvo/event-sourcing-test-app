package snapshot

import (
	"bytes"
	"encoding/gob"
	"github.com/go-redis/redis/v7"
	"time"
)

type Snapshot interface {
	Snapshot(t interface{}) error
	Restore(t interface{}) error
}

const (
	snapshotKey = "snapshot"
)

func NewRedisSnapshot(
	client *redis.Client,
) Snapshot {
	return &redisSnapshot{client: client}
}

type redisSnapshot struct {
	client *redis.Client
}

func (r *redisSnapshot) Snapshot(t interface{}) error {
	buf := bytes.Buffer{}
	ecd := gob.NewEncoder(&buf)
	err := ecd.Encode(t)
	if err != nil {
		return err
	}

	_, err = r.client.Set(snapshotKey, buf.Bytes(), time.Hour*1).Result()
	return err
}

func (r *redisSnapshot) Restore(t interface{}) error {
	snap, err := r.client.Get(snapshotKey).Result()
	if err != nil {
		return err
	}
	dcd := gob.NewDecoder(bytes.NewBuffer([]byte(snap)))
	return dcd.Decode(t)
}
