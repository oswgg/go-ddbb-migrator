package cmd

import (
	"fmt"
	"github.com/oswgg/migrator/internal/database/migrations"
	"github.com/spf13/cobra"
	"os"
)

var specificMigration string
var fromMigration string
var toMigration string
var env string

var migrateCmd = &cobra.Command{
	Use:   "migrate [up|down]",
	Short: "Run migrations",
	Long:  `Run migrations.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var upDownFlag = args[0]
		var upDownValue migrations.MigrationType

		var isSpecific = false
		var migrationName string
		var migratorValues *migrations.Migrator

		if specificMigration != "" {
			isSpecific = true
			migrationName = specificMigration
		}

		if upDownFlag == string(migrations.MigrationUp) {
			upDownValue = migrations.MigrationUp
		} else {
			upDownValue = migrations.MigrationDown
		}

		if isSpecific {
			migratorValues = &migrations.Migrator{
				Specific:          isSpecific,
				SpecificMigration: migrationName,
				MigrationType:     upDownValue,
				Env:               env,
			}
		} else {
			migratorValues = &migrations.Migrator{
				From:          fromMigration,
				To:            toMigration,
				MigrationType: upDownValue,
				Env:           env,
			}
		}

		migrator, err := migrations.NewMigrator(migratorValues)
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(0)
		}

		if upDownValue == migrations.MigrationUp {
			err := migrator.Up()
			if err != nil {
				fmt.Printf("Error: %v", err)
				os.Exit(0)
			}
		} else {
			err := migrator.Down()
			if err != nil {
				fmt.Printf("Error: %v", err)
				os.Exit(0)
			}
		}
	},
}

func init() {
	migrateCmd.Flags().StringVarP(&specificMigration, "name", "n", "", "Run specific migration")
	migrateCmd.Flags().StringVarP(&fromMigration, "from", "f", "", "Run from migration")
	migrateCmd.Flags().StringVarP(&toMigration, "to", "t", "", "Run to migration")
	migrateCmd.Flags().StringVarP(&env, "env", "e", "dev", "Run on environment")
	migratorCmd.AddCommand(migrateCmd)
}
