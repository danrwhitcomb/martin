package migrations

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"

	"martin/martin-device/app/models"
)

type SchemaMigration interface {
	Perform(database *gorm.DB) error
	Title() string
}

// Database migrations! Add new migration classes
// here and they will be run in order
var (
	log                       = revel.AppLog
	schemas []SchemaMigration = []SchemaMigration{InitialSetup{}}
)

type MigrationRunner struct{}

// Brings the given database up-to-date
func (c MigrationRunner) Run(database *gorm.DB) {

	// We are always going to have a starting schema.
	schemaVersion := 0
	schemaEntry, err := c.getSchemaEntry(database)
	if err == nil {
		log.Infof("Schema table found. Current Schema version: %d", schemaEntry.Version)
		schemaVersion = schemaEntry.Version
	}

	if schemaVersion != len(schemas) {
		log.Infof("Found out-of-date schema: %d. Starting to apply newer migrations", schemaVersion)
		err := c.applyMigrations(database, schemas[schemaVersion:])

		if err != nil {
			log.Fatalf("Schema migration failed with error: %s", err)
		}

		schemaEntry, err = c.getSchemaEntry(database)
		if err != nil {
			log.Fatal("Finished applying schema migrations, but no Schema data found?!?!?!")
		}

		log.Infof("Setting Schema version to %d", len(schemas))
		schemaEntry.Version = len(schemas)
		database.Save(&schemaEntry)
	} else {
		log.Info("DB schema is up-to-date. No new migrations to apply.")
	}

	log.Info("Schema migration complete.")
}

// Applys a set of migrations in order given
func (c MigrationRunner) applyMigrations(database *gorm.DB, migrations []SchemaMigration) error {
	for _, migration := range migrations {
		log.Debugf("Performing migration: %s", migration.Title())
		err := migration.Perform(database)
		if err != nil {
			return err
		}
	}

	return nil
}

// Attempts to retrieve the schema entry
func (c MigrationRunner) getSchemaEntry(database *gorm.DB) (models.Schema, error) {
	schemaEntry := models.Schema{}

	if database.HasTable(&models.Schema{}) {
		database.First(&schemaEntry)

		return schemaEntry, nil
	}

	return schemaEntry, errors.New("No table found for Schema model")
}
