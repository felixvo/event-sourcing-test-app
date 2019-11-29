package state

import (
	"github.com/felixvo/lmax/pkg/user"
	"github.com/felixvo/lmax/pkg/warehouse"
)

type State struct {
	LatestEventID string
	Users         map[int64]*user.User
	Items         map[string]*warehouse.Item
}

func (s *State) SetLatestEventID(id string) {
	s.LatestEventID = id
}
func (s *State) GetLatestEventID() string {
	return s.LatestEventID
}

// for inital state, state should calculated from events
func (s *State) SetUsers(users map[int64]*user.User) {
	s.Users = users
}

// for inital state, state should calculated from events
func (s *State) SetItems(items map[string]*warehouse.Item) {
	s.Items = items
}

func (s *State) GetUserByID(id int64) *user.User {
	u, exist := s.Users[id]
	if !exist {
		return &user.User{}
	}
	return u
}

func (s *State) GetItem(id string) *warehouse.Item {
	return s.Items[id]
}

func (s *State) GetItems(ids []string) []*warehouse.Item {
	rs := make([]*warehouse.Item, len(ids))
	for i, id := range ids {
		if item := s.GetItem(id); item != nil {
			rs[i] = s.GetItem(id)
		} else { // item not exist => should show error, this just for demo
			rs[i] = &warehouse.Item{
				ID:     id,
				Remain: 0,
			}
		}
	}
	return rs
}
