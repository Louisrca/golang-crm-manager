package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Louisrca/golang-crm-manager.git/utils/colorText"
	"github.com/google/uuid"
)

type User struct {
	Name  string
	Email string
}

var users = map[string]User{
	"3de3d7d4-4cc8-49c9-8b1d-aeddd035aef2": {
		Name:  "louis",
		Email: "rocca@mail.com",
	},
}

func AddUser(nameFlag *string, emailFlag *string) (string, *User) {

	randID := uuid.New().String()

	reader := bufio.NewReader(os.Stdin)

	var name string
	if nameFlag != nil && *nameFlag != "" {
		name = *nameFlag
	} else {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Print("Name please ? ")
		input, _ := reader.ReadString('\n')
		name = strings.TrimSpace(input)
	}

	var email string
	if emailFlag != nil && *emailFlag != "" {
		email = *emailFlag
	} else {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Print("Email please ? ")
		input, _ := reader.ReadString('\n')
		email = strings.TrimSpace(input)
	}

	user := User{
		Name:  name,
		Email: email,
	}

	users[randID] = user
	return randID, &user
}

func UpdateUser() *User {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println("Who's id ?")
	bufID, _ := reader.ReadString('\n')
	bufID = strings.TrimSpace(bufID)
	u, exists := users[bufID]

	if !exists {
		fmt.Println("User not found!")
		return nil
	}

	fmt.Print("New Name (leave empty to keep current)? ")
	bufName, _ := reader.ReadString('\n')
	bufName = strings.TrimSpace(bufName)
	if bufName != "" {
		u.Name = bufName
	}

	fmt.Print("New Email (leave empty to keep current)? ")
	bufEmail, _ := reader.ReadString('\n')
	bufEmail = strings.TrimSpace(bufEmail)
	if bufEmail != "" {
		u.Email = bufEmail
	}

	users[bufID] = u
	fmt.Println("User updated successfully!")
	return &u

}

func DeleteUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Print("Who's ID ? ")
	bufID, _ := reader.ReadString('\n')
	bufID = strings.TrimSpace(bufID)

	if _, exists := users[bufID]; exists {
		delete(users, bufID)
		fmt.Println("User deleted!")
	} else {
		fmt.Println("User not found!")
	}
}

func GetUsers() {
	colorText.CyanText("=== Users List ===")
	for id, u := range users {
		fmt.Printf("ID: %s | Name: %s | Email: %s\n", id, u.Name, u.Email)
	}
	colorText.CyanText("==================")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
