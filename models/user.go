package models

// User ..
type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UserRepository ..
type UserRepository interface {
	All() (*[]User, error)
}
