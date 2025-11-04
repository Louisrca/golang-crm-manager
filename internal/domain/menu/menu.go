package menu

import (
	"fmt"

	"github.com/Louisrca/golang-crm-manager.git/internal/domain/user"
	"github.com/Louisrca/golang-crm-manager.git/utils/colorText"
)

func MenuChoice() {
	fmt.Println("____________________________")
	fmt.Println("Welcome to CRM Manager")
	fmt.Println("____________________________")
	fmt.Println("What do you want to do ?")
	fmt.Println("\n")
	fmt.Println("1. Create a user")
	fmt.Println("2. Update a user")
	fmt.Println("3. Remove a user")
	fmt.Println("4. Users list")
	fmt.Println("5. Exit")
	fmt.Println("____________________________")
	fmt.Println()
}

func Menu() {
	for {

		MenuChoice()

		var choice int

		_, err := fmt.Scan(&choice)

		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("❌ Invalid input. Please enter a number.\n\n")
			continue
		}

		fmt.Println()

		switch choice {
		case 1:
			user.AddUser()
		case 2:
			user.UpdateUser()
		case 3:
			user.DeleteUser()
		case 4:
			user.GetUsers()
		case 5:
			colorText.CyanText("Bye bye")
			return
		default:
			colorText.RedText("Please Make a choice boy")
		}
	}
}
