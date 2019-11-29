package snapshot

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/consumer/state"
	"github.com/felixvo/lmax/pkg/user"
	"github.com/felixvo/lmax/pkg/warehouse"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisSnapshot(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	snapshotSrv := NewRedisSnapshot(client)
	st := &state.State{}
	st.SetUsers(map[int64]*user.User{
		1: {
			UseID:   1,
			Balance: 1,
		},
		2: {
			UseID:   2,
			Balance: 2,
		},
		3: {
			UseID:   3,
			Balance: 3,
		},
	})
	st.SetItems(map[string]*warehouse.Item{
		"cpu": {
			ID:     "cpu",
			Price:  100,
			Remain: 100,
		},
		"ram": {
			ID:     "ram",
			Price:  50,
			Remain: 150,
		},
	})
	err := Snapshot(*st)
	assert.NoError(t, err)
	restore := state.State{}
	err = Restore(&restore)
	assert.NoError(t, err)
	fmt.Println(restore.LatestEventID)
	for _, u := range restore.Users {
		fmt.Println(u)
	}
	for _, u := range restore.Items {
		fmt.Println(u)
	}
}
