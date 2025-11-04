package storer

import "fmt"

type User struct {
	ID    string
	Name  string
	Email string
}

type UserStorer interface {
	AddUser(user *User) error
	GetUsers() ([]*User, error)
	GetUserByID(id int) (*User, error)
	UpdateUser(id int, name string, email string) error
	DeleteUser(id int) error
}

var ErrorUserNotFound = func(id int) error {
	return fmt.Errorf("User with id %s: ", id, " not found")
}
