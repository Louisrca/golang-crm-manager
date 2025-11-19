package cmd

import (
	"github.com/Louisrca/golang-crm-manager/internal/handlers"
	"github.com/spf13/cobra"
)

var deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "This command allows you to delete an user with his ID",
	Long:  `./golang-crm-manager delete-user --store <storer>`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.DeleteContact(Store)
	}}

func init() {
	rootCmd.AddCommand(deleteUserCmd)
}
