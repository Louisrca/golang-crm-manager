package storage

import "fmt"

// MemoryStore est une implémentation CONCRÈTE de l'interface Storer.
// Elle utilise une map en mémoire pour stocker les données.
// Elle respecte le contrat Storer car elle possède toutes les méthodes demandées.
type MemoryStore struct {
	users  map[int]*User
	nextID int
}

// NewMemoryStore est un "constructeur" qui initialise proprement notre store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

func (ms *MemoryStore) Add(user *User) error {
	user.ID = ms.nextID
	ms.users[user.ID] = user
	ms.nextID++
	return nil
}

func (ms *MemoryStore) GetAll() ([]*User, error) {
	// On crée une slice pour éviter de retourner directement
	// une référence à la map interne.
	var allContacts []*User
	for _, c := range ms.users {
		allContacts = append(allContacts, c)
	}
	return allContacts, nil
}

func (ms *MemoryStore) GetByID(id int) (*User, error) {
	user, ok := ms.users[id]
	if !ok {
		return nil, fmt.Errorf("user avec l'ID %d non trouvé", id)
	}
	return user, nil
}

func (ms *MemoryStore) Update(id int, newName, newEmail string) error {
	user, err := ms.GetByID(id)
	if err != nil {
		return err // Retourne l'erreur "non trouvé"
	}

	if newName != "" {
		user.Name = newName
	}
	if newEmail != "" {
		user.Email = newEmail
	}
	return nil
}

func (ms *MemoryStore) Delete(id int) error {
	if _, ok := ms.users[id]; !ok {
		return fmt.Errorf("user avec l'ID %d non trouvé", id)
	}
	delete(ms.users, id)
	return nil
}
