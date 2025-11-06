package cmd

import (
	"github.com/Louisrca/golang-crm-manager/internal/handlers"
	"github.com/spf13/cobra"
)

var deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Hello add",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.DeleteContact()
	}}

func init() {
	rootCmd.AddCommand(deleteUserCmd)
}
