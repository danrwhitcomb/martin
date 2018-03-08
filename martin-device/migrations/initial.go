package migrations

import (
	"github.com/jinzhu/gorm"

	"martin/martin-device/app/models"
)

// Overrides migrations.SchemaMigration
type InitialSetup struct{}

func (c InitialSetup) Perform(database *gorm.DB) error {
	c.setupSchemaVersioning(database)
	c.setupSystemModels(database)

	return nil
}

func (c InitialSetup) setupSchemaVersioning(db *gorm.DB) {
	schema := models.Schema{}
	if !db.HasTable(&schema) {
		db.CreateTable(&models.Schema{})
		db.Save(&models.Schema{Version: 0})
	}
}

func (c InitialSetup) setupSystemModels(db *gorm.DB) {
	master := models.Master{}
	if !db.HasTable(&master) {
		db.CreateTable(&master)
	}

	config := models.Configuration{}
	if !db.HasTable(&config) {
		db.CreateTable(&config)
	}

	system := models.System{}
	if !db.HasTable(&system) {
		db.CreateTable(&system)
		system.State = models.Unconfigured
		db.Save(&system)
	}
}

func (c InitialSetup) Title() string {
	return "Initial Setup"
}
