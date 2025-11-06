package handlers

import (
	"bufio"
	"fmt"
	"os"

	readl "github.com/Louisrca/golang-crm-manager/utils/readline"
)

func DeleteContact() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the ID of the contact to delete: ")
	id, _ := readl.ReadLine(reader)

	err := store.Delete(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Contact with ID %d has been deleted.\n", id)
}
