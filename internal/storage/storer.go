package storage

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	ID    string
	Name  string
	Email string
}

func NewUser(name, email string) (*User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	id := uuid.New().String()

	if name == "" {
		fmt.Errorf("name cannot be empty")
		return nil, nil
	}

	if email == "" {
		fmt.Errorf("email cannot be empty")
		return nil, nil
	}

	if len(name) < 3 {
		fmt.Errorf("name must contain at least 3 characters")
		return nil, nil
	}

	return &User{ID: id, Name: name, Email: email}, nil
}

func (u User) Display() {
	fmt.Println()
}

type UserStorer interface {
	AddUser(user *User) error
	GetUsers() ([]*User, error)
	GetUserByID(id string) (*User, error)
	UpdateUser(id string, name string, email string) error
	DeleteUser(id string) error
}

var ErrorUserNotFound = func(id string) error {
	return fmt.Errorf("User with id %s: ", id, " not found")
}
