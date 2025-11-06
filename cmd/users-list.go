package cmd

import (
	 "github.com/Louisrca/golang-crm-manager/internal/handlers"
	"github.com/spf13/cobra"
)

var usersListCmd = &cobra.Command{
	Use:   "users-list",
	Short: "Hello add",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		handlers.UsersList()

	}}

func init() {
	rootCmd.AddCommand(usersListCmd)
}
