package handlers

import "fmt"

func UsersList() {
	contacts, err := store.GetAll()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println(" No contacts to display.")
		return
	}

	fmt.Println("\n--- Contact List ---")
	for _, contact := range contacts {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", contact.ID, contact.Name, contact.Email)
	}

}
