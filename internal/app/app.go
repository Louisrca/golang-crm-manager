package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Louisrca/golang-crm-manager.git/internal/storage"
	"github.com/Louisrca/golang-crm-manager.git/utils/colorText"
)

var store = storage.NewMemoryUser()

func MenuChoice() {
	fmt.Println("____________________________")
	fmt.Println("Welcome to CRM Manager")
	fmt.Println("____________________________")
	fmt.Println("What do you want to do ?")
	fmt.Println("\n")
	fmt.Println("1. Create a user")
	fmt.Println("2. Users list")
	fmt.Println("3. User by ID")
	fmt.Println("4. Exit")
	fmt.Println("____________________________")
	fmt.Println()
}

func Run() {

	for {

		MenuChoice()

		var choice int

		_, err := fmt.Scan(&choice)

		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Invalid input. Please enter a number.\n\n")
			continue
		}

		fmt.Println()

		switch choice {
		case 1:
			handleAddUser()
		case 2:
			handleGetUsers()
		case 3:
			handleGetUsersById()
		case 4:
			colorText.CyanText("Bye bye")
			return
		default:
			colorText.RedText("Please Make a choice boy")
		}
	}
}

func prompt(label string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	in, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	return strings.TrimSpace(in), nil
}

func handleAddUser() {

	name, err := prompt("\nName: ")
	if err != nil {
		colorText.RedText(fmt.Sprintf("Error reading name: %s\n", err))
		return
	}

	email, err := prompt("\nEmail: ")

	if err != nil {
		colorText.RedText(fmt.Sprintf("Error reading email: %s\n", err))
		return
	}
	newUser, err := storage.NewUser(name, email)

	id, err := store.AddUser(newUser)

	colorText.GreenText(fmt.Sprintf("User added successfully with ID: %s\n", id))
}

func handleGetUsers() {
	users, err := store.GetUsers()
	if err != nil {
		colorText.RedText(fmt.Sprintf("Error: %s\n", err))
		return
	}

	for _, u := range users {
		fmt.Printf("ID: %s | Name: %s | Email: %s\n", u.ID, u.Name, u.Email)
	}
	colorText.CyanText("==================")
}

func handleGetUsersById() {
	userID, err := prompt("\nGive me the UserID: ")
	if err != nil {
		colorText.RedText(fmt.Sprintf("Error reading ID: %s\n", err))
		return
	}

	userID = strings.TrimSpace(userID)
	if userID == "" {
		colorText.RedText("User ID cannot be empty\n")
		return
	}

	user, err := store.GetUserByID(userID)
	if err != nil {
		colorText.RedText(fmt.Sprintf("Error: %s\n", err))
		return
	}

	colorText.CyanText("\n=== User Found ===\n")
	fmt.Printf("ID: %s | Name: %s | Email: %s\n", user.ID, user.Name, user.Email)
	colorText.CyanText("==================\n\n")
}
