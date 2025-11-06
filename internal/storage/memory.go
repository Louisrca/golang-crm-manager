package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type JSONStore struct {
	filePath string
	users    map[string]*User
}

func NewJSONStore(filePath string) (*JSONStore, error) {
	store := &JSONStore{
		filePath: filePath,
		users:    make(map[string]*User),
	}

	if err := store.load(); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	return store, nil
}

func (js *JSONStore) Add(user *User) error {
	id := uuid.New().String()
	user.ID = id
	js.users[user.ID] = user
	return js.save()
}

func (js *JSONStore) GetAll() ([]*User, error) {
	all := make([]*User, 0, len(js.users))
	for _, u := range js.users {
		all = append(all, u)
	}
	return all, nil
}

func (js *JSONStore) GetByID(id string) (*User, error) {
	user, ok := js.users[id]
	if !ok {
		return nil, fmt.Errorf("user avec l'ID %d non trouvé", id)
	}
	return user, nil
}

func (js *JSONStore) Update(id string, newName, newEmail string) error {
	user, err := js.GetByID(id)
	if err != nil {
		return err
	}

	if newName != "" {
		user.Name = newName
	}
	if newEmail != "" {
		user.Email = newEmail
	}
	return js.save()
}

func (js *JSONStore) Delete(id string) error {
	if _, ok := js.users[id]; !ok {
		return fmt.Errorf("user avec l'ID %d non trouvé", id)
	}
	delete(js.users, id)
	return js.save()
}

func (js *JSONStore) save() error {
	file, err := os.Create(js.filePath)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier : %w", err)
	}
	defer file.Close()

	usersSlice := make([]*User, 0, len(js.users))
	for _, u := range js.users {
		usersSlice = append(usersSlice, u)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(usersSlice)
}

func (js *JSONStore) load() error {
	file, err := os.Open(js.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var users []*User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		return fmt.Errorf("erreur de décodage JSON : %w", err)
	}

	js.users = make(map[string]*User)
	for _, u := range users {
		js.users[u.ID] = u
	}

	return nil
}
