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
	fmt.Println("4. Update user")
	fmt.Println("5. Exit")
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
			handleAddUser(store)
		case 2:
			handleGetUsers(store)
		case 3:
			handleGetUsersById(store)
		case 4:
			handleUpdateUser(store)
		case 5:
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

func handleAddUser(storer storage.Storer) {

	name, _ := prompt("\nName: ")
	email, _ := prompt("\nEmail: ")

	newUser := storage.NewUser(name, email)

	err := store.AddUser(newUser)

	if err != nil {
		fmt.Println("ERRRRROOOOORR")
	}

	fmt.Println("User added successfully: ", newUser.ID)
}

func handleGetUsers(storer storage.Storer) {
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

func handleGetUsersById(storer storage.Storer) {
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

func handleUpdateUser(storer storage.Storer) {
	userID, err := prompt("\nGive me the UserID: ")
	if err != nil {
		colorText.RedText(" Error reading ID")
		return
	}

	userID = strings.TrimSpace(userID)
	if userID == "" {
		colorText.RedText(" User ID cannot be empty\n")
		return
	}

	currentUser, err := store.GetUserByID(userID)
	if err != nil {
		colorText.RedText(fmt.Sprintf(" Error: %s\n", err))
		return
	}

	fmt.Printf("Current user → ID: %s | Name: %s | Email: %s\n",
		currentUser.ID, currentUser.Name, currentUser.Email)

	newName, err := prompt("\nNew Name (leave empty to keep): ")
	if err != nil {
		colorText.RedText(fmt.Sprintf(" Error reading Name: %s\n", err))
		return
	}

	newEmail, err := prompt("New Email (leave empty to keep): ")
	if err != nil {
		colorText.RedText(fmt.Sprintf(" Error reading Email: %s\n", err))
		return
	}

	newName = strings.TrimSpace(newName)
	newEmail = strings.TrimSpace(newEmail)

	if newName == "" {
		newName = currentUser.Name
	} else if len(newName) < 3 {
		colorText.RedText(" Name must contain at least 3 characters\n")
		return
	}

	if newEmail == "" {
		newEmail = currentUser.Email
	} else if !strings.Contains(newEmail, "@") {
		colorText.RedText(" Email must contain @\n")
		return
	}

	err = store.UpdateUser(userID, newName, newEmail)
	if err != nil {
		colorText.RedText(fmt.Sprintf(" Error: %s\n", err))
		return
	}

	updatedUser, err := store.GetUserByID(userID)
	if err != nil {
		colorText.RedText(fmt.Sprintf(" Error: %s\n", err))
		return
	}

	colorText.GreenText("User Updated!\n")
	colorText.CyanText("\n=== Updated User ===\n")
	fmt.Printf("ID: %s | Name: %s | Email: %s\n",
		updatedUser.ID, updatedUser.Name, updatedUser.Email)
	colorText.CyanText("====================\n\n")
}
