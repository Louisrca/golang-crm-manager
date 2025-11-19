package handlers

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Louisrca/golang-crm-manager/internal/storage"
	rl "github.com/Louisrca/golang-crm-manager/utils/readline"
)

func UpdateUser(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the ID of the contact to update: ")
	id, _ := rl.ReadLine(reader)

	// On vérifie que le contact existe avant de demander les nouvelles infos
	existingContact, err := store.GetByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Updating '%s'. Leave blank to keep current value.\n", existingContact.Name)

	fmt.Printf("New name (%s): ", existingContact.Name)
	newName, err := rl.ReadLine(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("New email (%s): ", existingContact.Email)
	newEmail, err := rl.ReadLine(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = store.Update(id, newName, newEmail)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Contact updated successfully.")
}
