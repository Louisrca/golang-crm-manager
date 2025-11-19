package cmd

import (
	"github.com/Louisrca/golang-crm-manager/internal/handlers"
	"github.com/spf13/cobra"
)

var addUserCmd = &cobra.Command{
	Use:   "add",
	Short: "Allow you to add an user",
	Long:  `To use it :  ./golang-crm-manager add --store <storer>`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.AddUser(Store)
	}}

func init() {
	rootCmd.AddCommand(addUserCmd)
}
