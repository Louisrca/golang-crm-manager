package handlers

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Louisrca/golang-crm-manager/internal/storage"
	rl "github.com/Louisrca/golang-crm-manager/utils/readline"
)

func DeleteContact(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the ID of the contact to delete: ")
	id, _ := rl.ReadLine(reader)

	err := store.Delete(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Contact with ID %d has been deleted.\n", id)
}
