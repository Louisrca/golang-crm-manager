package handlers

import (
	"fmt"

	readi "github.com/Louisrca/golang-crm-manager/utils/readInteger"
)

func DeleteContact() {

	fmt.Print("Enter the ID of the contact to delete: ")
	id := readi.ReadInteger()
	if id == -1 {
		return
	}

	err := store.Delete(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Contact with ID %d has been deleted.\n", id)
}
