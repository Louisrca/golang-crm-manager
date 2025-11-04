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

func (u *User) Display() {
	fmt.Printf("User → Name: %s | Email: %s\n", u.Name, u.Email)
}

func NewUser(name, email string) (*User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}

	if len(name) < 3 {
		return nil, fmt.Errorf("name must contain at least 3 characters")
	}

	return &User{Name: name, Email: email}, nil
}

var users = map[string]*User{
	"3de3d7d4-4cc8-49c9-8b1d-aeddd035aef2": {
		Name:  "louis",
		Email: "rocca@mail.com",
	},
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

func AddUser() (*User, error) {
	name, err := prompt("\nName: ")
	if err != nil {
		return nil, fmt.Errorf("failed to read name: %w", err)
	}

	email, err := prompt("Email: ")
	if err != nil {
		return nil, fmt.Errorf("failed to read email: %w", err)
	}

	user, err := NewUser(name, email)
	if err != nil {
		return nil, fmt.Errorf("invalid user data: %w", err)
	}

	id := uuid.New().String()
	users[id] = user
	colorText.GreenText("✅ User added successfully!")

	return user, nil
}

func AddUserWithFlag(nameFlag, emailFlag *string) (*User, error) {
	if nameFlag == nil || emailFlag == nil {
		return nil, fmt.Errorf("invalid flags: name and email flags are required")
	}

	user, err := NewUser(*nameFlag, *emailFlag)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	id := uuid.New().String()
	users[id] = user
	colorText.GreenText("✅ User added via flags")
	return user, nil
}

func UpdateUser() (*User, error) {
	id, err := prompt("\nUser ID to update: ")
	if err != nil {
		return nil, fmt.Errorf("failed to read user ID: %w", err)
	}

	user, exists := users[id]
	if !exists {
		return nil, fmt.Errorf("user not found with ID: %s", id)
	}

	fmt.Printf("Current: ")
	user.Display()

	confirm, err := prompt("Update this user? (yes/no): ")
	if err != nil {
		return nil, fmt.Errorf("failed to read confirmation: %w", err)
	}

	if confirm != "yes" {
		colorText.YellowText("❌ Update cancelled")
		return nil, nil
	}

	newName, err := prompt("New name (leave empty to keep): ")
	if err != nil {
		return nil, fmt.Errorf("failed to read new name: %w", err)
	}

	if newName != "" {
		newName = strings.TrimSpace(newName)
		if len(newName) < 3 {
			return nil, fmt.Errorf("name must contain at least 3 characters")
		}
		user.Name = newName
	}

	newEmail, err := prompt("New email (leave empty to keep): ")
	if err != nil {
		return nil, fmt.Errorf("failed to read new email: %w", err)
	}

	if newEmail != "" {
		newEmail = strings.TrimSpace(newEmail)
		if newEmail == "" {
			return nil, fmt.Errorf("email cannot be empty")
		}
		user.Email = newEmail
	}

	colorText.GreenText("✅ User updated!")
	return user, nil
}

func DeleteUser() error {
	id, err := prompt("\nUser ID to delete: ")
	if err != nil {
		return fmt.Errorf("failed to read user ID: %w", err)
	}

	if _, exists := users[id]; exists {
		delete(users, id)
		colorText.BlueText("🗑 User deleted")
		return nil
	}

	return fmt.Errorf("user not found with ID: %s", id)
}

func GetUsers() {
	colorText.CyanText("\n=== Users List ===")
	if len(users) == 0 {
		colorText.YellowText("No users found")
	} else {
		for id, u := range users {
			fmt.Printf("ID: %s | Name: %s | Email: %s\n", id, u.Name, u.Email)
		}
	}
	colorText.CyanText("==================\n")
}
