package cmd

import (
	"github.com/Louisrca/golang-crm-manager/internal/handlers"
	"github.com/spf13/cobra"
)

var updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Hello add",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.UpdateUser()
	}}

func init() {
	rootCmd.AddCommand(updateUserCmd)
}
