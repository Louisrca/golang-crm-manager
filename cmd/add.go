package cmd

import (
	"github.com/Louisrca/golang-crm-manager/internal/handlers"
	"github.com/spf13/cobra"
)

var addUserCmd = &cobra.Command{
	Use:   "add",
	Short: "Hello add",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.AddUser()
	}}

func init() {
	rootCmd.AddCommand(addUserCmd)
}
