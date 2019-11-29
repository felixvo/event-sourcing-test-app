package main

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/pkg/event"
	"github.com/felixvo/lmax/cmd/pkg/handler"
	"github.com/felixvo/lmax/cmd/pkg/snapshot"
	"github.com/felixvo/lmax/cmd/pkg/state"
	"github.com/felixvo/lmax/cmd/pkg/user"
	"github.com/felixvo/lmax/cmd/pkg/warehouse"
	"github.com/go-redis/redis/v7"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	OrderStream = "orders"
)

var (
	startTime time.Time
	counter   int32 = 0
)

func main() {
	client, err := newRedisClient()
	if err != nil {
		panic(err)
	}
	snapshotSrv := snapshot.NewRedisSnapshot(client)
	st := initialState(snapshotSrv)
	go exeSnapshot(st, snapshotSrv)
	startTime = time.Now()
	consumeEvents(client, st, handler.HandlerFactory(st))
	quit := make(chan bool)
	<-quit
}
func exeSnapshot(st *state.State, snapshotSrv snapshot.Snapshot) {
	ticker := time.Tick(time.Second * 10)
	for {
		select {
		case <-ticker:
			err := snapshotSrv.Snapshot(st)
			if err != nil {
				fmt.Println("snapshot failed:", err)
				break
			}
			fmt.Println("snapshot success:", st.LatestEventID, " at ", time.Now())
		}

	}
}

func consumeEvents(client *redis.Client, st *state.State, handlerFactory func(t event.Type) handler.Handler) {

	for {
		start := "-"
		if len(st.LatestEventID) > 0 {
			splitted := strings.Split(st.LatestEventID, "-")
			counter, _ := strconv.Atoi(splitted[1])
			start = fmt.Sprintf("%s-%v", splitted[0], counter+1)

		}
		rr, err := client.XRange(OrderStream, start, "+").Result()
		if err != nil {
			panic(err)
		}
		if len(rr) == 0 {
			fmt.Println(time.Since(startTime), "total:", counter)
			return
		}
		for _, r := range rr {
			func() {
				atomic.AddInt32(&counter, 1)
				t := r.Values["type"].(string)
				e, err := event.New(event.Type(t))
				if err != nil {
					panic(err)
				}
				err = e.UnmarshalBinary([]byte(r.Values["data"].(string)))
				if err != nil {
					client.XDel("orders", r.ID)
					st.LatestEventID = r.ID
					fmt.Printf("fail to unmarshal event:%v\n", r.ID)
					return
				}
				h := handlerFactory(event.Type(t))
				e.SetID(r.ID)
				err = h.Handle(e)
				if err != nil {
					fmt.Printf("handle event error eventType:%v err:%v\n", t, err)
				}

			}()
		}
	}
}

func initialState(snapshotSrv snapshot.Snapshot) *state.State {
	st := state.State{
		Users: map[int64]*user.User{},
		Items: map[string]*warehouse.Item{},
	}
	err := snapshotSrv.Restore(&st)
	if err != nil { // inital state for demo purpose
		fmt.Println(err)
		return &st
	}
	fmt.Println("state restored ")
	for _, u := range st.Users {
		fmt.Printf("userID:%v balance:%v \n", u.UseID, u.Balance)
	}
	for _, item := range st.Items {
		fmt.Printf("itemID:%v remain:%v price:%v \n", item.ID, item.Remain, item.Price)
	}
	return &st
}
