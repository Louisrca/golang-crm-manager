package main

import (
	"flag"

	"github.com/Louisrca/golang-crm-manager.git/internal/domain/menu"
	"github.com/Louisrca/golang-crm-manager.git/internal/domain/user"
)

func main() {

	nameFlag := flag.String("name", "", "Nom du contact")
	emailFlag := flag.String("email", "", "email du contact")
	flag.Parse()

	if *nameFlag != "" && *emailFlag != "" {
		user.AddUserWithFlag(nameFlag, emailFlag)
	}
	menu.Menu()

}
