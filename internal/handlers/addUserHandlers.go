package handlers

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Louisrca/golang-crm-manager/internal/storage"
	readl "github.com/Louisrca/golang-crm-manager/utils/readline"
)

var store = storage.NewMemoryStore()

func AddUser() {
	reader := bufio.NewReader(os.Stdin)

	var (
		name  string
		email string
		err   error
	)

	for {
		fmt.Print("Enter contact name: ")
		name, err = readl.ReadLine(reader)
		if err != nil {
			fmt.Printf("Read error: %v. Please try again.\n", err)
			continue
		}
		if name != "" {
			break
		}
		fmt.Println("Name cannot be empty. Please enter a value.")
	}

	// Boucle pour l'email (obligatoire)
	for {
		fmt.Print("Enter contact email: ")
		email, err = readl.ReadLine(reader)
		if err != nil {
			fmt.Printf("Read error: %v. Please try again.\n", err)
			continue
		}
		if email != "" {
			break
		}
		fmt.Println("Email cannot be empty. Please enter a value.")
	}
	contact := &storage.User{Name: name, Email: email}
	err = store.Add(contact)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf(" Contact '%s' added with ID %d.\n", contact.Name, contact.ID)
}
