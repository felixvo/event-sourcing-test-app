package user

type User struct {
	UserID  int64 `json:"user_id"`
	Balance uint  `json:"balance"`
}
