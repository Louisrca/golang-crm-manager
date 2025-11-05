package storage

import (
	"fmt"
	"strings"
)

type User struct {
	ID    string
	Name  string
	Email string
}

func NewUser(name, email string) *User {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		fmt.Errorf("name cannot be empty")
		return nil
	}

	if email == "" {
		fmt.Errorf("email cannot be empty")
		return nil
	}

	if len(name) < 3 {
		fmt.Errorf("name must contain at least 3 characters")
		return nil
	}

	return &User{Name: name, Email: email}
}

type Storer interface {
	AddUser(user *User) error
	GetUsers() ([]*User, error)
	GetUserByID(id string) (*User, error)
	UpdateUser(id string, name string, email string) error
	DeleteUser(id string) error
}

var ErrorUserNotFound = func(id string) error {
	return fmt.Errorf("User with id %s: ", id, " not found")
}
