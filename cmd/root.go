package cmd

import (
	"fmt"

	"github.com/Louisrca/golang-crm-manager/internal/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Store storage.Storer

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "CRM CLI (YAML config + flags combinés)",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("✅ Config loaded from", viper.ConfigFileUsed())
		} else {
			fmt.Println("⚠️ No config.yaml found, using JSON storage by default")
		}

		storeType := viper.GetString("store")
		jsonPath := viper.GetString("json_path")
		dbPath := viper.GetString("db_path")

		switch storeType {
		case "JSON":
			s, err := storage.NewJSONStore(jsonPath)
			if err != nil {
				return err
			}
			Store = s
			fmt.Println("📦 Using JSON storage")

		case "GORM":
			db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
			if err != nil {
				return err
			}
			Store = storage.NewGORMStore(db)
			fmt.Println("📦 Using GORM storage")
		case "LOCAL":
			Store = storage.NewMemoryUserStore()
			fmt.Println("📦 Using Local Memory storage")

		default:
			return fmt.Errorf("invalid store type: %s", storeType)
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌", err)
	}
}

func init() {
	// Définition des flags Cobra
	rootCmd.PersistentFlags().String("store", "", "Type de stockage (JSON ou GORM)")
	rootCmd.PersistentFlags().String("json_path", "", "Chemin du fichier JSON")
	rootCmd.PersistentFlags().String("db_path", "", "Chemin de la base SQLite")

	// Liaison flags → Viper
	viper.BindPFlag("store", rootCmd.PersistentFlags().Lookup("store"))
	viper.BindPFlag("json_path", rootCmd.PersistentFlags().Lookup("json_path"))
	viper.BindPFlag("db_path", rootCmd.PersistentFlags().Lookup("db_path"))

	// Valeurs par défaut
	viper.SetDefault("store", "JSON")
	viper.SetDefault("json_path", "users.json")
	viper.SetDefault("db_path", "users.db")

}
