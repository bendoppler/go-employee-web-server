package models

// UserCallCount represents a user's call count
type UserCallCount struct {
	Username  string `json:"username"`
	CallCount int64  `json:"call_count"`
}
