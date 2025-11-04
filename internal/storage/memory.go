package storage

import (
	"fmt"

	"github.com/google/uuid"
)

type MemoryUserStore struct {
	users map[string]*User
}

func NewMemoryUser() *MemoryUserStore {
	return &MemoryUserStore{users: make(map[string]*User)}
}

func (m *MemoryUserStore) AddUser(user *User) (string, error) {
	if user == nil {
		return "", fmt.Errorf("user cannot be nil")
	}

	id := uuid.New().String()

	user.ID = id

	m.users[id] = user
	return id, nil
}

func (m *MemoryUserStore) GetUsers() ([]*User, error) {
	users := make([]*User, 0, len(m.users))
	for _, u := range m.users {
		users = append(users, u)
	}
	return users, nil
}

func (m *MemoryUserStore) GetUserByID(id string) (*User, error) {

	user, exists := m.users[id]

	if !exists {
		return nil, fmt.Errorf("user not found with ID: %s", id)
	}

	return user, nil
}

func (m *MemoryUserStore) UpdateUser(id string, name string, email string) error {
	user, exists := m.users[id]

	if !exists {
		return fmt.Errorf("user not found with ID: %s", id)
	}

	user.Name = name
	user.Email = email

	return nil
}

func (m *MemoryUserStore) DeleteUser(id string) error {
	if _, exists := m.users[id]; !exists {
		return fmt.Errorf("user not found with ID: %s", id)
	}

	delete(m.users, id)

	return nil
}
