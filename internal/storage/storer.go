package storage

import "fmt"

// Contact est notre structure de données centrale.
type User struct {
	ID    string
	Name  string
	Email string
}

// Storer est notre CONTRAT de stockage.
// Il définit un ensemble de comportements (méthodes) que tout type
// de stockage doit respecter. On ne se soucie pas de COMMENT c'est fait
// (en mémoire, fichier, BDD), seulement de CE QUI peut être fait.
type Storer interface {
	Add(contact *User) error
	GetAll() ([]*User, error)
	GetByID(id string) (*User, error) // Méthode pratique pour la suite
	Update(id string, newName string, newEmail string) error
	Delete(id string) error
}

var ErrContactNotFound = func(id string) error {
	return fmt.Errorf("User avec l'ID %d non trouvé", id)
}
