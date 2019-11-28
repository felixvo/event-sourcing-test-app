package warehouse

type Item struct {
	ID     string
	Remain int64
}

// Repository for item
type Repository interface {
	GetItems(itemIDs []string) ([]*Item, error)
	UpdateRemains(items []*Item) error
}
