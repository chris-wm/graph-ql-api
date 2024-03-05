package migrations

import (
	"github.com/electivetechnology/utility-library-go/migrator"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/adapter"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/entity"
	"gorm.io/gorm"
)

const (
	VERSION_20210924104500 = "20210924104500"
)

func MigrateUp20210924104500(db *gorm.DB, currentVersions []migrator.Migration) {
	if migrator.VersionExists(currentVersions, VERSION_20210924104500) {
		return
	}

	migrator.AddVersion(db, VERSION_20210924104500)

	// Templates
	db.Exec(
		"CREATE TABLE `templates` (" +
			"`id` VARCHAR(16) NOT NULL COLLATE 'utf8mb4_unicode_ci'," +
			"`created_at` DATETIME(3) NULL DEFAULT NULL," +
			"`updated_at` DATETIME(3) NULL DEFAULT NULL," +
			"`deleted_at` DATETIME(3) NULL DEFAULT NULL," +
			"`type` VARCHAR(16) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'," +
			"`name` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'," +
			"`visibility` VARCHAR(16) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'," +
			"`organisation` VARCHAR(16) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'," +
			"`content` LONGTEXT NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'," +
			"PRIMARY KEY (`id`) USING BTREE," +
			"INDEX `idx_templates_deleted_at` (`deleted_at`) USING BTREE" +
			")" +
			"COLLATE='utf8mb4_unicode_ci'" +
			"ENGINE=InnoDB;")
}

func MigrateDown20210924104500(db *gorm.DB) {
	migration := migrator.Migration{}
	migration.Version = VERSION_20210924104500

	adapter.GetDb().Unscoped().Where("version = ?", VERSION_20210924104500).Delete(&migration)

	db.Migrator().DropTable(entity.Template{})
}
