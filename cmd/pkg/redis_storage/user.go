package redis_storage

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/pkg/user"
	"github.com/go-redis/redis/v7"
	"strconv"
)

const (
	userHashKey    = "users"
	balanceHashKey = "balance"
)

type userRespostory struct {
	client *redis.Client
}

func NewUserRepository(client *redis.Client) user.Repository {
	return &userRespostory{
		client: client,
	}
}

func (u *userRespostory) Get(userID int64) (*user.User, error) {
	vv, err := u.client.HGetAll(userKey(userID)).Result()
	if err != nil {
		return nil, err
	}
	return parseUser(userID, vv)
}
func (u *userRespostory) MultiGet(userIDs []int64) ([]*user.User, error) {
	keys := userKeys(userIDs)
	pipeline := u.client.Pipeline()
	cmds := make([]*redis.StringStringMapCmd, len(keys))
	for i, k := range keys {
		cmds[i] = pipeline.HGetAll(k)
	}
	_, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	// parse data after exe pipeline
	rs := make([]*user.User, len(keys))
	for i, c := range cmds {
		data, err := c.Result()
		if err != nil {
			return nil, err
		}
		usr, err := parseUser(userIDs[i], data)
		if err != nil {
			return nil, err
		}
		rs[i] = usr
	}
	return rs, nil
}

func (u *userRespostory) MultiSetBalance(userIDs []int64, balance []int64) error {
	pipeline := u.client.TxPipeline()
	for i, userID := range userIDs {
		pipeline.HSet(userKey(userID), balanceHashKey, balance[i])
	}
	_, err := pipeline.Exec()
	return err
}

func userKey(userID int64) string {
	return fmt.Sprintf("%s.%v", userHashKey, userID)
}
func userKeys(userIDs []int64) []string {
	rs := make([]string, len(userIDs))
	for i, uID := range userIDs {
		rs[i] = userKey(uID)
	}
	return rs
}

func parseUser(userID int64, v map[string]string) (*user.User, error) {
	usr := user.User{
		UseID: userID,
	}
	bStr, exist := v[balanceHashKey]
	if !exist {
		return &usr, nil
	}
	b, err := strconv.ParseInt(bStr, 10, 64)
	if err != nil {
		return nil, err
	}
	usr.Balance = b
	return &usr, nil
}
