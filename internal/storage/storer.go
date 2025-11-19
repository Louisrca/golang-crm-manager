package storage

import (
	"fmt"

	"gorm.io/gorm"
)

// Type générique utilisé dans la logique métier
type User struct {
	ID    string
	Name  string
	Email string
}

// Modèle spécifique à GORM (peu évoluer à l'avenir)
type GORMUser struct {
	gorm.Model
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
	Email string `json:"email" gorm:"type:varchar(100);not null"`
}

// Modèle spécifique au JSON (peu évoluer à l'avenir)
type JSONUser struct {
	ID    string
	Name  string
	Email string
}

// Storer est notre CONTRAT de stockage.
// Il définit un ensemble de comportements (méthodes) que tout type
// de stockage doit respecter. On ne se soucie pas de COMMENT c'est fait
// (en mémoire, fichier, BDD), seulement de CE QUI peut être fait.
type Storer interface {
	Add(user *User) error
	GetAll() ([]*User, error)
	GetByID(id string) (*User, error) // Méthode pratique pour la suite
	Update(id string, newName string, newEmail string) error
	Delete(id string) error
}

var ErrContactNotFound = func(id string) error {
	return fmt.Errorf("User avec l'ID %d non trouvé", id)
}
