package cmd

import (
	"fmt"
	coreConfig "github.com/nonamecat19/go-orm/core/lib/config"
	"github.com/spf13/cobra"
	"time"
)

var migrationCreateCmd = &cobra.Command{
	Use:   "migration:create",
	Short: "Create migration",
	Long:  `Create a new migration file with the given name`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		config, err := coreConfig.ParseYamlConfig(configPath)
		if err != nil {
			panic(err)
		}
		migrationConfig := (*config).Migrations

		name, _ := cmd.Flags().GetString("name")
		if migrationConfig.AddTimestamp {
			timestamp := time.Now().Format("20060102150405")
			name = fmt.Sprintf("%s_%s", timestamp, name)
		}

		println("Creating migration", name)

	},
}

func init() {
	rootCmd.AddCommand(migrationCreateCmd)

	migrationCreateCmd.Flags().StringP("config", "c", "", "Path to the config file")
	migrationCreateCmd.Flags().StringP("name", "n", "", "Path to the config file")
}
