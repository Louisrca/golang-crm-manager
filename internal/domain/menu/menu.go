package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Louisrca/golang-crm-manager.git/internal/domain/user"
	"github.com/Louisrca/golang-crm-manager.git/utils/colorText"
)

func MenuChoice() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Welcome to CRM Manager")
	fmt.Println("-------------------------------------")
	fmt.Println("\n")
	fmt.Println("What do you want to do ?")
	fmt.Println("\n")
	fmt.Println("1. Create a user")
	fmt.Println("2. Update a user")
	fmt.Println("3. Remove a user")
	fmt.Println("4. Users list")
	fmt.Println("-------------------------------------")
	fmt.Println("-------------------------------------")
	fmt.Println("-------------End of menu-------------")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func OptionSelector(opt string) {
	switch {
	case opt == "1":
		user.AddUser()
	case opt == "2":
		user.UpdateUser()
	case opt == "3":
		user.DeleteUser()
	case opt == "4":
		user.GetUsers()
	default:
		colorText.RedText("Please Make a choice boy")
	}
}

func Menu() {
	for {

		reader := bufio.NewReader(os.Stdin)
		MenuChoice()
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		OptionSelector(option)
	}
}
