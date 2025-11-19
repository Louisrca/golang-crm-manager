package cmd

import (
	"github.com/Louisrca/golang-crm-manager/internal/handlers"
	"github.com/spf13/cobra"
)

var updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "This command allows you to update an user with his ID",
	Long:  `./golang-crm-manager update-user --store <storer>`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.UpdateUser(Store)
	}}

func init() {
	rootCmd.AddCommand(updateUserCmd)
}
