package readi

import (
	"bufio"
	"os"
	"strconv"

	rl "github.com/Louisrca/golang-crm-manager/utils/readline"
)

func ReadInteger() int {
	reader := bufio.NewReader(os.Stdin)

	id, err := rl.ReadLine(reader)
	if err != nil {
		return -1 // Renvoie -1 pour un choix invalide
	}
	nbrId, err := strconv.Atoi(id)
	if err != nil {
		return -1
	}

	return nbrId
}
