package postgres

import (
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

type MigrateConfig struct {
	MigrationsFilePath string
	ConnectionURI      string
}

func runMigrations(config *MigrateConfig, mode string, version uint) error {
	filePath := fmt.Sprintf("file://%s", config.MigrationsFilePath)

	m, err := migrate.New(filePath, config.ConnectionURI)
	if err != nil {
		panic(err)
	}

	defer m.Close()

	switch strings.ToLower(mode) {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "version":
		err = m.Migrate(version)
	default:
		err = fmt.Errorf("invalid migration mode '%v'", mode)
	}

	// DO NOT consider ErrNoChange as migration failure
	if err == migrate.ErrNoChange {
		err = nil
	}
	return err
}

// Apply all migrations under migrations file path
func MigrateUp(config *MigrateConfig) {
	// Migrations is always performed on default database
	err := runMigrations(config, "up", 0)
	if err != nil {
		log.Panicf("Migrations failed! Error: %v", err)
	} else {
		log.Infof("Migrations applied successfully")
	}
}

// Revert all migrations under migrations file path
// NOTE: This function is used only while testing to cleanup database
func MigrateDown(config *MigrateConfig) {
	err := runMigrations(config, "down", 0)
	if err != nil {
		log.Panicf("Migrations drop failed! Error: %v", err)
	} else {
		log.Infof("Migrations dropped successfully")
	}
}

func Migrate(config *MigrateConfig, version uint) {
	err := runMigrations(config, "version", version)
	if err != nil {
		log.Panicf("Migrations failed! Error: %v", err)
	} else {
		log.Infof("Migrations applied successfully")
	}
}
