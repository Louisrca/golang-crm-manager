package storage

import "fmt"

type User struct {
	ID    string
	Name  string
	Email string
}

type UserStorer interface {
	AddUser(user *User) error
	GetUsers() ([]*User, error)
	GetUserByID(id string) (*User, error)
	UpdateUser(id string, name string, email string) error
	DeleteUser(id string) error
}

var ErrorUserNotFound = func(id int) error {
	return fmt.Errorf("User with id %s: ", id, " not found")
}
