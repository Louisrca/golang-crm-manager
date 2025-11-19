package handlers

import (
	"fmt"

	"github.com/Louisrca/golang-crm-manager/internal/storage"
)

func UsersList(store storage.Storer) {
	contacts, err := store.GetAll()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println(" No user to display.")
		return
	}

	fmt.Println("\n--- User List ---")
	for _, user := range contacts {
		fmt.Printf("ID: %s, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

}
