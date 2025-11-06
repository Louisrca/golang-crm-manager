package storage

import (
	"github.com/google/uuid"
)

type MemoryUserStore struct {
	users map[string]*User
}

func NewMemoryUser() *MemoryUserStore {
	return &MemoryUserStore{users: make(map[string]*User)}
}

func (m *MemoryUserStore) AddUser(user *User) error {

	id := uuid.New().String()

	user.ID = id

	m.users[id] = user
	return nil
}

func (m *MemoryUserStore) GetUsers() ([]*User, error) {
	users := make([]*User, 0, len(m.users))
	for _, u := range m.users {
		users = append(users, u)
	}
	return users, nil
}

func (m *MemoryUserStore) GetUserByID(id string) (*User, error) {

	user := m.users[id]

	return user, nil
}

func (m *MemoryUserStore) UpdateUser(id string, name string, email string) error {
	user := m.users[id]
	user.Name = name
	user.Email = email

	return nil
}

func (m *MemoryUserStore) DeleteUser(id string) error {

	delete(m.users, id)

	return nil
}
