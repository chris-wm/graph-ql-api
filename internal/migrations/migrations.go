package migrations

import (
	"github.com/electivetechnology/utility-library-go/logger"
	"github.com/electivetechnology/utility-library-go/migrator"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/adapter"
	"gorm.io/gorm"
)

var log logger.Logging
var version string

func init() {
	// Add generic logger
	log = logger.NewLogger("migrations")

	// Register Migration
	log.Printf("Registering migration with Model")
	adapter.RegisterMigration(MigrationUp)
}

func MigrationUp(db *gorm.DB) {
	log.Printf("Starting MigrationUp")

	versions := migrator.MigrateUp(db)
	log.Printf("Current Migration: %v", versions)

	MigrateUp1(db, versions)

	log.Printf("Finished MigrationUp")
}

func MigrationDown(db *gorm.DB) {
	log.Printf("Starting MigrationDown")

	MigrateDown1(db)

	log.Printf("Finished MigrationDown")
}
