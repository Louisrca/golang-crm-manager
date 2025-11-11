package cmd

import (
	"fmt"
	"os"

	"github.com/Louisrca/golang-crm-manager/internal/storage"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	storeType string
	Store     storage.Storer
)

var rootCmd = &cobra.Command{
	Use:   "crm-manager",
	Short: "CRM avec plusieurs backends (JSON/GORM)",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		switch storeType {
		case "JSON":
			var err error
			Store, err = storage.NewJSONStore("users.json")
			if err != nil {
				return fmt.Errorf("failed to init JSON store: %w", err)
			}
			fmt.Println("📦 Using JSON storage")
		case "GORM":
			db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
			if err != nil {
				return fmt.Errorf("cannot init GORM: %w", err)
			}
			Store = storage.NewGORMStore(db)
			fmt.Println("📦 Using GORM storage")
		default:
			return fmt.Errorf("invalid store type: %s (use JSON or GORM)", storeType)
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&storeType, "store", "JSON", "Type de stockage: JSON ou GORM")
}
