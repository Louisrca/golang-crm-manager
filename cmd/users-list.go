package cmd

import (
	"github.com/Louisrca/golang-crm-manager/internal/handlers"
	"github.com/spf13/cobra"
)

var usersListCmd = &cobra.Command{
	Use:   "users-list",
	Short: "Display a list of users",
	Long:  `./golang-crm-manager users-list --store <storer>`,
	Run: func(cmd *cobra.Command, args []string) {

		handlers.UsersList(Store)

	}}

func init() {
	rootCmd.AddCommand(usersListCmd)
}
