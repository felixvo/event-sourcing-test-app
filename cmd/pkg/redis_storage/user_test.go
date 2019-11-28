package redis_storage

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	repo := NewUserRepository(client)
	userIDs := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	balances := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	err := repo.MultiSetBalance(userIDs, balances)
	assert.NoError(t, err)

	uu, err := repo.MultiGet(userIDs)
	assert.NoError(t, err)
	for _, u := range uu {
		fmt.Println(u)
		assert.Equal(t, u.Balance, balances[u.UseID-1])
	}
}
