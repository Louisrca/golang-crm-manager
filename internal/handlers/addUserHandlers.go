package handlers

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Louisrca/golang-crm-manager/internal/storage"
	rl "github.com/Louisrca/golang-crm-manager/utils/readline"
)

func AddUser(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)

	var (
		name  string
		email string
		err   error
	)

	for {
		fmt.Print("Enter contact name: ")
		name, err = rl.ReadLine(reader)
		if err != nil {
			fmt.Printf("Read error: %v. Please try again.\n", err)
			continue
		}
		if name != "" {
			break
		}
		fmt.Println("Name cannot be empty. Please enter a value.")
	}

	for {
		fmt.Print("Enter contact email: ")
		email, err = rl.ReadLine(reader)
		if err != nil {
			fmt.Printf("Read error: %v. Please try again.\n", err)
			continue
		}
		if email != "" {
			break
		}
		fmt.Println("Email cannot be empty. Please enter a value.")
	}
	newUser := &storage.User{Name: name, Email: email}
	err = store.Add(newUser)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf(" Contact '%s' added with ID %s.\n", newUser.Name, newUser.ID)
}
