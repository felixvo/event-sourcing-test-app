package user

type User struct {
	UseID   int64
	Balance int64
}

// UserDAO --
type Repository interface {
	Get(userID int64) (*User, error)
	MultiGet(userIDs []int64) ([]*User, error)
	MultiSetBalance(userIDs []int64, balance []int64) error
}
