package state

import (
	"github.com/felixvo/lmax/cmd/pkg/user"
	"github.com/felixvo/lmax/cmd/pkg/warehouse"
)

type State struct {
	users map[int64]*user.User
	items map[string]*warehouse.Item
}

func (s *State) GetUserByID(id int64) *user.User {
	return s.users[id]
}

func (s *State) GetItem(id string) *warehouse.Item {
	return s.items[id]
}

func (s *State) GetItems(ids []string) ([]*warehouse.Item) {
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
