package types

import (
	"errors"
	"fmt"
	"github.com/oswgg/migrator/internal/config"
	"github.com/oswgg/migrator/internal/database"
	"github.com/oswgg/migrator/pkg/tools"
)

type Migrator struct {
	Env               string
	Specific          bool
	SpecificMigration string
	MigrationType     string
	From              string
	To                string
	Migrations        []Migration
	Connection        database.DatabaseImpl
}

type Migration struct {
	Path string
	Name string
}

type MigrationRunner interface {
	Up() error
	Down() error
}

func (m *Migrator) Up() error {
	configurations, err := config.GetUserYAMLConfig(m.Env)
	if err != nil {
		return err
	}

	migrationsTableExists, err := m.Connection.VerifyTableExists(configurations.MigrationsTableName)
	if err != nil {
		return err
	}

	if !migrationsTableExists {
		err := m.Connection.CreateMigrationsTable()
		if err != nil {
			return err
		}
	}

	if len(m.Migrations) == 0 {
		fmt.Println("No migrations pending")
		return nil
	}

	for _, migration := range m.Migrations {
		readFile, err := tools.ReadFile(migration.Path)
		if err != nil {
			return err
		}
		fmt.Printf("========= Migrating: %s =========\n", migration.Name)
		err = m.Connection.ExecMigrationFileContent(string(readFile), migration.Name, "up")
		if err != nil {
			return err
		}
		fmt.Printf("========= Migrated: %s =========\n\n", migration.Name)
	}
	return nil
}

func (m *Migrator) Down() error {
	fmt.Printf("Down from %v", m.From)
	fmt.Printf("Down to %v", m.To)

	return nil
}
