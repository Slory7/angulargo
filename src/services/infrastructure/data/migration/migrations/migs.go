package migrations

import "services/infrastructure/data/migration"

var MigrationVersions = []*migration.Migration{
	v201810091551(),
	v201810311029(),
}
