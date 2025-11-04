package memory

import "github.com/Louisrca/golang-crm-manager.git/internal/storer"

type MemoryUserStore struct {
	users map[string]*storer.User
}

func NewMemoryUser() *MemoryUserStore {
	return &MemoryUserStore{users: make(map[string]*storer.User)}
}

func (m MemoryUserStore) Add(){

}
