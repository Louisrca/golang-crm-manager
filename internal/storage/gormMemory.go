package storage

import (
	"fmt"

	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore(db *gorm.DB) *GORMStore {
	db.AutoMigrate(&GORMUser{})
	return &GORMStore{db: db}
}

func (gs *GORMStore) Add(user *User) error {
	gormUser := GORMUser{Name: user.Name, Email: user.Email}
	return gs.db.Create(&gormUser).Error
}

func (gs *GORMStore) GetAll() ([]*User, error) {
	var gormUsers []GORMUser
	if err := gs.db.Find(&gormUsers).Error; err != nil {
		return nil, err
	}

	users := make([]*User, len(gormUsers))
	for i, u := range gormUsers {
		users[i] = &User{
			ID:    fmt.Sprint(u.ID),
			Name:  u.Name,
			Email: u.Email,
		}
	}
	return users, nil
}

func (gs *GORMStore) GetByID(id string) (*User, error) {
	var gu GORMUser
	if err := gs.db.First(&gu, id).Error; err != nil {
		return nil, err
	}
	return &User{
		ID:    fmt.Sprint(gu.ID),
		Name:  gu.Name,
		Email: gu.Email,
	}, nil
}

func (gs *GORMStore) Update(id string, newName string, newEmail string) error {
	return gs.db.Model(&GORMUser{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":  newName,
			"email": newEmail,
		}).Error
}

func (gs *GORMStore) Delete(id string) error {
	return gs.db.Delete(&GORMUser{}, id).Error
}
