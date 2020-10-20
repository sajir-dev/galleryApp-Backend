package domain

// User ...
type User struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
