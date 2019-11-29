package main

import (
	"fmt"
	"github.com/felixvo/lmax/cmd/consumer/handler"
	"github.com/felixvo/lmax/cmd/consumer/state"
	"github.com/felixvo/lmax/pkg/event"
	"github.com/felixvo/lmax/pkg/snapshot"
	"github.com/felixvo/lmax/pkg/user"
	"github.com/felixvo/lmax/pkg/warehouse"
	"github.com/go-redis/redis/v7"
	"strconv"
	"strings"
	"time"
)

const (
	OrderStream = "orders"
)

func main() {
	client, err := newRedisClient()
	if err != nil {
		panic(err)
	}
	snapshotSrv := snapshot.NewRedisSnapshot(client)
	st := initialState(snapshotSrv)
	//
	go exeSnapshot(st, snapshotSrv)

	// start fetch events
	events := eventFetcher(client, st)

	// start consume events
	consumeEvents(events, handler.HandlerFactory(st))

	quit := make(chan bool)
	<-quit
}
func exeSnapshot(st *state.State, snapshotSrv snapshot.Snapshot) {
	ticker := time.Tick(time.Second * 30)
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

// start fetch new event starting from st.LatestEventID
func eventFetcher(client *redis.Client, st *state.State) chan event.Event {
	c := make(chan event.Event, 100)
	start := "-"
	if len(st.LatestEventID) > 0 {
		splitted := strings.Split(st.LatestEventID, "-")
		counter, _ := strconv.Atoi(splitted[1])
		start = fmt.Sprintf("%s-%v", splitted[0], counter+1)
	}
	go func() {
		for {
			func() {
				defer func() { // increase start by once after processed all the new messages
					splitted := strings.Split(start, "-")
					counter, _ := strconv.Atoi(splitted[1])
					start = fmt.Sprintf("%s-%v", splitted[0], counter+1)
				}()
				rr, err := client.XRange(OrderStream, start, "+").Result()
				if err != nil {
					panic(err)
				}

				for _, r := range rr {
					start = r.ID
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
					e.SetID(r.ID)
					c <- e
				}
			}()
		}
	}()
	return c
}

func consumeEvents(events chan event.Event, handlerFactory func(t event.Type) handler.Handler) {
	for {
		select {
		case e := <-events:
			h := handlerFactory(e.GetType())
			err := h.Handle(e)
			if err != nil {
				fmt.Printf("handle event error eventType:%v err:%v\n", e.GetType(), err)
			}
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
