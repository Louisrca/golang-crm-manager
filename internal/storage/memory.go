package storage

import (
	"fmt"

	"github.com/google/uuid"
)

type MemoryUserStore struct {
	users map[string]*User
}

func NewMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{
		users: make(map[string]*User),
	}
}

func (m *MemoryUserStore) Add(user *User) error {
	id := uuid.New().String()
	user.ID = id
	m.users[id] = user
	return nil
}

func (m *MemoryUserStore) GetAll() ([]*User, error) {
	res := make([]*User, 0, len(m.users))
	for _, u := range m.users {
		res = append(res, u)
	}
	return res, nil
}

func (m *MemoryUserStore) GetByID(id string) (*User, error) {
	user := m.users[id]
	if user == nil {
		return nil, fmt.Errorf("user avec l'ID %s non trouvé", id)
	}
	return user, nil
}

func (m *MemoryUserStore) Update(id string, name string, email string) error {
	user := m.users[id]
	if user == nil {
		return fmt.Errorf("user avec l'ID %s non trouvé", id)
	}
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	return nil
}

func (m *MemoryUserStore) Delete(id string) error {
	if _, ok := m.users[id]; !ok {
		return fmt.Errorf("user avec l'ID %s non trouvé", id)
	}
	delete(m.users, id)
	return nil
}
