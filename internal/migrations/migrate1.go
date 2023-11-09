package migrations

import (
	"github.com/electivetechnology/utility-library-go/migrator"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/adapter"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/entity"
	"gorm.io/gorm"
)

const (
	VERSION_1 = "1"
)

func MigrateUp1(db *gorm.DB, currentVersions []migrator.Migration) {
	if migrator.VersionExists(currentVersions, VERSION_1) {
		return
	}

	migrator.AddVersion(db, VERSION_1)

	// Users
	db.Exec(
		"CREATE TABLE `users` (" +
			"`id` CHAR(36) NOT NULL COLLATE 'utf8mb4_unicode_ci'," +
			"`created_at` DATETIME(3) NULL DEFAULT NULL," +
			"`updated_at` DATETIME(3) NULL DEFAULT NULL," +
			"`deleted_at` DATETIME(3) NULL DEFAULT NULL," +
			"`name` VARCHAR(36) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'," +
			"`email` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'," +
			"`password` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'," +
			"PRIMARY KEY (`id`) USING BTREE)" +
			"COLLATE='utf8mb4_unicode_ci'" +
			"ENGINE=InnoDB;")

}

func MigrateDown1(db *gorm.DB) {
	migration := migrator.Migration{}
	migration.Version = VERSION_1

	adapter.GetDb().Unscoped().Where("version = ?", VERSION_1).Delete(&migration)

	db.Migrator().DropTable(entity.User{})
}
